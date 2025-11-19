package model

type OrganizationRecordProxy struct {
	SulfurRecordProxy
}

func (a *OrganizationRecordProxy) Organization() string {
	return a.GetString("organization")
}

func (a *OrganizationRecordProxy) SetOrganization(org string) {
	a.Set("organization", org)
}
