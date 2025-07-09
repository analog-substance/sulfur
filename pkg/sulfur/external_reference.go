package sulfur

type ExternalReference struct {
	CollectionId   string `json:"collectionId"`
	CollectionName string `json:"collectionName"`
	Id             string `json:"id"`
	Name           string `json:"name"`
	Value          string `json:"value"`
}

type ExternalReferenceResponse struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"perPage"`
	TotalPages int         `json:"totalPages"`
	TotalItems int         `json:"totalItems"`
	Items      []DNSRecord `json:"items"`
}
