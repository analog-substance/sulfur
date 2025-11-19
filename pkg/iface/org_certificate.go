package iface

import (
	"time"

	"github.com/pocketbase/pocketbase/tools/types"
)

type OrgCertificate interface {
	Organization() string
	SetOrganization(string)
	Serial() string
	SetSerial(string)
	Issuer() string
	SetIssuer(string)
	NotBefore() types.DateTime
	SetNotBefore(time time.Time)
	NotAfter() types.DateTime
	SetNotAfter(time time.Time)

	Subject() string
	SetSubject(string)
	SubjectAlternativeNames() string
	SetSubjectAlternativeNames(string)

	SetExternalReference(extRef ExternalReference)
	ExternalReferenceId() string

	LastSeen() types.DateTime
	SetLastSeen(lastSeen time.Time)
	Save() error
	Id() string
}
