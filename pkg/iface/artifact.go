package iface

import (
	"github.com/pocketbase/pocketbase/tools/filesystem"
	"github.com/pocketbase/pocketbase/tools/types"
)

type Artifact interface {
	Save() error
	Created() types.DateTime
	Updated() types.DateTime
	IPAddress() string
	SetIPAddress(val string)
	SetArtifacts(val []*filesystem.File)
	DNSRecord() string
	SetDNSRecord(val string)
	Id() string

	GetIP() IPAddress
	GetDNSRecord() DNSRecord
}
