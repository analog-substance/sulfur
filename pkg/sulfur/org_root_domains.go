package sulfur

type OrgRootDomain struct {
	Domain    string `json:"domain"`
	Registrar string `json:"registrar"`
}

type OrgDomain struct {
	CollectionId   string `json:"collectionId"`
	CollectionName string `json:"collectionName"`
	Created        string `json:"created"`
	Description    string `json:"description"`
	Expand         struct {
		Organization Organization `json:"organization"`
		RootDomain   RootDomain   `json:"root_domain"`
	} `json:"expand"`
	Id           string `json:"id"`
	LastSeen     string `json:"last_seen"`
	Organization string `json:"organization"`
	Registrar    string `json:"registrar"`
	RootDomain   string `json:"root_domain"`
	Updated      string `json:"updated"`
}

type OrgDomainsListResponse struct {
	Items      []OrgDomain `json:"items"`
	Page       int         `json:"page"`
	PerPage    int         `json:"perPage"`
	TotalItems int         `json:"totalItems"`
	TotalPages int         `json:"totalPages"`
}
