package sulfur

type Organization struct {
	CollectionId   string `json:"collectionId"`
	CollectionName string `json:"collectionName"`
	Id             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Created        string `json:"created"`
	Updated        string `json:"updated"`
}

type OrganizationListResponse struct {
	Page       int            `json:"page"`
	PerPage    int            `json:"perPage"`
	TotalPages int            `json:"totalPages"`
	TotalItems int            `json:"totalItems"`
	Items      []Organization `json:"items"`
}
