package iface

import (
	"github.com/pocketbase/pocketbase/tools/types"
	"time"
)

type OrgRootDomain interface {
	Registrar() string
	SetRegistrar(string)
	DomainName() string
	DNSRecords() []DNSRecord
	LastSeen() types.DateTime
	SetLastSeen(lastSeen time.Time)

	//SubDomains() []OrgRootDomain
	Save() error
	Id() string
}
