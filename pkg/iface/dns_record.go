package iface

import (
	"github.com/pocketbase/pocketbase/tools/types"
	"time"
)

type DNSRecord interface {
	SetName(string)
	SetValue(string)
	SetType(string)
	SetTTL(time.Duration)

	SetResolveErr(resolveErr string)
	SetResolveErrCount(errCount int)
	SetLastResolved(lastResolved time.Time)
	SetLastSeen(lastSeen time.Time)

	Name() string
	Value() string
	Type() string
	TTL() time.Duration
	RootDomain() RootDomain
	ResolveError() string
	ResolveErrorCount() string
	LastResolved() types.DateTime
	LastSeen() types.DateTime
	Created() types.DateTime
	Updated() types.DateTime
	Save() error
}
