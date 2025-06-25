package iface

import "github.com/analog-substance/sulfur/pkg/sulfur"

type Organization interface {
	Name() string
	Save() error
	Id() string
	SubdomainTakeovers() ([]sulfur.SubdomainTakeover, error)
}
