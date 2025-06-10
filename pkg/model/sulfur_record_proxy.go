package model

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

type SulfurRecordProxy struct {
	core.BaseRecordProxy
}

func (a *SulfurRecordProxy) Id() string {
	return a.ProxyRecord().Id
}

func (a *SulfurRecordProxy) Save() error {
	return app_state.GetApp().Save(a)
}

func (a *SulfurRecordProxy) Created() types.DateTime {
	return a.GetDateTime("created")
}

func (a *SulfurRecordProxy) Updated() types.DateTime {
	return a.GetDateTime("updated")
}
