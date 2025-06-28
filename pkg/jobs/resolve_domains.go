package jobs

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/miekg/dns"
	"github.com/projectdiscovery/dnsx/libs/dnsx"
	"github.com/projectdiscovery/retryabledns"
	"log"
	"log/slog"
	"time"
)

var dnsClient *retryabledns.Client

func init() {
	var err error
	// Create DNS Resolver with default options
	retryablednsOptions := retryabledns.Options{
		BaseResolvers: dnsx.DefaultOptions.BaseResolvers,
		MaxRetries:    dnsx.DefaultOptions.MaxRetries,
		Hostsfile:     dnsx.DefaultOptions.Hostsfile,
		Proxy:         dnsx.DefaultOptions.Proxy,
	}

	dnsClient, err = retryabledns.NewWithOptions(retryablednsOptions)
	if err != nil {
		log.Panic("err: %v\n", err)
	}
	dnsClient.TCPFallback = true
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
	domainsToResolve, err := model.GetDNSLookupQueue()
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
			logger.Error("error in resolution results", "error", result[i].Error, "dnsData", result[i].DNSData)
		} else {
			saveRecords("A", result[i].Name, result[i].DNSData.A, logger)
			saveRecords("AAAA", result[i].Name, result[i].DNSData.AAAA, logger)
			saveRecords("CNAME", result[i].Name, result[i].DNSData.CNAME, logger)
			saveRecords("NS", result[i].Name, result[i].DNSData.NS, logger)
			saveRecords("MX", result[i].Name, result[i].DNSData.MX, logger)
			saveRecords("TXT", result[i].Name, result[i].DNSData.TXT, logger)
			saveRecords("SRV", result[i].Name, result[i].DNSData.SRV, logger)
		}
	}
}

type checkDNSStatus struct {
	Name    string
	DNSData *retryabledns.DNSData
	Error   error
}

func CheckDNSWorker(input chan string, output chan checkDNSStatus) {
	for record := range input {
		dnsData, err := DNSQueryMultiple(record, []uint16{dns.TypeA, dns.TypeAAAA, dns.TypeCNAME, dns.TypeTXT, dns.TypeNS, dns.TypeMX, dns.TypeSRV})
		output <- checkDNSStatus{
			Name:    record,
			DNSData: dnsData,
			Error:   err,
		}
	}
}

func saveRecords(recordType string, name string, values []string, logger *slog.Logger) {
	for _, v := range values {
		r, err := model.DNSRecordFirstOrCreate(name, v, recordType)
		if err != nil {
			logger.Error("error in dns record", "error", err, "name", name, "value", v)
			continue
		}

		r.SetLastResolved(time.Now())
		r.SetResolveErr("")
		if err := r.Save(); err != nil {
			logger.Error("Failed to save DNS record", "name", name, "error", err)
		}
	}
}

func DNSQueryMultiple(record string, types []uint16) (*retryabledns.DNSData, error) {
	return dnsClient.QueryMultiple(record, types)
}
