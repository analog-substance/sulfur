package iface

import "github.com/pocketbase/pocketbase/core"

type Organization interface {
	Name() string
	Save() error
	ProxyRecord() *core.Record
}
