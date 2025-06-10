package sulfur

type DNSRecord struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
	TTL   int    `json:"ttl"`
}
