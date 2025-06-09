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

			r, err := model.IPPortFirstOrCreate(result[i].IPAddrRecord.ProxyRecord().Id, v.Port, v.Protocol.String())
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
		options := runner.Options{
			Host:     goflags.StringSlice{scanRes.IPAddrRecord.Address()},
			ScanType: "c",
			OnResult: func(hr *result.HostResult) {
				scanRes.HostResults = hr
				output <- scanRes
			},
			Ports:   "80,443",
			Silent:  true,
			Timeout: 45 * time.Second,
		}

		naabuRunner, err := runner.NewRunner(&options)
		if err != nil {
			log.Fatal(err)
		}
		defer naabuRunner.Close()

		err = naabuRunner.RunEnumeration(context.Background())
		if err != nil {
			log.Println("error running port scan", err, scanRes.IPAddrRecord.Address())
		}
	}
}
