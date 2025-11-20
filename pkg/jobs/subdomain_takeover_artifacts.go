package jobs

import (
	"log"

	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/pocketbase/pocketbase/tools/filesystem"
)

func init() {

}

func SubdomainTakeoverArtifacts() {
	logger := app_state.GetApp().Logger().WithGroup("SubdomainTakeoverArtifacts")
	dnsRecords, err := model.GetTakeoversNeedingArtifacts()
	if err != nil {
		logger.Error("Error getting dns records", "error", err)
		return
	}

	total := len(dnsRecords)
	logger.Info("processing results", "count", total)

	for _, record := range dnsRecords {
		logger.Info("processing record", "record", record)
		log.Println("processing record", "record", record)
		screenshot, err := Screenshot(record.Name(), record.Value(), logger)
		if err != nil {
			logger.Error("Error getting screenshot", "record", record, "error", err)
			continue
		}

		artifact, err := model.ArtifactFirstOrCreate(record.Id(), "")
		if err != nil {
			logger.Error("Error creating artifact obj", "record", record, "error", err)
			continue
		}

		f, err := filesystem.NewFileFromBytes(screenshot, "subdomain-takeover-screenshot.png")
		artifact.SetArtifacts([]*filesystem.File{f})

		err = artifact.Save()
		if err != nil {
			logger.Error("Error saving artifact obj", "record", record, "error", err)
		}
	}
}
