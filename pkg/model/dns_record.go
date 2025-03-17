package model

import (
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"time"
)

const DNSRecords = "dns_records"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*DNSRecord)(nil)

type DNSRecord struct {
	core.BaseRecordProxy
}

func (a *DNSRecord) Save() error {
	return GetApp().Save(a)
}

func (a *DNSRecord) RootDomain() iface.RootDomain {
	return &RootDomain{}
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
	record, err := FirstOrCreateByFilter(
		DNSRecords,
		"name={:name} && value={:value} && type={:type}",
		dbx.Params{
			"name":  recordName,  // case insensitive match
			"value": recordValue, // case insensitive match
			"type":  recordType,
		})

	if err != nil {
		return nil, err
	}
	dnsR.SetProxyRecord(record)
	return dnsR, nil
}
