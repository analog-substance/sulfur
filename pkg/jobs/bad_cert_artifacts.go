package jobs

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"sync"
	"time"

	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/chromedp/chromedp"
	"github.com/pocketbase/pocketbase/tools/filesystem"
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
		logger.Info("processing record", "record", record)
		log.Println("processing record", "job", "BadCertArtifacts", "record", record)

		parsedIP := net.ParseIP(record.Value())
		if parsedIP == nil {
			logger.Error("Error parsing ip", "ip", record.Value())
			continue
		}

		ipAddr, err := model.IPAddressFirstOrCreate(record.Value())
		if err != nil {
			logger.Error("Error getting ip address", "error", err)
			continue
		}

		screenshot, err := Screenshot(record.Name(), parsedIP.String(), logger)
		if err != nil {
			logger.Error("Error getting screenshot", "record", record, "error", err)
			continue
		}

		artifact, err := model.ArtifactFirstOrCreate(record.Id(), ipAddr.Id())
		if err != nil {
			logger.Error("Error creating artifact obj", "record", record, "error", err)
			continue
		}

		f, err := filesystem.NewFileFromBytes(screenshot, "bad-cert-screenshot.png")
		if err != nil {
			logger.Error("Error creating file from bytes", "record", record, "error", err)
			continue
		}
		artifact.SetArtifacts([]*filesystem.File{f})

		err = artifact.Save()
		if err != nil {
			logger.Error("Error saving artifact obj", "record", record, "error", err)
		}
	}
}

var screenshotMutex sync.Mutex

func Screenshot(host, ip string, logger *slog.Logger) ([]byte, error) {

	//screenshotMutex.Lock()
	//defer screenshotMutex.Unlock()

	//userData, err := os.MkdirTemp("", "sulfur-screenshot-*")
	//if err != nil {
	//	return nil, err
	//}
	//
	//defer func() {
	//	log.Println("Cleanup screenshot")
	//
	//	if err := os.RemoveAll(userData); err != nil {
	//		time.Sleep(3 * time.Second)
	//		if err := os.RemoveAll(userData); err != nil {
	//			logger.Error("failed to delete chrome data dir", "dir", userData, "err", err)
	//		}
	//	}
	//
	//	snapPath := filepath.Join("/tmp/snap-private-tmp/snap.chromium", userData)
	//	if _, err := os.Stat(snapPath); !os.IsNotExist(err) {
	//		if err := os.RemoveAll(snapPath); err != nil {
	//			time.Sleep(3 * time.Second)
	//			if err := os.RemoveAll(snapPath); err != nil {
	//				logger.Error("failed to delete chrome data dir", "dir", snapPath, "err", err)
	//			}
	//		}
	//	}
	//
	//}()

	log.Println("screenshot: new allocator")
	// create context
	allocatorCtx, allocatorCancel := chromedp.NewExecAllocator(
		context.Background(),
		chromedp.Flag("host-resolver-rules", fmt.Sprintf("MAP %s %s", host, ip)),
		chromedp.Flag("ignore-certificate-errors", "1"),
		chromedp.Flag("headless", true),
		//chromedp.UserDataDir(userData),
	)
	defer allocatorCancel()

	log.Println("screenshot: new context")
	ctx, cancel := chromedp.NewContext(
		allocatorCtx,
		//chromedp.WithDebugf(debug.Printf),
	)
	defer cancel()

	var buf []byte
	// capture entire browser viewport, returning png with quality=90
	log.Println("screenshot: navigate")

	err := chromedp.Run(ctx, chromedp.Navigate(fmt.Sprintf("https://%s", host)))
	if err != nil {
		logger.Error("Ignoring navigation error", "error", err)
	}

	log.Println("screenshot: fullScreenshot()")

	err = chromedp.Run(ctx, fullScreenshot(fmt.Sprintf("https://%s", host), 90, &buf))

	log.Println("screenshot: return")

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

//type Chrome struct {
//	userDataDir     string
//	logger          *slog.Logger
//	allocatorCtx    context.Context
//	allocatorCancel context.CancelFunc
//	browserCtx      context.Context
//	browserCancel   context.CancelFunc
//	jobs            <-chan string
//	results         <-chan int
//}
//
//func NewChrome(logger *slog.Logger) *Chrome {
//	userData, err := os.MkdirTemp("", "sulfur-screenshot-*")
//	if err != nil {
//		panic(err)
//	}
//
//	c := &Chrome{
//		userDataDir: userData,
//		logger:      logger,
//	}
//
//	go c.worker()
//
//	return c
//}
//
//func (c *Chrome) setup() {
//	// create context
//	c.allocatorCtx, c.allocatorCancel = chromedp.NewExecAllocator(
//		context.Background(),
//		//chromedp.Flag("host-resolver-rules", fmt.Sprintf("MAP %s %s", host, ip)),
//		chromedp.Flag("ignore-certificate-errors", "1"),
//		chromedp.Flag("headless", true),
//		chromedp.UserDataDir(c.userDataDir),
//	)
//
//	c.browserCtx, c.browserCancel = chromedp.NewContext(
//		c.allocatorCtx,
//	)
//}
//
//func (c *Chrome) stop() {
//	if err := os.RemoveAll(c.userDataDir); err != nil {
//		time.Sleep(3 * time.Second)
//		if err := os.RemoveAll(c.userDataDir); err != nil {
//			c.logger.Error("failed to delete chrome data dir", "dir", c.userDataDir, "err", err)
//		}
//	}
//}
//
//func (c *Chrome) Screenshot(host, ip string, logger *slog.Logger) ([]byte, error) {
//
//}
//
//func (c *Chrome) worker() {
//	for j := range c.jobs {
//
//		var buf []byte
//		// capture entire browser viewport, returning png with quality=90
//		err := chromedp.Run(c.browserCtx, chromedp.Navigate(fmt.Sprintf(j)))
//		if err != nil {
//			c.logger.Error("Ignoring navigation error", "error", err)
//		}
//		err = chromedp.Run(c.browserCtx, chromedp.Sleep(2 * time.Second), chromedp.FullScreenshot(&buf, 90))
//		if err != nil {
//			c.logger.Error("Ignoring navigation error", "error", err)
//		}
//
//
//	}
//}
//
//
//type screenshotResult struct {
//	dns
//}
