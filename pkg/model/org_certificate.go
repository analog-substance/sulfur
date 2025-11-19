package model

import (
	"time"

	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

const OrgCertificateCollection = "org_certificates"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*OrgCertificate)(nil)

type OrgCertificate struct {
	LastSeenExternalReferenceRecordProxy
}

func (a *OrgCertificate) Serial() string {
	return a.GetString("serial")
}

func (a *OrgCertificate) SetSerial(serial string) {
	a.Set("serial", serial)
}

func (a *OrgCertificate) Issuer() string {
	return a.GetString("issuer")
}
func (a *OrgCertificate) SetIssuer(issuer string) {
	a.Set("issuer", issuer)
}

func (a *OrgCertificate) Subject() string {
	return a.GetString("subject")
}
func (a *OrgCertificate) SetSubject(subject string) {
	a.Set("subject", subject)
}

func (a *OrgCertificate) SubjectAlternativeNames() string {
	return a.GetString("subject_alternative_names")
}
func (a *OrgCertificate) SetSubjectAlternativeNames(subjectAlternativeNames string) {
	a.Set("subject_alternative_names", subjectAlternativeNames)
}

func (a *OrgCertificate) NotBefore() types.DateTime {
	return a.GetDateTime("not_before")
}

func (a *OrgCertificate) SetNotBefore(notBefore time.Time) {
	a.Set("not_before", notBefore)
}

func (a *OrgCertificate) NotAfter() types.DateTime {
	return a.GetDateTime("not_after")
}

func (a *OrgCertificate) SetNotAfter(notAfter time.Time) {
	a.Set("not_after", notAfter)
}

func OrgCertificateFirstOrCreate(serial, orgID string) (iface.OrgCertificate, error) {
	dnsR := &OrgCertificate{}

	record, err := FirstOrCreateByFilter(
		OrgCertificateCollection,
		//"name={:name} && value={:value} && type={:type}",
		"serial={:serial} AND organization={:organization}",

		dbx.Params{
			"serial":       serial,
			"organization": orgID,
		})

	if err != nil {
		return nil, err
	}
	dnsR.SetProxyRecord(record)
	return dnsR, nil
}
