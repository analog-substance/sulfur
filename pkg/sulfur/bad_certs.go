package sulfur

type BadCertsListResponse struct {
	Page       int              `json:"page"`
	PerPage    int              `json:"perPage"`
	TotalPages int              `json:"totalPages"`
	TotalItems int              `json:"totalItems"`
	Items      []BadCertificate `json:"items"`
}

type BadCertificate struct {
	CollectionId     string `json:"collectionId,omitempty"`
	CollectionName   string `json:"collectionName,omitempty"`
	Id               string `json:"id"`
	ArtifactId       string `json:"artifact_id"`
	Domain           string `json:"domain"`
	IpAddress        string `json:"ip_address"`
	Fingerprint      string `json:"fingerprint"`
	Subject          string `json:"subject"`
	AlternativeNames string `json:"alternative_names"`
	ExternalId       string `json:"external_id,omitempty"`
	ExternalName     string `json:"external_name,omitempty"`
	Organization     string `json:"organization"`
	Certificate      string `json:"certificate"`
	RootDomain       string `json:"root_domain"`
}
