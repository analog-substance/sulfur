package model

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

const OrganizationCollection = "organizations"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*Organization)(nil)

type Organization struct {
	core.BaseRecordProxy
}

func (a *Organization) Save() error {
	return app_state.GetApp().Save(a)
}

func (a *Organization) Name() string {
	return a.GetString("name")
}

func (a *Organization) Created() types.DateTime {
	return a.GetDateTime("created")
}

func (a *Organization) Updated() types.DateTime {
	return a.GetDateTime("updated")
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
