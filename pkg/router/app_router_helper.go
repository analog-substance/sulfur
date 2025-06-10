package router

import (
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/pocketbase/pocketbase/core"
)

func AttachRoutes(se *core.ServeEvent) error {

	se.Router.POST(sulfur.ConsumeDNSRecordsPath, consumeDNSRecord)
	se.Router.POST("/consume/{org_id}/assets/root_domains", consumeAssetRootDomains)

	return se.Next()
}
