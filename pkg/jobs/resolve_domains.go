package jobs

import (
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/projectdiscovery/dnsx/libs/dnsx"
	"log"
	"net"
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
	domainsToResolve, err := model.GetARecordsToResolve()
	if err != nil {
		log.Println(err)
		return
	}

	_, sharedSpace, _ := net.ParseCIDR("100.64.0.0/10")

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
					log.Println("error creating dns record", err, result[i].Name, v)
					continue
				}

				r.SetLastResolved(time.Now())
				r.SetResolveErr("")
				if err := r.Save(); err != nil {
					log.Println("FAILED TO SAVE RECORD", result[i].Name, err)
				}

				parsedIP := net.ParseIP(v)
				if parsedIP != nil {
					ipRecord, err := model.IPAddressFirstOrCreate(parsedIP.String())
					if err != nil {
						log.Println("failed to create ip addr", err)
					}

					ipRecord.SetIs6(parsedIP.To4() == nil)
					// not sure what i was thinking whn i created this field....
					//ipRecord.SetIsEphemeral(parsedIP.IsEp)
					ipRecord.SetIsGlobalUnicast(parsedIP.IsGlobalUnicast())
					ipRecord.SetIsInterfaceLocalMulticast(parsedIP.IsLinkLocalMulticast())
					ipRecord.SetIsLoopback(parsedIP.IsLoopback())
					ipRecord.SetIsLinkLocalMulticast(parsedIP.IsLinkLocalMulticast())
					ipRecord.SetIsLinkLocalUnicast(parsedIP.IsLinkLocalUnicast())
					ipRecord.SetIsMulticast(parsedIP.IsMulticast())
					ipRecord.SetIsPrivate(parsedIP.IsPrivate())
					ipRecord.SetIsShared(sharedSpace.Contains(parsedIP))
					ipRecord.SetIsUnspecified(parsedIP.IsUnspecified())

					if err := ipRecord.Save(); err != nil {
						log.Println("failed to save IP", err)
					}
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
