package iface

import (
	"github.com/pocketbase/pocketbase/tools/types"
	"time"
)

type IPPort interface {
	Save() error
	Created() types.DateTime
	Updated() types.DateTime
	IPAddress() string
	SetIPAddress(val string)
	Port() int
	SetPort(val int)
	Banner() string
	SetBanner(val string)
	Service() string
	SetService(val string)
	AppProtocol() string
	SetAppProtocol(val string)
	Protocol() string
	SetProtocol(val string)
	LastSeen() types.DateTime
	SetLastSeen(lastSeen time.Time)
	Id() string

	GetIP() IPAddress
}
