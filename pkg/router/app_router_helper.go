package router

import (
	"github.com/pocketbase/pocketbase/core"
)

func AttachRoutes(se *core.ServeEvent) error {

	se.Router.POST("/consume/dns_record", consumeDNSRecord)

	return se.Next()
}
