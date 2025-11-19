package model

import (
	"log"
	"time"

	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

const IPAddressesCollection = "ip_addresses"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*IPAddress)(nil)

type IPAddress struct {
	LastSeenRecordProxy
}

func (a *IPAddress) Address() string {
	return a.GetString("address")
}
func (a *IPAddress) SetAddress(val string) {
	a.Set("address", val)
}

func (a *IPAddress) Is6() bool {
	return a.GetBool("is_6")
}

func (a *IPAddress) SetIs6(val bool) {
	a.Set("is_6", val)
}

func (a *IPAddress) IsGlobalUnicast() bool {
	return a.GetBool("is_global_unicast")
}

func (a *IPAddress) SetIsGlobalUnicast(val bool) {
	a.Set("is_global_unicast", val)
}

func (a *IPAddress) IsInterfaceLocalMulticast() bool {
	return a.GetBool("is_interface_local_multicast")
}

func (a *IPAddress) SetIsInterfaceLocalMulticast(val bool) {
	a.Set("is_interface_local_multicast", val)
}

func (a *IPAddress) IsLinkLocalMulticast() bool {
	return a.GetBool("is_link_local_multicast")
}

func (a *IPAddress) SetIsLinkLocalMulticast(val bool) {
	a.Set("is_link_local_multicast", val)
}

func (a *IPAddress) IsLinkLocalUnicast() bool {
	return a.GetBool("is_link_local_unicast")
}

func (a *IPAddress) SetIsLinkLocalUnicast(val bool) {
	a.Set("is_link_local_unicast", val)
}

func (a *IPAddress) IsLoopback() bool {
	return a.GetBool("is_loopback")
}

func (a *IPAddress) SetIsLoopback(val bool) {
	a.Set("is_loopback", val)
}

func (a *IPAddress) IsMulticast() bool {
	return a.GetBool("is_multicast")
}

func (a *IPAddress) SetIsMulticast(val bool) {
	a.Set("is_multicast", val)
}

func (a *IPAddress) IsPrivate() bool {
	return a.GetBool("is_private")
}

func (a *IPAddress) SetIsPrivate(val bool) {
	a.Set("is_private", val)
}

func (a *IPAddress) IsUnspecified() bool {
	return a.GetBool("is_unspecified")
}

func (a *IPAddress) SetIsUnspecified(val bool) {
	a.Set("is_unspecified", val)
}

func (a *IPAddress) IsShared() bool {
	return a.GetBool("is_shared")
}

func (a *IPAddress) SetIsShared(val bool) {
	a.Set("is_shared", val)
}

func (a *IPAddress) IsEphemeral() bool {
	return a.GetBool("is_ephemeral")
}

func (a *IPAddress) SetIsEphemeral(val bool) {
	a.Set("is_ephemeral", val)
}

func (a *IPAddress) LastSimplePortScan() types.DateTime {
	return a.GetDateTime("last_simple_port_scan")
}

func (a *IPAddress) SetLastSimplePortScan(lastScan time.Time) {
	a.Set("last_simple_port_scan", lastScan)
}

func (a *IPAddress) GetDomains() []string {

	ret := []string{}
	dnsRecords, err := a.GetDNSRecords()
	if err != nil {
		log.Println("Error getting DNS records:", err)
		return ret
	}

	for _, dnsRecord := range dnsRecords {
		ret = append(ret, dnsRecord.Name())
	}

	return ret
}

func (a *IPAddress) GetDNSRecords() ([]iface.DNSRecord, error) {

	return GetDNSRecordsForIP(a.Address())

}

func IPAddressFirstOrCreate(ipAddr string) (iface.IPAddress, error) {
	ipAddrRecord := &IPAddress{}
	record, err := FirstOrCreateByFilter(
		IPAddressesCollection,
		//"name={:name} && value={:value} && type={:type}",
		"address={:address}",

		dbx.Params{
			"address": ipAddr,
		})

	if err != nil {
		return nil, err
	}
	ipAddrRecord.SetProxyRecord(record)
	return ipAddrRecord, nil
}
