package iface

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"time"
)

type IPPortCertificate interface {
	Save() error
	Created() types.DateTime
	Updated() types.DateTime
	IPPort() string
	SetIPPort(val string)
	Certificate() string
	SetCertificate(val string)
	LastSeen() types.DateTime
	SetLastSeen(lastSeen time.Time)
	ProxyRecord() *core.Record
}
