package router

import (
	"encoding/json"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/pocketbase/pocketbase/core"
	"io"
	"log"
	"net/http"
)

type assetRootDomain struct {
	Domain    string `json:"domain"`
	Registrar string `json:"registrar"`
}

func consumeAssetRootDomains(e *core.RequestEvent) error {

	orgID := e.Request.PathValue("org_id")

	org, err := model.FindOrgByID(orgID)
	if err != nil {
		return e.String(http.StatusNotFound, "invalid request")
	}

	domains := []assetRootDomain{}
	jsonBytes, err := io.ReadAll(e.Request.Body)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	err = json.Unmarshal(jsonBytes, &domains)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	for _, domain := range domains {

		rootDomain, err := model.AssetRootDomainFirstOrCreate(domain.Domain, org.ProxyRecord().Id)
		if err != nil {
			log.Println("unable to find or create root domain", err)
			continue
		}
		//rootDomain.Set("domain", domain.Domain)

		err = rootDomain.Save()
		if err != nil {
			log.Println("err saving root domain", err)
		}
	}
	return e.String(http.StatusOK, "done")
}
