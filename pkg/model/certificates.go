package model

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

const CertificateCollection = "certificates"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*Certificate)(nil)

type Certificate struct {
	SulfurRecordProxy
}

func (a *Certificate) Fingerprint() string {
	return a.GetString("fingerprint")
}
func (a *Certificate) SetFingerprint(val string) {
	a.Set("fingerprint", val)
}

func (a *Certificate) Issuer() string {
	return a.GetString("issuer")
}
func (a *Certificate) SetIssuer(val string) {
	a.Set("issuer", val)
}

func (a *Certificate) Subject() string {
	return a.GetString("subject")
}
func (a *Certificate) SetSubject(val string) {
	a.Set("subject", val)
}

func (a *Certificate) AlternativeNames() string {
	return a.GetString("alternative_names")
}
func (a *Certificate) SetAlternativeNames(val string) {
	a.Set("alternative_names", val)
}

func (a *Certificate) NotBefore() types.DateTime {
	return a.GetDateTime("not_before")
}

func (a *Certificate) SetNotBefore(lastSeen time.Time) {
	a.Set("not_before", lastSeen)
}
func (a *Certificate) Expires() types.DateTime {
	return a.GetDateTime("expires")
}

func (a *Certificate) SetExpires(lastSeen time.Time) {
	a.Set("expires", lastSeen)
}

func (a *Certificate) SetSerial(serialNumber *big.Int) {

	bytes := serialNumber.Bytes()
	hexWithColons := ""
	for i, b := range bytes {
		hexWithColons += fmt.Sprintf("%02x", b)
		if i < len(bytes)-1 {
			hexWithColons += ":"
		}
	}

	//// Convert the big.Int serial number to a byte slice
	//serialBytes := serialNumber.Bytes()
	//// Encode the byte slice to a hexadecimal string
	//serialHex := hex.EncodeToString(serialBytes)

	a.Set("serial", hexWithColons)
}

func CertificateFirstOrCreate(fingerprint string) (iface.Certificate, error) {
	ipPort := &Certificate{}

	record, err := FirstOrCreateByFilter(
		CertificateCollection,
		//"name={:name} && value={:value} && type={:type}",
		"fingerprint={:fingerprint}",

		dbx.Params{
			"fingerprint": fingerprint,
		})

	if err != nil {
		return nil, err
	}
	ipPort.SetProxyRecord(record)
	return ipPort, nil
}

func GetAllCertificates() ([]iface.Certificate, error) {
	ret := []iface.Certificate{}
	records, err := app_state.GetApp().FindAllRecords(CertificateCollection)

	for _, record := range records {

		// load into proxy
		proxyStruct := &Certificate{}
		proxyStruct.SetProxyRecord(record)

		ret = append(ret, proxyStruct)

	}

	return ret, err
}

func GetAllDomainsFromCertificates() ([]string, error) {
	certs, err := GetAllCertificates()
	if err != nil {
		return nil, err
	}

	domainMap := map[string]bool{}
	domains := []string{}
	for _, cert := range certs {
		domainMap[strings.ToLower(cert.Subject())] = true
		for _, alt := range strings.Split(cert.AlternativeNames(), ",") {
			domainMap[strings.ToLower(alt)] = true
		}
	}

	for domain := range domainMap {
		domains = append(domains, domain)
	}
	return domains, nil
}
