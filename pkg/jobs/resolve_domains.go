package jobs

import (
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/projectdiscovery/dnsx/libs/dnsx"
	"log"
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

func ResolveDomains() {
	domainsToResolve, err := model.GetAuditScopeDNS()
	if err != nil {
		log.Println(err)
		return
	}

	total := len(domainsToResolve)
	log.Printf("total records: %v\n", total)

	input := make(chan string, total)
	output := make(chan checkDNSStatus, total)

	for w := 1; w <= 20; w++ {
		go CheckDNSWorker(input, output)
	}

	for _, record := range domainsToResolve {
		input <- record.Host
	}

	result := make([]checkDNSStatus, total)
	for i, _ := range result {
		result[i] = <-output
		if result[i].Error != nil {
			log.Println(result[i].Error)
		} else {
			for _, v := range result[i].Value {
				r, err := model.DNSRecordFirstOrCreate(result[i].Name, v, "A")
				if err != nil {
					log.Println(err)
					continue
				}

				r.SetLastResolved(time.Now())
				r.SetResolveErr("")

				err = r.Save()
				if err != nil {
					log.Println("FAILED TO SAVE RECORD", result[i].Name, err)
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

	result, err := dnsClient.Lookup(record)
	if err != nil {
		if err.Error() == "no ips found" {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}
