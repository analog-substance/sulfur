package sulfur

type IPAddress struct {
	Address                   string `json:"address"`
	CollectionId              string `json:"collectionId"`
	CollectionName            string `json:"collectionName"`
	Created                   string `json:"created"`
	Id                        string `json:"id"`
	Is6                       bool   `json:"is_6"`
	IsEphemeral               bool   `json:"is_ephemeral"`
	IsGlobalUnicast           bool   `json:"is_global_unicast"`
	IsInterfaceLocalMulticast bool   `json:"is_interface_local_multicast"`
	IsLinkLocalMulticast      bool   `json:"is_link_local_multicast"`
	IsLinkLocalUnicast        bool   `json:"is_link_local_unicast"`
	IsLoopback                bool   `json:"is_loopback"`
	IsMulticast               bool   `json:"is_multicast"`
	IsPrivate                 bool   `json:"is_private"`
	IsShared                  bool   `json:"is_shared"`
	IsUnspecified             bool   `json:"is_unspecified"`
	LastSeen                  string `json:"last_seen"`
	LastSimplePortScan        string `json:"last_simple_port_scan"`
	Updated                   string `json:"updated"`
}
