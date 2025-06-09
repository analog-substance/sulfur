package model

import (
	"net"
)

type CIDR struct {
	CollectionId              string `json:"collectionId"`
	CollectionName            string `json:"collectionName"`
	ID                        string `json:"id"`
	Organization              string `json:"organization"`
	Cidr                      string `json:"cidr"`
	Asn                       int    `json:"asn"`
	Is6                       bool   `json:"is_6"`
	IsGlobalUnicast           bool   `json:"is_global_unicast"`
	IsInterfaceLocalMulticast bool   `json:"is_interface_local_multicast"`
	IsLinkLocalMulticast      bool   `json:"is_link_local_multicast"`
	IsLinkLocalUnicast        bool   `json:"is_link_local_unicast"`
	IsLoopback                bool   `json:"is_loopback"`
	IsMulticast               bool   `json:"is_multicast"`
	IsPrivate                 bool   `json:"is_private"`
	IsUnspecified             bool   `json:"is_unspecified"`
	IsShared                  bool   `json:"is_shared"`
	IsEphemeral               bool   `json:"is_ephemeral"`
	LastSeen                  string `json:"last_seen"`
	Created                   string `json:"created"`
	Updated                   string `json:"updated"`
}

func CIDRFromNetCIDR(c *net.IPNet) *CIDR {
	ci := CIDRFromNetIP(&c.IP)
	ci.Cidr = c.String()
	return ci
}

func CIDRFromNetIP(c *net.IP) *CIDR {

	is6 := c.To4() == nil

	ci := &CIDR{
		Cidr:                      c.String(),
		IsGlobalUnicast:           c.IsGlobalUnicast(),
		IsInterfaceLocalMulticast: c.IsInterfaceLocalMulticast(),
		IsLinkLocalMulticast:      c.IsLinkLocalMulticast(),
		IsLinkLocalUnicast:        c.IsLinkLocalUnicast(),
		IsLoopback:                c.IsLoopback(),
		IsMulticast:               c.IsMulticast(),
		IsPrivate:                 c.IsPrivate(),
		IsUnspecified:             c.IsUnspecified(),
		Is6:                       is6,
	}

	return ci
}
