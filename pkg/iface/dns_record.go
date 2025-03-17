package iface

import "time"

type DNSRecord interface {
	Name() string
	Value() string
	TTL() time.Duration
	RootDomain() RootDomain
}
