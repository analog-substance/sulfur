package model

import (
	"database/sql"
	"errors"
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

func (a *DNSRecord) SetTTL(newTTL int) {
	a.Set("ttl", newTTL)
}
func (a *DNSRecord) SetType(recordType string) {
	a.Set("type", recordType)
}

func FindDNSRecord(recordName, recordValue, recordType string) (iface.DNSRecord, error) {

	dnsR := &DNSRecord{}

	record, err := GetApp().FindFirstRecordByFilter(
		DNSRecords,
		"name={:recordName} AND value={:recordValue} AND type={:recordType}",
		dbx.Params{
			"recordName":  recordName,  // case insensitive match
			"recordValue": recordValue, // case insensitive match
			"recordType":  recordType,
		})
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		coll, err := GetApp().FindCollectionByNameOrId(DNSRecords)
		if err != nil {
			return nil, err
		}
		
	}

	return dnsR, nil
}
