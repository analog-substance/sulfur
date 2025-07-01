package model

import (
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

const ExternalReferenceCollection = "external_references"

// ensures that the Article struct satisfy the core.RecordProxy interface
var _ core.RecordProxy = (*ExternalReference)(nil)

type ExternalReference struct {
	SulfurRecordProxy
}

func (a *ExternalReference) Name() string {
	return a.GetString("name")
}
func (a *ExternalReference) SetName(val string) {
	a.Set("name", val)
}

func (a *ExternalReference) Description() string {
	return a.GetString("description")
}
func (a *ExternalReference) SetDescription(val string) {
	a.Set("description", val)
}

func (a *ExternalReference) Value() string {
	return a.GetString("value")
}
func (a *ExternalReference) SetValue(val string) {
	a.Set("value", val)
}

func ExternalReferenceFirstOrCreate(value string) (iface.ExternalReference, error) {
	recordProxy := &ExternalReference{}

	record, err := FirstOrCreateByFilter(
		ExternalReferenceCollection,
		//"name={:name} && value={:value} && type={:type}",
		"value={:value}",

		dbx.Params{
			"value": value,
		})

	if err != nil {
		return nil, err
	}
	recordProxy.SetProxyRecord(record)
	return recordProxy, nil
}
