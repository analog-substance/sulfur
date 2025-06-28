package model

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"log"
	"net"
	"strings"
	"time"
)

const DNSRecordCollection = "dns_records"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*DNSRecord)(nil)

var sharedSpace *net.IPNet

func init() {
	_, sharedSpace, _ = net.ParseCIDR("100.64.0.0/10")

}

type DNSRecord struct {
	SulfurRecordProxy
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

func (a *DNSRecord) SetRootDomain(domain iface.RootDomain) {
	a.Set("root_domain", domain.Id())
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

	recordValue = strings.TrimSpace(recordValue)
	recordName = strings.TrimSpace(recordName)
	recordName = strings.TrimRight(recordName, ".")
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

	rootDomain, err := RootDomainFirstOrCreate(strings.ToLower(recordName))
	if err == nil {
		if rootDomain.Id() == "" {
			err = rootDomain.Save()
			if err != nil {
				log.Println("unable to save root domain", err)
			}
		}

		dnsR.SetRootDomain(rootDomain)
	} else {
		log.Println("unable to find or create root domain record", err)
	}

	parsedIP := net.ParseIP(recordValue)
	if parsedIP != nil {
		ipRecord, err := IPAddressFirstOrCreate(parsedIP.String())
		if err != nil {
			log.Println("failed to create ip addr", err)
		}

		ipRecord.SetIs6(parsedIP.To4() == nil)
		// not sure what i was thinking whn i created this field....
		//ipRecord.SetIsEphemeral(parsedIP.IsEp)
		ipRecord.SetIsGlobalUnicast(parsedIP.IsGlobalUnicast())
		ipRecord.SetIsInterfaceLocalMulticast(parsedIP.IsLinkLocalMulticast())
		ipRecord.SetIsLoopback(parsedIP.IsLoopback())
		ipRecord.SetIsLinkLocalMulticast(parsedIP.IsLinkLocalMulticast())
		ipRecord.SetIsLinkLocalUnicast(parsedIP.IsLinkLocalUnicast())
		ipRecord.SetIsMulticast(parsedIP.IsMulticast())
		ipRecord.SetIsPrivate(parsedIP.IsPrivate())
		ipRecord.SetIsShared(sharedSpace.Contains(parsedIP))
		ipRecord.SetIsUnspecified(parsedIP.IsUnspecified())

		if err := ipRecord.Save(); err != nil {
			log.Println("failed to save IP", err)
		}
	}

	return dnsR, nil
}

func GetDNSRecordsForIP(ipAddress string) ([]iface.DNSRecord, error) {

	ret := []iface.DNSRecord{}
	// retrieve multiple "articles" records with optional dbx expressions
	records, err := app_state.GetApp().FindAllRecords(DNSRecordCollection,
		dbx.NewExp("value = {:ip_address}", dbx.Params{"ip_address": ipAddress}),
	)

	if err != nil {
		return nil, err
	}
	for _, record := range records {

		// load into proxy
		dnsRecord := &DNSRecord{}
		dnsRecord.SetProxyRecord(record)

		ret = append(ret, dnsRecord)

	}

	return ret, nil
}

type DNSScope struct {
	Host string `db:"host" json:"host"`
}

const SQLDNSLookupQueue = `
SELECT dns_records.name host
FROM dns_records
WHERE dns_records.name NOT IN (
	SELECT dns_records.name
	FROM dns_records
	WHERE (
		(last_resolved > datetime('now', '-4 hours'))
		OR (resolve_error != '')
	)
	GROUP BY dns_records.name
) AND  (
		(last_resolved IS NULL OR last_resolved < datetime('now', '-4 hours'))
		OR (resolve_error IS NOT NULL OR resolve_error != '')
	)
GROUP BY dns_records.name
`

const SQLAddrResolveQueue = `
SELECT dns_records.value host
FROM dns_records
WHERE dns_records.value NOT IN (
	SELECT dns_records.value
	FROM dns_records
	WHERE type = 'A' AND (
		last_resolved > datetime('now', '-4 hours')
	)
	GROUP BY dns_records.value
) AND type = 'A' AND (
		(last_resolved IS NULL OR last_resolved < datetime('now', '-4 hours'))
		OR (resolve_error IS NOT NULL AND resolve_error != '')
	)
GROUP BY dns_records.value
`

func GetDNSLookupQueue() ([]DNSScope, error) {
	accounts := []DNSScope{}
	err := app_state.GetApp().DB().
		NewQuery(SQLDNSLookupQueue).
		All(&accounts)

	return accounts, err
}

func GetActiveIPs() ([]DNSScope, error) {
	accounts := []DNSScope{}
	err := app_state.GetApp().DB().
		NewQuery(SQLAddrResolveQueue).
		All(&accounts)

	return accounts, err
}

const SQLPortScanQueue = `
SELECT ip_addresses.address host
FROM ip_addresses
LEFT JOIN ip_ports ON ip_addresses.id=ip_ports.ip_address
WHERE ip_addresses.is_private = false AND ip_addresses.is_shared = false AND ip_addresses.is_loopback = false
  AND (ip_addresses.last_simple_port_scan IS NULL OR ip_addresses.last_simple_port_scan < datetime('now', '-4 hours'))
GROUP BY ip_addresses.address
`

func GetSimplePortScanInput() ([]DNSScope, error) {
	accounts := []DNSScope{}
	err := app_state.GetApp().DB().
		NewQuery(SQLPortScanQueue).
		All(&accounts)

	return accounts, err
}
