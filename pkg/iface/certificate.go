package iface

import (
	"math/big"
	"time"

	"github.com/pocketbase/pocketbase/tools/types"
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
	SetNotBefore(lastSeen time.Time)
	Expires() types.DateTime
	SetExpires(lastSeen time.Time)
	SetSerial(serialNumber *big.Int)
	Id() string
}
