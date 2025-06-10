package iface

type Organization interface {
	Name() string
	Save() error
	Id() string
}
