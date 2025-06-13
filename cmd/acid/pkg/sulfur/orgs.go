package sulfur

import "github.com/analog-substance/sulfur/pkg/sulfur"

func (a *APIClient) ListOrganizations() (*sulfur.OrganizationListResponse, error) {
	resStruct := sulfur.OrganizationListResponse{}

	err := a.GetStruct(sulfur.OrganizationsPath, &resStruct)
	if err != nil {
		return nil, err
	}

	return &resStruct, nil
}
