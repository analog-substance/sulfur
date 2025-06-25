package iface

import (
	"github.com/pocketbase/pocketbase/tools/types"
	"time"
)

type OrgIPAddress interface {
	Organization() string
	SetOrganization(string)
	IPAddress() string
	SetIPAddress(string)
	LastSeen() types.DateTime
	SetLastSeen(lastSeen time.Time)

	//SubDomains() []OrgRootDomain
	Save() error
	Id() string
}
