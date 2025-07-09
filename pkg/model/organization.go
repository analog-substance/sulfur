package model

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

const OrganizationCollection = "organizations"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*Organization)(nil)

type Organization struct {
	SulfurRecordProxy
}

func (a *Organization) Name() string {
	return a.GetString("name")
}

func FindOrgByID(orgID string) (iface.Organization, error) {

	rdr := &Organization{}

	record, err := app_state.GetApp().FindRecordById(OrganizationCollection, orgID)
	if err != nil {
		return nil, err
	}

	rdr.SetProxyRecord(record)
	return rdr, nil

}

func (a *Organization) SubdomainTakeovers() ([]sulfur.SubdomainTakeover, error) {
	var results []sulfur.SubdomainTakeover
	err := app_state.GetApp().DB().
		Select("*").
		From("subdomain_takeovers").
		AndWhere(dbx.NewExp("org_id={:org_id}", dbx.Params{"org_id": a.Id()})).
		All(&results)

	return results, err
}

func (a *Organization) BadCertificates() ([]sulfur.BadCertificate, error) {
	var results []sulfur.BadCertificate
	err := app_state.GetApp().DB().
		Select(
			"id",
			"COALESCE(artifact_id, '') AS artifact_id",
			"domain",
			"ip_address",
			"fingerprint",
			"subject",
			"alternative_names",
			"COALESCE(external_id, '') as external_id",
			"COALESCE(external_name, '') as external_name",
			"organization",
			"certificate",
			"root_domain").
		From("bad_certificates").
		AndWhere(dbx.NewExp("organization={:org_id}", dbx.Params{"org_id": a.Id()})).
		All(&results)

	return results, err
}
