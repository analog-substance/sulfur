package iface

import "github.com/pocketbase/pocketbase/core"

type AssetRootDomain interface {
	Registrar() string
	DomainName() string
	DNSRecords() []DNSRecord
	//SubDomains() []AssetRootDomain
	Save() error
	ProxyRecord() *core.Record
}
