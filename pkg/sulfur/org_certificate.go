package sulfur

import "time"

type OrgCertificateImport struct {
	Serial                  string    `json:"serial"`
	Issuer                  string    `json:"issuer"`
	NotBefore               time.Time `json:"not_before"`
	NotAfter                time.Time `json:"not_after"`
	Subject                 string    `json:"subject"`
	SubjectAlternativeNames []string  `json:"subject_alternative_names"`
	ExternalReference       string    `json:"external_reference"`
	LastSeen                string    `json:"last_seen"`
}

type OrgCertificate struct {
	CollectionId            string `json:"collectionId"`
	CollectionName          string `json:"collectionName"`
	Created                 string `json:"created"`
	Id                      string `json:"id"`
	Serial                  string `json:"serial"`
	Issuer                  string `json:"issuer"`
	NotBefore               string `json:"not_before"`
	NotAfter                string `json:"not_after"`
	Subject                 string `json:"subject"`
	SubjectAlternativeNames string `json:"subject_alternative_names"`
	ExternalReference       string `json:"external_reference"`
	LastSeen                string `json:"last_seen"`
	Organization            string `json:"organization"`
	Updated                 string `json:"updated"`
	Expand                  struct {
		Organization      Organization      `json:"organization"`
		ExternalReference ExternalReference `json:"external_reference"`
	} `json:"expand"`
}

type OrgCertificatesListResponse struct {
	Items      []OrgCertificate `json:"items"`
	Page       int              `json:"page"`
	PerPage    int              `json:"perPage"`
	TotalItems int              `json:"totalItems"`
	TotalPages int              `json:"totalPages"`
}
