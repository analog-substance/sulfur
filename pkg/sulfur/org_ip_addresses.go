package sulfur

type OrgIPAddress struct {
	CollectionId   string `json:"collectionId"`
	CollectionName string `json:"collectionName"`
	Created        string `json:"created"`
	Expand         struct {
		IpAddress    IPAddress    `json:"ip_address"`
		Organization Organization `json:"organization"`
	} `json:"expand"`
	Id           string `json:"id"`
	IpAddress    string `json:"ip_address"`
	LastSeen     string `json:"last_seen"`
	Organization string `json:"organization"`
	Updated      string `json:"updated"`
}

type OrgIpAddressesListResponse struct {
	Items      []OrgIPAddress `json:"items"`
	Page       int            `json:"page"`
	PerPage    int            `json:"perPage"`
	TotalItems int            `json:"totalItems"`
	TotalPages int            `json:"totalPages"`
}
