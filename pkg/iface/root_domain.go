package iface

type RootDomain interface {
	DomainName() string
	DNSRecords() []DNSRecord
	Save() error
	Id() string
}
