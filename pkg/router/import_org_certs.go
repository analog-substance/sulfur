package router

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/pocketbase/pocketbase/core"
)

func importOrgCertificatesRoute(e *core.RequestEvent) error {

	orgID := e.Request.PathValue("org_id")

	org, err := model.FindOrgByID(orgID)
	if err != nil {
		return e.String(http.StatusNotFound, "invalid request")
	}

	orgCerts := []sulfur.OrgCertificateImport{}
	jsonBytes, err := io.ReadAll(e.Request.Body)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	err = json.Unmarshal(jsonBytes, &orgCerts)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	go importOrgCerts(orgCerts, org)
	return e.String(http.StatusOK, "done")
}

func importOrgCerts(orgCerts []sulfur.OrgCertificateImport, org iface.Organization) {
	logger := app_state.GetApp().Logger().WithGroup("importOrgIPAddresses")
	logger.Info("import org ip addresses started", "count", len(orgCerts))
	for _, orgCert := range orgCerts {

		orgCertRecord, err := model.OrgCertificateFirstOrCreate(orgCert.Serial, org.Id())
		if err != nil {
			logger.Error("unable to find or create org cert", "serial", orgCert.Serial, "orgId", org.Id(), "err", err)
			continue
		}

		if orgCert.ExternalReference != "" {
			extRef, err := model.ExternalReferenceFirstOrCreate(orgCert.ExternalReference)
			if err != nil {
				logger.Error("unable to find or create external reference", "extRef", orgCert.ExternalReference, "err", err)
			} else {
				if err := extRef.Save(); err != nil {
					logger.Error("unable to save external reference", "extRef", orgCert.ExternalReference, "err", err)
				} else {
					orgCertRecord.SetExternalReference(extRef)
				}
			}
		}

		orgCertRecord.SetLastSeen(time.Now())
		orgCertRecord.SetIssuer(orgCert.Issuer)
		orgCertRecord.SetNotAfter(orgCert.NotAfter)
		orgCertRecord.SetNotBefore(orgCert.NotBefore)
		orgCertRecord.SetSubject(orgCert.Subject)
		orgCertRecord.SetSubjectAlternativeNames(strings.Join(orgCert.SubjectAlternativeNames, ","))

		err = orgCertRecord.Save()
		if err != nil {
			logger.Error("err saving org cert", "err", err)
		}
	}

	logger.Info("import org certs complete", "count", len(orgCerts))
}
