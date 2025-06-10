package router

import (
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/pocketbase/pocketbase/core"
)

func AttachRoutes(se *core.ServeEvent) error {

	se.Router.POST(sulfur.ImportDNSRecordsPath, importDNSRecords)
	se.Router.POST(sulfur.ImportOrgRootDomainsPath, importOrgRootDomains)

	return se.Next()
}
