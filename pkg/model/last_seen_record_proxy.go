package model

import (
	"time"

	"github.com/pocketbase/pocketbase/tools/types"
)

type LastSeenRecordProxy struct {
	SulfurRecordProxy
}

func (a *LastSeenRecordProxy) LastSeen() types.DateTime {
	return a.GetDateTime("last_seen")
}

func (a *LastSeenRecordProxy) SetLastSeen(lastSeen time.Time) {
	a.Set("last_seen", lastSeen)
}

type LastSeenExternalReferenceRecordProxy struct {
	ExternalReferenceRecordProxy
}

func (a *LastSeenExternalReferenceRecordProxy) LastSeen() types.DateTime {
	return a.GetDateTime("last_seen")
}

func (a *LastSeenExternalReferenceRecordProxy) SetLastSeen(lastSeen time.Time) {
	a.Set("last_seen", lastSeen)
}

func (a *LastSeenExternalReferenceRecordProxy) Organization() string {
	return a.GetString("organization")
}

func (a *LastSeenExternalReferenceRecordProxy) SetOrganization(org string) {
	a.Set("organization", org)
}
