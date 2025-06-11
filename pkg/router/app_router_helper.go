package router

import (
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func AttachRoutes(se *core.ServeEvent) error {

	se.Router.POST(sulfur.ImportDNSRecordsPath, importDNSRecords).Bind(apis.RequireSuperuserAuth())
	se.Router.POST(sulfur.ImportOrgRootDomainsPath, importOrgRootDomains).Bind(apis.RequireSuperuserAuth())

	return se.Next()
}
