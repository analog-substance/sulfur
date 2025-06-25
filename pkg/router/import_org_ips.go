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

func importOrgIPAddresses(e *core.RequestEvent) error {

	orgID := e.Request.PathValue("org_id")

	org, err := model.FindOrgByID(orgID)
	if err != nil {
		return e.String(http.StatusNotFound, "invalid request")
	}

	ipAddresses := []string{}
	jsonBytes, err := io.ReadAll(e.Request.Body)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	err = json.Unmarshal(jsonBytes, &ipAddresses)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	for _, ipAddress := range ipAddresses {

		log.Println(ipAddress)

		ipAddr, err := model.IPAddressFirstOrCreate(ipAddress)
		if err != nil {
			log.Println("cant find ip address", err)
			continue
		}

		log.Println("Found ip address", ipAddr.Address())

		if ipAddr.Id() == "" {
			err = ipAddr.Save()
			if err != nil {
				log.Println("unable to save ip address", err)
				continue
			}
		}

		orgIPAddr, err := model.OrgIPAddressFirstOrCreate(ipAddr.Id(), org.Id())
		if err != nil {
			log.Println("unable to find or create org ip addr", ipAddr.Id(), org.Id(), err)
			continue
		}

		orgIPAddr.SetLastSeen(time.Now())

		err = orgIPAddr.Save()
		if err != nil {
			log.Println("err saving org ip addr", err)
		}
	}
	return e.String(http.StatusOK, "done")
}
