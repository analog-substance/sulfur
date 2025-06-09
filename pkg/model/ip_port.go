package model

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"time"
)

const IPPortCollection = "ip_ports"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*IPPort)(nil)

type IPPort struct {
	core.BaseRecordProxy
}

func (a *IPPort) Save() error {
	return app_state.GetApp().Save(a)
}

func (a *IPPort) Created() types.DateTime {
	return a.GetDateTime("created")
}

func (a *IPPort) Updated() types.DateTime {
	return a.GetDateTime("updated")
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
