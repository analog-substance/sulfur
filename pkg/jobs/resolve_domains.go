package jobs

import (
	"fmt"
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/projectdiscovery/dnsx/libs/dnsx"
	"log"
	"log/slog"
	"strings"
	"time"
)

var dnsClient *dnsx.DNSX

func init() {
	var err error
	// Create DNS Resolver with default options
	dnsClient, err = dnsx.New(dnsx.DefaultOptions)
	if err != nil {
		log.Panic("err: %v\n", err)
	}
}

func ResolveCertificateDomains() {
	logger := app_state.GetApp().Logger().WithGroup("ResolveCertificateDomains")
	domainsToResolve, err := model.GetAllDomainsFromCertificates()
	if err != nil {
		logger.Error("failed to get domain queue: ", "err", err)
		return
	}

	ResolveDomains(domainsToResolve, logger)
}

func ResolveDNSRecordDomains() {
	logger := app_state.GetApp().Logger().WithGroup("ResolveDNSRecordDomains")
	domainsToResolve, err := model.GetARecordsToResolve()
	if err != nil {
		logger.Error("failed to get domain queue: ", "err", err)
		return
	}

	domains := []string{}
	for _, domain := range domainsToResolve {
		domains = append(domains, domain.Host)
	}

	ResolveDomains(domains, logger)
}

func ResolveDomains(domainsToResolve []string, logger *slog.Logger) {

	total := len(domainsToResolve)
	logger.Info("records to resolve", "total", total)

	input := make(chan string, total)
	output := make(chan checkDNSStatus, total)

	for w := 1; w <= 20; w++ {
		go CheckDNSWorker(input, output)
	}

	for _, record := range domainsToResolve {
		input <- record
	}

	result := make([]checkDNSStatus, total)
	for i, _ := range result {
		result[i] = <-output
		if result[i].Error != nil {
			logger.Error("error in resolution results", "error", result[i].Error)
		} else {
			for _, v := range result[i].Value {
				r, err := model.DNSRecordFirstOrCreate(result[i].Name, v, "A")
				if err != nil {
					logger.Error("error in dns record", "error", err, "name", result[i].Name, "value", v)
					continue
				}

				r.SetLastResolved(time.Now())
				r.SetResolveErr("")
				if err := r.Save(); err != nil {
					logger.Error("Failed to save DNS record", "name", result[i].Name, "error", err)
				}
			}
		}
	}
}

type checkDNSStatus struct {
	Name  string
	Value []string
	Error error
}

func CheckDNSWorker(input chan string, output chan checkDNSStatus) {
	for record := range input {

		resolvable, err := getRecords(record)

		output <- checkDNSStatus{
			Name:  record,
			Value: resolvable,
			Error: err,
		}
	}
}

func getRecords(record string) ([]string, error) {

	// check if it is a wildcard, if so change to random value
	if strings.HasPrefix(record, "*.") {
		record = strings.Replace(record, "*", fmt.Sprintf("rand-%s", time.Now().String()), 1)
	}

	app_state.GetApp().Logger().Debug("looking up dns record", "record", record)

	result, err := dnsClient.Lookup(record)
	if err != nil {
		if err.Error() == "no ips found" {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}
