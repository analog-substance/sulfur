package sulfur

import (
	"fmt"
	"github.com/analog-substance/sulfur/pkg/sulfur"
)

func (a *APIClient) ListDNSRecords() (*sulfur.DNSRecordListResponse, error) {
	resStruct := sulfur.DNSRecordListResponse{}

	apiPath := fmt.Sprintf("%s?perPage=10000", sulfur.DNSRecordsPath)

	err := a.GetStruct(apiPath, &resStruct)
	if err != nil {
		return nil, err
	}

	return &resStruct, nil
}
