package sulfur

import (
	"encoding/json"
	"fmt"
	"github.com/analog-substance/sulfur/pkg/sulfur"
)

func (a *APIClient) ListExternalReferences() (*sulfur.ExternalReferenceResponse, error) {
	resStruct := sulfur.ExternalReferenceResponse{}

	apiPath := fmt.Sprintf("%s?perPage=10000", sulfur.ExternalReferencesPath)

	err := a.GetStruct(apiPath, &resStruct)
	if err != nil {
		return nil, err
	}

	return &resStruct, nil
}

func (a *APIClient) ImportExternalReferences(externalReferencesToImport []sulfur.ExternalReference) error {
	body, err := json.Marshal(externalReferencesToImport)
	if err != nil {
		return err
	}

	resp, err := a.PostJSON(sulfur.ImportExternalReferencesPath, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
