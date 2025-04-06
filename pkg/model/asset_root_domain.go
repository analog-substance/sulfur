package model

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"golang.org/x/net/publicsuffix"
	"strings"
)

const AssetRootDomainCollection = "asset_root_domains"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*AssetRootDomain)(nil)

type AssetRootDomain struct {
	core.BaseRecordProxy
}

func (a *AssetRootDomain) Save() error {
	return app_state.GetApp().Save(a)
}

func (a *AssetRootDomain) Registrar() string {
	return a.GetString("registrar")
}

func (a *AssetRootDomain) DomainName() string {
	return a.GetString("domain")
}

func (a *AssetRootDomain) DNSRecords() []iface.DNSRecord {
	var s []iface.DNSRecord

	m := &DNSRecord{}

	s = append(s, m)
	return s
}

//func (a *AssetRootDomain) SubDomains() (domains []*AssetRootDomain) {
//	return domains
//}

func (a *AssetRootDomain) Created() types.DateTime {
	return a.GetDateTime("created")
}

func (a *AssetRootDomain) Updated() types.DateTime {
	return a.GetDateTime("updated")
}

func FindAssetRootDomain(rootDomainName string) (iface.AssetRootDomain, error) {

	if strings.HasSuffix(rootDomainName, ".") {
		rootDomainName = rootDomainName[:len(rootDomainName)-1]
	}
	rootDomain, err := publicsuffix.EffectiveTLDPlusOne(rootDomainName)

	rdr := &AssetRootDomain{}

	err = app_state.GetApp().RecordQuery(AssetRootDomainCollection).
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

func AssetRootDomainFirstOrCreate(domain, orgID string) (iface.AssetRootDomain, error) {
	dnsR := &AssetRootDomain{}

	record, err := FirstOrCreateByFilter(
		AssetRootDomainCollection,
		//"name={:name} && value={:value} && type={:type}",
		"LOWER(domain)={:name} AND organization={:orgID}",

		dbx.Params{
			"name":         strings.ToLower(domain),
			"organization": orgID,
		})

	if err != nil {
		return nil, err
	}
	dnsR.SetProxyRecord(record)
	return dnsR, nil
}
