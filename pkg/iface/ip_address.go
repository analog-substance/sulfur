package iface

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"time"
)

type IPAddress interface {
	Save() error
	Created() types.DateTime
	Updated() types.DateTime
	Address() string
	SetAddress(val string)
	Is6() bool
	SetIs6(val bool)
	IsGlobalUnicast() bool
	SetIsGlobalUnicast(val bool)
	IsInterfaceLocalMulticast() bool
	SetIsInterfaceLocalMulticast(val bool)
	IsLinkLocalMulticast() bool
	SetIsLinkLocalMulticast(val bool)
	IsLinkLocalUnicast() bool
	SetIsLinkLocalUnicast(val bool)
	IsLoopback() bool
	SetIsLoopback(val bool)
	IsMulticast() bool
	SetIsMulticast(val bool)
	IsPrivate() bool
	SetIsPrivate(val bool)
	IsUnspecified() bool
	SetIsUnspecified(val bool)
	IsShared() bool
	SetIsShared(val bool)
	IsEphemeral() bool
	SetIsEphemeral(val bool)
	LastSeen() types.DateTime
	SetLastSeen(lastSeen time.Time)
	LastSimplePortScan() types.DateTime
	SetLastSimplePortScan(lastScan time.Time)
	ProxyRecord() *core.Record

	GetDomains() []string
	GetDNSRecords() ([]DNSRecord, error)
}
