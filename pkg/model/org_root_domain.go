package model

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"golang.org/x/net/publicsuffix"
	"strings"
	"time"
)

const OrgRootDomainCollection = "org_domains"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*OrgRootDomain)(nil)

type OrgRootDomain struct {
	SulfurRecordProxy
}

func (a *OrgRootDomain) Registrar() string {
	return a.GetString("registrar")
}

func (a *OrgRootDomain) DomainName() string {
	return a.GetString("domain")
}

func (a *OrgRootDomain) LastSeen() types.DateTime {
	return a.GetDateTime("last_seen")
}

func (a *OrgRootDomain) DNSRecords() []iface.DNSRecord {
	var s []iface.DNSRecord

	m := &DNSRecord{}

	s = append(s, m)
	return s
}

//func (a *RootDomain) SubDomains() (domains []*RootDomain) {
//	return domains
//}

func (a *OrgRootDomain) SetRegistrar(registrar string) {
	a.Set("registrar", registrar)
}

func (a *OrgRootDomain) SetLastSeen(lastSeen time.Time) {
	a.Set("last_seen", lastSeen)
}

func FindOrgRootDomain(rootDomainName string) (iface.OrgRootDomain, error) {

	if strings.HasSuffix(rootDomainName, ".") {
		rootDomainName = rootDomainName[:len(rootDomainName)-1]
	}
	rootDomain, err := publicsuffix.EffectiveTLDPlusOne(rootDomainName)

	rdr := &OrgRootDomain{}

	err = app_state.GetApp().RecordQuery(OrgRootDomainCollection).
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

func OrgRootDomainFirstOrCreate(rootDomainID, orgID string) (iface.OrgRootDomain, error) {
	dnsR := &OrgRootDomain{}

	record, err := FirstOrCreateByFilter(
		OrgRootDomainCollection,
		//"name={:name} && value={:value} && type={:type}",
		"root_domain={:root_domain} AND organization={:organization}",

		dbx.Params{
			"root_domain":  rootDomainID,
			"organization": orgID,
		})

	if err != nil {
		return nil, err
	}
	dnsR.SetProxyRecord(record)
	return dnsR, nil
}
