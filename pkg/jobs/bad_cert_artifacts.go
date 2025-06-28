package jobs

import (
	"context"
	"fmt"
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/chromedp/chromedp"
	"github.com/pocketbase/pocketbase/tools/filesystem"
	"log/slog"
	"time"
)

func init() {

}

func BadCertArtifacts() {
	logger := app_state.GetApp().Logger().WithGroup("BadCertArtifacts")
	dnsRecords, err := model.GetDNSRecordsNeedingArtifacts()
	if err != nil {
		logger.Error("Error getting dns records", "error", err)
		return
	}

	total := len(dnsRecords)
	logger.Info("processing results", "count", total)

	for _, record := range dnsRecords {
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

		f, err := filesystem.NewFileFromBytes(screenshot, "bad-cert-screenshot.png")
		artifact.SetArtifacts([]*filesystem.File{f})

		err = artifact.Save()
		if err != nil {
			logger.Error("Error saving artifact obj", "record", record, "error", err)
		}
	}
}

func Screenshot(host, ip string, logger *slog.Logger) ([]byte, error) {
	// create context
	allocatorCtx, allocatorCancel := chromedp.NewExecAllocator(
		context.Background(),
		chromedp.Flag("host-resolver-rules", fmt.Sprintf("MAP %s %s", host, ip)),
		chromedp.Flag("ignore-certificate-errors", "1"),
		chromedp.Flag("headless", true),
	)
	defer allocatorCancel()

	ctx, cancel := chromedp.NewContext(
		allocatorCtx,
		chromedp.IgnoreCertErrors,
		chromedp.Ign,

		//chromedp.WithDebugf(debug.Printf),
	)
	defer cancel()

	var buf []byte
	// capture entire browser viewport, returning png with quality=90
	err := chromedp.Run(ctx, chromedp.Navigate(fmt.Sprintf("https://%s", host)))
	if err != nil {
		logger.Error("Ignoring navigation error", "error", err)
	}

	err = chromedp.Run(ctx, fullScreenshot(fmt.Sprintf("https://%s", host), 90, &buf))
	return buf, err
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Sleep(2 * time.Second),
		chromedp.FullScreenshot(res, quality),
	}
}
