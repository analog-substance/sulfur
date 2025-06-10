package model

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"time"
)

const CertificateCollection = "certificates"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*Certificate)(nil)

type Certificate struct {
	core.BaseRecordProxy
}

func (a *Certificate) Save() error {
	return app_state.GetApp().Save(a)
}

func (a *Certificate) Created() types.DateTime {
	return a.GetDateTime("created")
}

func (a *Certificate) Updated() types.DateTime {
	return a.GetDateTime("updated")
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

func (a *Certificate) Issued() types.DateTime {
	return a.GetDateTime("issued")
}

func (a *Certificate) SetIssued(lastSeen time.Time) {
	a.Set("issued", lastSeen)
}
func (a *Certificate) Expires() types.DateTime {
	return a.GetDateTime("expires")
}

func (a *Certificate) SetExpires(lastSeen time.Time) {
	a.Set("expires", lastSeen)
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
