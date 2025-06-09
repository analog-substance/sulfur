package router

import (
	"encoding/json"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/pocketbase/pocketbase/core"
	"io"
	"log"
	"net/http"
	"time"
)

type orgRootDomainRequest struct {
	Domain    string `json:"domain"`
	Registrar string `json:"registrar"`
}

func consumeAssetRootDomains(e *core.RequestEvent) error {

	orgID := e.Request.PathValue("org_id")

	org, err := model.FindOrgByID(orgID)
	if err != nil {
		return e.String(http.StatusNotFound, "invalid request")
	}

	domains := []orgRootDomainRequest{}
	jsonBytes, err := io.ReadAll(e.Request.Body)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	err = json.Unmarshal(jsonBytes, &domains)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	for _, domain := range domains {

		rootDomain, err := model.RootDomainFirstOrCreate(domain.Domain)
		if err != nil {
			log.Println("cant find root domain", err)
			continue
		}

		log.Println("Found root domain", rootDomain.DomainName())

		if rootDomain.ProxyRecord().Id == "" {
			err = rootDomain.Save()
			if err != nil {
				log.Println("unable to save root domain", err)
				continue
			}
		}

		orgRootDomain, err := model.OrgRootDomainFirstOrCreate(rootDomain.ProxyRecord().Id, org.ProxyRecord().Id)
		if err != nil {
			log.Println("unable to find or create root domain", rootDomain.ProxyRecord().Id, org.ProxyRecord().Id, err)
			continue
		}

		orgRootDomain.SetRegistrar(domain.Registrar)
		orgRootDomain.SetLastSeen(time.Now())

		err = orgRootDomain.Save()
		if err != nil {
			log.Println("err saving root domain", err)
		}
	}
	return e.String(http.StatusOK, "done")
}
