package jobs

import (
	"errors"
	copper "github.com/analog-substance/copper/pkg/lib"
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/model"
	"time"
)

var portsToScan = []int{}

func init() {
	portsToScan = copper.GetTopPopularPorts("tcp", 1000)
}

type SimplePortScanResults struct {
	IPAddr      string
	SimplePorts []simplePort
}

type simplePort struct {
	Port     int
	Protocol string
}

func SimplePortScanWorkers() {
	logger := app_state.GetApp().Logger().WithGroup("SimplePortScan")
	ipsToScan, err := model.GetSimplePortScanInput()
	if err != nil {
		logger.Error("Error getting IPs to scan", "error", err)
		return
	}

	total := len(ipsToScan)
	logger.Info("preparing to scan", "count", total)

	input := make(chan *SimplePortScanResults, total)
	output := make(chan *SimplePortScanResults, total)

	for w := 1; w <= 20; w++ {
		go PortScanWorker(input, output)
	}

	queue := 0
	for _, record := range ipsToScan {
		queue++
		input <- &SimplePortScanResults{
			IPAddr:      record.Host,
			SimplePorts: []simplePort{},
		}
	}
	close(input)
	logger.Info("queued", "count", queue)

	for a := queue; a > 0; a-- {
		portScanResult := <-output

		err := SaveOpenPortsForIP(portScanResult.IPAddr, portScanResult.SimplePorts)
		if err != nil {
			logger.Error("Failed to save port scan results", "error", err)
			continue
		}
	}
	close(output)
}

func PortScanWorker(input chan *SimplePortScanResults, output chan *SimplePortScanResults) {
	for scanRes := range input {

		openPorts := copper.GetOpenPortsOnHost(scanRes.IPAddr, portsToScan, 500)
		for _, openPort := range openPorts {
			scanRes.SimplePorts = append(scanRes.SimplePorts, simplePort{
				Port:     openPort,
				Protocol: "tcp",
			})
		}

		output <- scanRes
	}
}

func SaveOpenPortsForIP(ipAddrStr string, ports []simplePort) error {
	ipAddr, err := model.IPAddressFirstOrCreate(ipAddrStr)
	if err != nil {
		return errors.New("failed to get IP model for results: " + err.Error())
	}

	// we can update the time to now.
	ipAddr.SetLastSimplePortScan(time.Now())
	if err := ipAddr.Save(); err != nil {
		return errors.New("failed to update last scan date for IP: " + err.Error())
	}

	for _, v := range ports {
		r, err := model.IPPortFirstOrCreate(ipAddr.Id(), v.Port, v.Protocol)
		if err != nil {
			return errors.New("failed to get IP Port model: " + err.Error())
		}

		r.SetLastSeen(time.Now())
		err = r.Save()
		if err != nil {
			return errors.New("failed to save IP Port model: " + err.Error())
		}

	}
	return nil
}
