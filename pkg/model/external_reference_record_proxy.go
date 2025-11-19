package model

import (
	"github.com/analog-substance/sulfur/pkg/iface"
)

type ExternalReferenceRecordProxy struct {
	SulfurRecordProxy
}

func (a *ExternalReferenceRecordProxy) SetExternalReference(extRef iface.ExternalReference) {
	a.Set("external_reference", extRef.Id())
}

func (a *ExternalReferenceRecordProxy) ExternalReferenceId() string {
	return a.GetString("external_reference")
}
