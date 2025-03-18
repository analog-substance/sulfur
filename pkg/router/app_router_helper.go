package router

import (
	"github.com/pocketbase/pocketbase/core"
)

func AttachRoutes(se *core.ServeEvent) error {

	se.Router.POST("/consume/dns_records", consumeDNSRecord)
	se.Router.POST("/consume/{org_id}/assets/root_domains", consumeAssetRootDomains)

	return se.Next()
}
