package jobs

import (
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/naabu/v2/pkg/result"
	"github.com/projectdiscovery/naabu/v2/pkg/runner"
	"golang.org/x/net/context"
	"log"
	"time"
)

type SimplePortScanResults struct {
	IPAddrRecord iface.IPAddress
	HostResults  *result.HostResult
}

func SimplePortScan() {
	domainsToResolve, err := model.GetSimplePortScanInput()
	if err != nil {
		log.Println(err)
		return
	}

	total := len(domainsToResolve)
	log.Printf("total ips to scan: %v\n", total)

	hosts := goflags.StringSlice{}

	for _, record := range domainsToResolve {
		hosts = append(hosts, record.Host)
	}

	scaRes, err := RunScan(hosts)
	if err != nil {
		log.Println(err)
		return
	}

	for _, hr := range scaRes {

		ipAddrStr := hr.Host
		ipAddr, err := model.IPAddressFirstOrCreate(ipAddrStr)
		if err != nil {
			log.Println(err)
			continue
		}

		if ipAddr.Id() == "" {
			err = ipAddr.Save()
			if err != nil {
				log.Println("unable to save new ip", err)
				continue
			}
		}

		for _, v := range hr.Ports {
			r, err := model.IPPortFirstOrCreate(ipAddr.Id(), v.Port, v.Protocol.String())
			if err != nil {
				log.Println("error saving ip port combo", err)
				continue
			}

			r.SetLastSeen(time.Now())
			err = r.Save()
			if err != nil {
				log.Println("FAILED TO SAVE RECORD", ipAddr.Address(), err)
				continue
			}

			ipAddr.SetLastSimplePortScan(time.Now())
			if err := ipAddr.Save(); err != nil {
				log.Println("Error updating last scan date", err)
			}
		}
	}
}

func SimplePortScanWorkers() {
	domainsToResolve, err := model.GetSimplePortScanInput()
	if err != nil {
		log.Println(err)
		return
	}

	total := len(domainsToResolve)
	log.Printf("total ips to scan: %v\n", total)

	input := make(chan *SimplePortScanResults, total)
	output := make(chan *SimplePortScanResults, total)

	for w := 1; w <= 20; w++ {
		go PortScanWorker(input, output)
	}

	for _, record := range domainsToResolve {

		ipAddr, err := model.IPAddressFirstOrCreate(record.Host)
		if err != nil {
			log.Println("error getting ip", err)
			continue
		}

		input <- &SimplePortScanResults{
			IPAddrRecord: ipAddr,
		}
	}

	result := make([]*SimplePortScanResults, total)
	for i, _ := range result {
		result[i] = <-output
		for _, v := range result[i].HostResults.Ports {

			r, err := model.IPPortFirstOrCreate(result[i].IPAddrRecord.Id(), v.Port, v.Protocol.String())
			if err != nil {
				log.Println("error saving ip port combo", err)
				continue
			}

			r.SetLastSeen(time.Now())
			err = r.Save()
			if err != nil {
				log.Println("FAILED TO SAVE RECORD", result[i].IPAddrRecord.Address(), err)
				continue
			}

			result[i].IPAddrRecord.SetLastSimplePortScan(time.Now())
			if err := result[i].IPAddrRecord.Save(); err != nil {
				log.Println("Error updating last scan date", err)
			}

		}
	}
}

func PortScanWorker(input chan *SimplePortScanResults, output chan *SimplePortScanResults) {
	for scanRes := range input {
		hr, err := RunScan(goflags.StringSlice{scanRes.IPAddrRecord.Address()})
		if err != nil {
			log.Println(err)
		}
		scanRes.HostResults = hr[0]
		output <- scanRes
	}
}

func RunScan(hosts goflags.StringSlice) ([]*result.HostResult, error) {
	var hostResults []*result.HostResult

	options := runner.Options{
		Host:               hosts,
		InputReadTimeout:   500 * time.Millisecond,
		Stream:             true,
		DisableStdin:       true,
		DisableUpdateCheck: true,
		ScanType:           "c",
		OnResult: func(hr *result.HostResult) {
			hostResults = append(hostResults, hr)
		},
		Ports:   "443",
		Silent:  true,
		Timeout: 500 * time.Millisecond,

		Retries: 0,
		//Debug:   true,
		//Verbose: true,
	}

	naabuRunner, err := runner.NewRunner(&options)
	if err != nil {
		log.Fatal(err)
	}
	defer naabuRunner.Close()

	err = naabuRunner.RunEnumeration(context.Background())
	if err != nil {
		return nil, err
	}

	return hostResults, nil
}
