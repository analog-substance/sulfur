package model

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/filesystem"
	"log"
)

const ArtifactCollection = "artifacts"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*Artifact)(nil)

type Artifact struct {
	SulfurRecordProxy
}

func (a *Artifact) IPAddress() string {
	return a.GetString("ip_address")
}
func (a *Artifact) SetIPAddress(val string) {
	a.Set("ip_address", val)
}

func (a *Artifact) SetArtifacts(val []*filesystem.File) {
	a.Set("artifacts", val)
}

func (a *Artifact) DNSRecord() string {
	return a.GetString("dns_record")
}
func (a *Artifact) SetDNSRecord(val string) {
	a.Set("dns_record", val)
}

func (a *Artifact) GetIP() iface.IPAddress {

	coreRecord, err := FindRecordByID(IPAddressesCollection, a.IPAddress())
	if err != nil {
		log.Println("failed to get IP", err)
		return nil
	}

	// load into proxy
	obj := &IPAddress{}
	obj.SetProxyRecord(coreRecord)

	return obj
}

func (a *Artifact) GetDNSRecord() iface.DNSRecord {

	coreRecord, err := FindRecordByID(DNSRecordCollection, a.DNSRecord())
	if err != nil {
		log.Println("failed to get dns record", err)
		return nil
	}

	// load into proxy
	obj := &DNSRecord{}
	obj.SetProxyRecord(coreRecord)

	return obj
}

func ArtifactFirstOrCreate(dnsRecordId string, ipAddrId string) (iface.Artifact, error) {
	artifact := &Artifact{}

	record, err := FirstOrCreateByFilter(
		ArtifactCollection,
		//"name={:name} && value={:value} && type={:type}",
		"ip_address={:ip_address} AND dns_record={:dns_record}",

		dbx.Params{
			"dns_record": dnsRecordId,
			"ip_address": ipAddrId,
		})

	if err != nil {
		return nil, err
	}
	artifact.SetProxyRecord(record)
	return artifact, nil
}

func GetDNSRecordsNeedingArtifacts() ([]iface.DNSRecord, error) {
	ret := []iface.DNSRecord{}

	var result []dbId

	err := app_state.GetApp().DB().
		Select("*").
		From("bad_certificates").
		AndWhere(dbx.NewExp("artifact_id IS NULL")).
		All(&result)

	idSlice := []interface{}{}
	for _, record := range result {
		idSlice = append(idSlice, record.Id)
	}

	records, err := app_state.GetApp().FindAllRecords(DNSRecordCollection,
		dbx.In("id", idSlice...),
	)

	for _, record := range records {

		// load into proxy
		proxyRec := &DNSRecord{}
		proxyRec.SetProxyRecord(record)

		ret = append(ret, proxyRec)
	}

	return ret, err

}
