package model

import (
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

const IPPortCertificateCollection = "ip_port_certificates"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*IPPortCertificate)(nil)

type IPPortCertificate struct {
	LastSeenRecordProxy
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
