package sulfur

type SubdomainTakeover struct {
	CollectionId   string `json:"collectionId"`
	CollectionName string `json:"collectionName"`
	Id             string `json:"id"`
	OrgId          string `json:"org_id"`
	OrgName        string `json:"org_name"`
	Domain         string `json:"domain"`
	IpAddress      string `json:"ip_address"`
	LastResolved   string `json:"last_resolved"`
}

type SubdomainTakeoversListResponse struct {
	Page       int                 `json:"page"`
	PerPage    int                 `json:"perPage"`
	TotalPages int                 `json:"totalPages"`
	TotalItems int                 `json:"totalItems"`
	Items      []SubdomainTakeover `json:"items"`
}
