package iface

import (
	"github.com/pocketbase/pocketbase/tools/types"
)

type ExternalReference interface {
	SetName(string)
	Name() string
	Value() string
	SetValue(string)
	SetDescription(string)
	Description() string

	Created() types.DateTime
	Updated() types.DateTime
	Save() error
	Id() string
}
