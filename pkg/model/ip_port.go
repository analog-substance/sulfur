package model

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

const IPPortCollection = "ip_ports"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*IPPort)(nil)

type IPPort struct {
	SulfurRecordProxy
}

func (a *IPPort) IPAddress() string {
	return a.GetString("ip_address")
}
func (a *IPPort) SetIPAddress(val string) {
	a.Set("ip_address", val)
}

func (a *IPPort) Port() int {
	return a.GetInt("port")
}
func (a *IPPort) SetPort(val int) {
	a.Set("port", val)
}

func (a *IPPort) Banner() string {
	return a.GetString("banner")
}
func (a *IPPort) SetBanner(val string) {
	a.Set("banner", val)
}
func (a *IPPort) Service() string {
	return a.GetString("service")
}
func (a *IPPort) SetService(val string) {
	a.Set("service", val)
}

func (a *IPPort) AppProtocol() string {
	return a.GetString("app_protocol")
}
func (a *IPPort) SetAppProtocol(val string) {
	a.Set("app_protocol", val)
}

func (a *IPPort) Protocol() string {
	return a.GetString("protocol")
}
func (a *IPPort) SetProtocol(val string) {
	a.Set("protocol", val)
}

func (a *IPPort) LastSeen() types.DateTime {
	return a.GetDateTime("last_seen")
}

func (a *IPPort) SetLastSeen(lastSeen time.Time) {
	a.Set("last_seen", lastSeen)
}

func (a *IPPort) GetIP() iface.IPAddress {

	coreRecord, err := FindRecordByID(IPAddressesCollection, a.IPAddress())
	if err != nil {
		log.Println("failed to get IP", err)
		return nil
	}

	// load into proxy
	obj := &IPAddress{}
	obj.SetProxyRecord(coreRecord)

	return obj
}

func IPPortFirstOrCreate(ipAddrId string, port int, protocol string) (iface.IPPort, error) {
	ipPort := &IPPort{}

	record, err := FirstOrCreateByFilter(
		IPPortCollection,
		//"name={:name} && value={:value} && type={:type}",
		"ip_address={:ip_address} AND port={:port} AND protocol={:protocol}",

		dbx.Params{
			"ip_address": ipAddrId,
			"port":       port,
			"protocol":   protocol,
		})

	if err != nil {
		return nil, err
	}
	ipPort.SetProxyRecord(record)
	return ipPort, nil
}

func GetIPsWithPorts(ports ...int) ([]iface.IPPort, error) {
	ret := []iface.IPPort{}
	portStr := ""
	for _, port := range ports {
		if len(portStr) > 0 {
			portStr += ","
		}
		portStr = portStr + strconv.Itoa(port)
	}

	// retrieve multiple "articles" records with optional dbx expressions
	records, err := app_state.GetApp().FindAllRecords(IPPortCollection,
		dbx.NewExp(fmt.Sprintf("port in (%s)", portStr)),
	)

	for _, record := range records {

		// load into proxy
		ipPort := &IPPort{}
		ipPort.SetProxyRecord(record)

		ret = append(ret, ipPort)

	}

	return ret, err
}

const SQLCertQueue = `
SELECT ip_ports.* 
FROM ip_ports
INNER JOIN ip_addresses on ip_addresses.id = ip_ports.ip_address
LEFT JOIN ip_port_certificates on ip_port_certificates.ip_port = ip_ports.id 
WHERE ip_ports.port = 443
AND (ip_port_certificates.id IS NULL OR ip_port_certificates.last_seen < datetime('now', '-4 hours'))
LIMIT 1000
`

type dbId struct {
	Id string `db:"id"`
}

func GetCertsQueue() ([]iface.IPPort, error) {
	ipPorts := []dbId{}
	err := app_state.GetApp().DB().
		NewQuery(SQLCertQueue).
		All(&ipPorts)

	ret := []iface.IPPort{}
	idStr := ""
	for _, port := range ipPorts {
		if len(idStr) > 0 {
			idStr += ","
		}
		idStr += fmt.Sprintf(`"%s"`, port.Id)
	}

	// retrieve multiple "articles" records with optional dbx expressions
	records, err := app_state.GetApp().FindAllRecords(IPPortCollection,
		dbx.NewExp(fmt.Sprintf("id in (%s)", idStr)),
	)

	for _, record := range records {

		// load into proxy
		ipPort := &IPPort{}
		ipPort.SetProxyRecord(record)

		ret = append(ret, ipPort)

	}

	return ret, err
}
