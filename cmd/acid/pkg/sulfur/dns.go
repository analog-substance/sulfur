package sulfur

import (
	"encoding/json"
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

func (a *APIClient) ImportDNSRecords(domainsToImport []sulfur.DNSRecord) error {
	body, err := json.Marshal(domainsToImport)
	if err != nil {
		return err
	}

	resp, err := a.PostJSON(sulfur.ImportDNSRecordsPath, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
