package iface

import "github.com/pocketbase/pocketbase/core"

type RootDomain interface {
	DomainName() string
	DNSRecords() []DNSRecord
	Save() error
	ProxyRecord() *core.Record
}
