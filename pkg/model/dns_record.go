package model

import (
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"strings"
	"time"
)

const DNSRecordCollection = "dns_records"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*DNSRecord)(nil)

type DNSRecord struct {
	core.BaseRecordProxy
}

func (a *DNSRecord) Save() error {
	return GetApp().Save(a)
}

func (a *DNSRecord) RootDomain() iface.AssetRootDomain {
	return &AssetRootDomain{}
}

func (a *DNSRecord) Name() string {
	return a.GetString("name")
}

func (a *DNSRecord) Value() string {
	return a.GetString("value")
}
func (a *DNSRecord) Type() string {
	return a.GetString("type")
}

func (a *DNSRecord) TTL() time.Duration {
	ttl := a.GetInt("ttl")
	return time.Duration(ttl) * time.Second
}

func (a *DNSRecord) ResolveError() string {
	return a.GetString("resolve_error")
}

func (a *DNSRecord) ResolveErrorCount() string {
	return a.GetString("resolve_error_count")
}

func (a *DNSRecord) LastResolved() types.DateTime {
	return a.GetDateTime("last_resolved")
}

func (a *DNSRecord) LastSeen() types.DateTime {
	return a.GetDateTime("last_seen")
}

func (a *DNSRecord) Created() types.DateTime {
	return a.GetDateTime("created")
}

func (a *DNSRecord) Updated() types.DateTime {
	return a.GetDateTime("updated")
}

func (a *DNSRecord) SetName(name string) {
	a.Set("name", name)
}

func (a *DNSRecord) SetValue(val string) {
	a.Set("value", val)
}

func (a *DNSRecord) SetTTL(newTTL time.Duration) {
	a.Set("ttl", newTTL)
}

func (a *DNSRecord) SetType(recordType string) {
	a.Set("type", recordType)
}

func (a *DNSRecord) SetRootDomain(domain iface.AssetRootDomain) {
	a.Set("root_domain", domain.ProxyRecord().Id)
}

func (a *DNSRecord) SetResolveErr(resolveErr string) {
	a.Set("resolve_error", resolveErr)
}

func (a *DNSRecord) SetResolveErrCount(errCount int) {
	a.Set("resolve_error_count", errCount)
}

func (a *DNSRecord) SetLastResolved(lastResolved time.Time) {
	a.Set("last_resolved", lastResolved)
}

func (a *DNSRecord) SetLastSeen(lastSeen time.Time) {
	a.Set("last_seen", lastSeen)
}

func DNSRecordFirstOrCreate(recordName, recordValue, recordType string) (iface.DNSRecord, error) {
	dnsR := &DNSRecord{}

	//err := GetApp().RecordQuery(DNSRecordCollection).
	//	AndWhere(dbx.NewExp("LOWER(name)={:name} AND value={:value} AND UPPER(type)={:type}", dbx.Params{
	//		"name":  strings.ToLower(recordName),
	//		"value": recordValue,
	//		"type":  strings.ToUpper(recordType),
	//	})).
	//	Limit(1).
	//	One(dnsR)
	//
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//
	//	}
	//}
	record, err := FirstOrCreateByFilter(
		DNSRecordCollection,
		//"name={:name} && value={:value} && type={:type}",
		"LOWER(name)={:name} AND value={:value} AND UPPER(type)={:type}",

		dbx.Params{
			"name":  strings.ToLower(recordName),
			"value": recordValue,
			"type":  strings.ToUpper(recordType),
		})

	if err != nil {
		return nil, err
	}
	dnsR.SetProxyRecord(record)
	return dnsR, nil
}
