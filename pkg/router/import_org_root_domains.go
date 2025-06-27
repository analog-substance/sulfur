package router

import (
	"encoding/json"
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/pocketbase/pocketbase/core"
	"io"
	"log"
	"net/http"
	"time"
)

func importOrgRootDomains(e *core.RequestEvent) error {

	orgID := e.Request.PathValue("org_id")

	org, err := model.FindOrgByID(orgID)
	if err != nil {
		return e.String(http.StatusNotFound, "invalid request")
	}

	domains := []sulfur.OrgRootDomain{}
	jsonBytes, err := io.ReadAll(e.Request.Body)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	err = json.Unmarshal(jsonBytes, &domains)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	go importOrgDomains(domains, org)
	return e.String(http.StatusOK, "done")
}

func importOrgDomains(domains []sulfur.OrgRootDomain, org iface.Organization) {
	logger := app_state.GetApp().Logger().WithGroup("importOrgDomains")
	logger.Info("Imported org root started", "count", len(domains))

	for _, domain := range domains {
		rootDomain, err := model.RootDomainFirstOrCreate(domain.Domain)
		if err != nil {
			logger.Error("unable to find root domain", "err", err)
			continue
		}

		log.Println("Found root domain", rootDomain.DomainName())

		if rootDomain.Id() == "" {
			err = rootDomain.Save()
			if err != nil {
				logger.Error("unable to save root domain", "err", err)
				continue
			}
		}

		orgRootDomain, err := model.OrgRootDomainFirstOrCreate(rootDomain.Id(), org.Id())
		if err != nil {
			logger.Error("unable to find or create root domain", "rootDomainId", rootDomain.Id(), "orgId", org.Id(), "err", err)
			continue
		}

		orgRootDomain.SetRegistrar(domain.Registrar)
		orgRootDomain.SetLastSeen(time.Now())

		err = orgRootDomain.Save()
		if err != nil {
			logger.Error("error saving org root domain", "err", err)
		}
	}
	logger.Info("Imported org root domains", "count", len(domains))
}
