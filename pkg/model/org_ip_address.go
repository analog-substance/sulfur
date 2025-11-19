package model

import (
	"time"

	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

const OrgIPAddressCollection = "org_ip_addresses"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*OrgIPAddress)(nil)

type OrgIPAddress struct {
	OrganizationRecordProxy
}

func (a *OrgIPAddress) IPAddress() string {
	return a.GetString("ip_address")
}

func (a *OrgIPAddress) LastSeen() types.DateTime {
	return a.GetDateTime("last_seen")
}

func (a *OrgIPAddress) SetIPAddress(ipAddress string) {
	a.Set("ip_address", ipAddress)
}

func (a *OrgIPAddress) SetLastSeen(lastSeen time.Time) {
	a.Set("last_seen", lastSeen)
}

func (a *OrgIPAddress) SetExternalReference(extRef iface.ExternalReference) {
	a.Set("external_reference", extRef.Id())
}

func OrgIPAddressFirstOrCreate(ipAddressId, orgId string) (iface.OrgIPAddress, error) {
	dnsR := &OrgIPAddress{}

	record, err := FirstOrCreateByFilter(
		OrgIPAddressCollection,
		//"name={:name} && value={:value} && type={:type}",
		"ip_address={:ip_address} AND organization={:organization}",

		dbx.Params{
			"ip_address":   ipAddressId,
			"organization": orgId,
		})

	if err != nil {
		return nil, err
	}
	dnsR.SetProxyRecord(record)
	return dnsR, nil
}
