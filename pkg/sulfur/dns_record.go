package sulfur

type DNSRecord struct {
	CollectionId      string `json:"collectionId"`
	CollectionName    string `json:"collectionName"`
	Id                string `json:"id"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Value             string `json:"value"`
	TTL               int    `json:"ttl"`
	ResolveError      string `json:"resolve_error"`
	ResolveErrorCount int    `json:"resolve_error_count"`
	LastResolved      string `json:"last_resolved"`
	LastSeen          string `json:"last_seen"`
	RootDomain        string `json:"root_domain"`
	Created           string `json:"created"`
	Updated           string `json:"updated"`
	Expand            struct {
		RootDomain RootDomain `json:"root_domain"`
	} `json:"expand"`
}

type DNSRecordListResponse struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"perPage"`
	TotalPages int         `json:"totalPages"`
	TotalItems int         `json:"totalItems"`
	Items      []DNSRecord `json:"items"`
}
