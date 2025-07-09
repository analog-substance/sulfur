package router

import (
	"encoding/json"
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/pocketbase/pocketbase/core"
	"io"
	"net/http"
)

func importExternalRefsRoute(e *core.RequestEvent) error {

	var extRefs = []sulfur.ExternalReference{}
	jsonBytes, err := io.ReadAll(e.Request.Body)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	err = json.Unmarshal(jsonBytes, &extRefs)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	go importExternalRefs(extRefs)
	return e.String(http.StatusOK, "done")
}

func importExternalRefs(extRefs []sulfur.ExternalReference) {
	logger := app_state.GetApp().Logger().WithGroup("importExternalRefs")
	logger.Info("import external references", "count", len(extRefs))
	for _, newExtRef := range extRefs {
		extRef, err := model.ExternalReferenceFirstOrCreate(newExtRef.Value)
		if err != nil {
			logger.Error("unable to find or create external reference", "extRef", newExtRef.Value, "err", err)
			continue
		}
		extRef.SetName(newExtRef.Name)
		if err := extRef.Save(); err != nil {
			logger.Error("unable to save external reference", "extRef", newExtRef.Value, "err", err)
		}
	}

	logger.Info("import org ip addresses complete", "count", len(extRefs))
}
