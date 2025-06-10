package model

import (
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"time"
)

const IPPortCertificateCollection = "ip_port_certificates"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*IPPortCertificate)(nil)

type IPPortCertificate struct {
	SulfurRecordProxy
}

func (a *IPPortCertificate) IPPort() string {
	return a.GetString("ip_port")
}
func (a *IPPortCertificate) SetIPPort(val string) {
	a.Set("ip_port", val)
}

func (a *IPPortCertificate) Certificate() string {
	return a.GetString("certificate")
}
func (a *IPPortCertificate) SetCertificate(val string) {
	a.Set("certificate", val)
}

func (a *IPPortCertificate) LastSeen() types.DateTime {
	return a.GetDateTime("last_seen")
}

func (a *IPPortCertificate) SetLastSeen(lastSeen time.Time) {
	a.Set("last_seen", lastSeen)
}

func IPPortCertificateFirstOrCreate(ip_port, certificate string) (iface.IPPortCertificate, error) {
	portCert := &IPPortCertificate{}

	record, err := FirstOrCreateByFilter(
		IPPortCertificateCollection,
		//"name={:name} && value={:value} && type={:type}",
		"ip_port={:ip_port} AND certificate={:certificate}",

		dbx.Params{
			"ip_port":     ip_port,
			"certificate": certificate,
		})

	if err != nil {
		return nil, err
	}
	portCert.SetProxyRecord(record)
	return portCert, nil
}

type SimpleIPPort struct {
	IPAddress string `db:"address"`
	Domain    string `db:"domain"`
	Port      int    `db:"port"`
}
