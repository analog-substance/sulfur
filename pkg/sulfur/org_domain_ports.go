package sulfur

type OrgDomainPort struct {
	CollectionId   string `json:"collectionId"`
	CollectionName string `json:"collectionName"`
	Id             string `json:"id"`
	Organization   string `json:"organization"`
	Domain         string `json:"domain"`
	Port           int    `json:"port"`
	Service        string `json:"service"`
	LastSeen       string `json:"last_seen"`
}

type OrgDomainPortListResponse struct {
	Page       int             `json:"page"`
	PerPage    int             `json:"perPage"`
	TotalPages int             `json:"totalPages"`
	TotalItems int             `json:"totalItems"`
	Items      []OrgDomainPort `json:"items"`
}
