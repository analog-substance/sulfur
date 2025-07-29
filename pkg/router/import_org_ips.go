package router

import (
	"encoding/json"
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/pocketbase/pocketbase/core"
	"io"
	"net/http"
	"time"
)

func importOrgIPAddresses(e *core.RequestEvent) error {

	orgID := e.Request.PathValue("org_id")

	org, err := model.FindOrgByID(orgID)
	if err != nil {
		return e.String(http.StatusNotFound, "invalid request")
	}

	ipAddresses := []sulfur.OrgIPAddressImport{}
	jsonBytes, err := io.ReadAll(e.Request.Body)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	err = json.Unmarshal(jsonBytes, &ipAddresses)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	go importOrgIPs(ipAddresses, org)
	return e.String(http.StatusOK, "done")
}

func importOrgIPs(ipAddresses []sulfur.OrgIPAddressImport, org iface.Organization) {
	logger := app_state.GetApp().Logger().WithGroup("importOrgIPAddresses")
	logger.Info("import org ip addresses started", "count", len(ipAddresses))
	for _, ipAddress := range ipAddresses {
		ipAddr, err := model.IPAddressFirstOrCreate(ipAddress.IpAddress)
		if err != nil {
			logger.Error("cant find ip address", "err", err)
			continue
		}

		if ipAddr.Id() == "" {
			err = ipAddr.Save()
			if err != nil {
				logger.Error("cant save ip address", "err", err)
				continue
			}
		}

		orgIPAddr, err := model.OrgIPAddressFirstOrCreate(ipAddr.Id(), org.Id())
		if err != nil {
			logger.Error("unable to find or create org ip addr", "ipAddrId", ipAddr.Id(), "orgId", org.Id(), "err", err)
			continue
		}

		if ipAddress.ExternalReference != "" {
			extRef, err := model.ExternalReferenceFirstOrCreate(ipAddress.ExternalReference)
			if err != nil {
				logger.Error("unable to find or create external reference", "extRef", ipAddress.ExternalReference, "err", err)
			} else {
				if err := extRef.Save(); err != nil {
					logger.Error("unable to save external reference", "extRef", ipAddress.ExternalReference, "err", err)
				} else {
					orgIPAddr.SetExternalReference(extRef)
				}
			}
		}
		orgIPAddr.SetLastSeen(time.Now())

		err = orgIPAddr.Save()
		if err != nil {
			logger.Error("err saving org ip addr", "err", err)
		}
	}

	logger.Info("import org ip addresses complete", "count", len(ipAddresses))
}
