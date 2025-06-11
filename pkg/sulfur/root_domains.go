package sulfur

type RootDomain struct {
	CollectionId   string `json:"collectionId"`
	CollectionName string `json:"collectionName"`
	Id             string `json:"id"`
	Domain         string `json:"domain"`
	LastSeen       string `json:"last_seen"`
	Created        string `json:"created"`
	Updated        string `json:"updated"`
}

type RootDomainsListResponse struct {
	Page       int          `json:"page"`
	PerPage    int          `json:"perPage"`
	TotalPages int          `json:"totalPages"`
	TotalItems int          `json:"totalItems"`
	Items      []RootDomain `json:"items"`
}
