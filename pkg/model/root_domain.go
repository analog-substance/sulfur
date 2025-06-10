package model

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"golang.org/x/net/publicsuffix"
	"strings"
)

const RootDomainCollection = "root_domains"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*RootDomain)(nil)

type RootDomain struct {
	SulfurRecordProxy
}

func (a *RootDomain) DomainName() string {
	return a.GetString("domain")
}

func (a *RootDomain) DNSRecords() []iface.DNSRecord {
	var s []iface.DNSRecord

	m := &DNSRecord{}

	s = append(s, m)
	return s
}

//func (a *RootDomain) SubDomains() (domains []*RootDomain) {
//	return domains
//}

func FindRootDomain(domain string) (iface.RootDomain, error) {
	rootDomain, err := ToRootDomain(domain)

	if err != nil {
		return nil, err
	}

	rdr := &RootDomain{}

	err = app_state.GetApp().RecordQuery(RootDomainCollection).
		AndWhere(dbx.NewExp("LOWER(domain)={:domain}", dbx.Params{
			"domain": strings.ToLower(rootDomain),
		})).
		Limit(1).
		One(rdr)

	if err != nil {
		return nil, err
	}

	return rdr, nil

}

func RootDomainFirstOrCreate(domain string) (iface.RootDomain, error) {
	dnsR := &RootDomain{}

	domain, err := ToRootDomain(domain)
	if err != nil {
		return nil, err
	}

	record, err := FirstOrCreateByFilter(
		RootDomainCollection,
		"LOWER(domain)={:domain}",

		dbx.Params{
			"domain": strings.ToLower(domain),
		})

	if err != nil {
		return nil, err
	}
	dnsR.SetProxyRecord(record)
	return dnsR, nil
}

func ToRootDomain(domain string) (string, error) {
	domain = NormalizeDomain(domain)
	return publicsuffix.EffectiveTLDPlusOne(domain)
}

func NormalizeDomain(domain string) string {
	if strings.HasSuffix(domain, ".") {
		domain = domain[:len(domain)-1]
	}
	return domain
}
