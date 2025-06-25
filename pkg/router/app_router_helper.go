package router

import (
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func AttachRoutes(se *core.ServeEvent) error {

	se.Router.GET(sulfur.ExportOrgSubdomainTakeovers, exportOrgSubdomainTakeovers).Bind(apis.RequireSuperuserAuth())
	se.Router.POST(sulfur.ImportDNSRecordsPath, importDNSRecords).Bind(apis.RequireSuperuserAuth())
	se.Router.POST(sulfur.ImportDomainAndResolvePath, importDomainsAndResolve).Bind(apis.RequireSuperuserAuth())
	se.Router.POST(sulfur.ImportOrgIPAddressesPath, importOrgIPAddresses).Bind(apis.RequireSuperuserAuth())
	se.Router.POST(sulfur.ImportOrgRootDomainsPath, importOrgRootDomains).Bind(apis.RequireSuperuserAuth())

	return se.Next()
}
