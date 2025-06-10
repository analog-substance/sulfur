package dns

import (
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/projectdiscovery/dnsx/libs/dnsx"
	"log"
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

func ResolveDomains(domainsToResolve []string) []sulfur.DNSRecord {
	results := []sulfur.DNSRecord{}

	total := len(domainsToResolve)
	log.Printf("total records: %v\n", total)

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
			log.Println(result[i].Error)
		} else {
			results = append(results, result[i].DNSRecords...)
		}
	}

	return results
}

type checkDNSStatus struct {
	Name       string
	DNSRecords []sulfur.DNSRecord
	Error      error
}

func CheckDNSWorker(input chan string, output chan checkDNSStatus) {
	for record := range input {

		resolvable, err := getRecords(record)

		output <- checkDNSStatus{
			Name:       record,
			DNSRecords: resolvable,
			Error:      err,
		}
	}
}

func getRecords(record string) ([]sulfur.DNSRecord, error) {
	dnsRecords := []sulfur.DNSRecord{}
	results, err := dnsClient.QueryMultiple(record)
	if err != nil {
		if err.Error() == "no ips found" {
			return nil, nil
		}
		return nil, err
	}

	for _, res := range results.A {
		dnsRecords = append(dnsRecords, sulfur.DNSRecord{
			Name:  results.Host,
			Value: res,
			Type:  "A",
			TTL:   int(results.TTL),
		})

	}
	return dnsRecords, nil
}
