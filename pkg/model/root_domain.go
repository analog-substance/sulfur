package model

import (
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"strings"
)

const EnvRootDomains = "env_root_domains"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*RootDomain)(nil)

type RootDomain struct {
	core.BaseRecordProxy
}

func (a *RootDomain) Save() error {
	return GetApp().Save(a)
}

func (a *RootDomain) Registrar() string {
	return a.GetString("registrar")
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

func (a *RootDomain) Created() types.DateTime {
	return a.GetDateTime("created")
}

func (a *RootDomain) Updated() types.DateTime {
	return a.GetDateTime("updated")
}

func FindRootDomain(rootDomainName string) (iface.RootDomain, error) {

	rootDomain := &RootDomain{}

	err := GetApp().RecordQuery(EnvRootDomains).
		AndWhere(dbx.NewExp("LOWER(domain)={:domain}", dbx.Params{
			"domain": strings.ToLower(rootDomainName), // case insensitive match
		})).
		Limit(1).
		One(rootDomain)

	if err != nil {
		return nil, err
	}

	return rootDomain, nil

}
