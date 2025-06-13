package sulfur

import "github.com/analog-substance/sulfur/pkg/sulfur"

func (a *APIClient) ListDNSRecords() (*sulfur.DNSRecordListResponse, error) {
	resStruct := sulfur.DNSRecordListResponse{}

	err := a.GetStruct(sulfur.DNSRecordsPath, &resStruct)
	if err != nil {
		return nil, err
	}

	return &resStruct, nil
}
