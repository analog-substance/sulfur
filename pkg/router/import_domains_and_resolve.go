package router

import (
	"encoding/json"
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/jobs"
	"github.com/pocketbase/pocketbase/core"
	"io"
	"net/http"
)

func importDomainsAndResolve(e *core.RequestEvent) error {
	domains := []string{}
	jsonBytes, err := io.ReadAll(e.Request.Body)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	err = json.Unmarshal(jsonBytes, &domains)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}
	logger := app_state.GetApp().Logger().WithGroup("ImportAndResolveDomains")
	go jobs.ResolveDomains(domains, logger)

	return e.String(http.StatusOK, "done")
}
