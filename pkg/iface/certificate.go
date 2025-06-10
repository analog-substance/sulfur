package iface

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"time"
)

type Certificate interface {
	Save() error
	Created() types.DateTime
	Updated() types.DateTime
	Fingerprint() string
	SetFingerprint(val string)
	Issuer() string
	SetIssuer(val string)
	Subject() string
	SetSubject(val string)
	AlternativeNames() string
	SetAlternativeNames(val string)
	Issued() types.DateTime
	SetIssued(lastSeen time.Time)
	Expires() types.DateTime
	SetExpires(lastSeen time.Time)
	ProxyRecord() *core.Record
}
