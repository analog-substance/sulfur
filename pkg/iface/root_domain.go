package iface

type RootDomain interface {
	Registrar() string
	DomainName() string
	//RootDomain() RootDomain
	DNSRecords() []DNSRecord
	//SubDomains() []RootDomain
}
