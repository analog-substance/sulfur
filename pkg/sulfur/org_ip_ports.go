package sulfur

type OrgIPPort struct {
	CollectionId   string `json:"collectionId"`
	CollectionName string `json:"collectionName"`
	Id             string `json:"id"`
	Organization   string `json:"organization"`
	Address        string `json:"address"`
	Port           int    `json:"port"`
	LastSeen       string `json:"last_seen"`
	DnsNames       string `json:"dns_names"`
}

type OrgIPPortListResponse struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"perPage"`
	TotalPages int         `json:"totalPages"`
	TotalItems int         `json:"totalItems"`
	Items      []OrgIPPort `json:"items"`
}
