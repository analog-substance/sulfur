package jobs

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"sync"
	"time"
)

func AddJobs() {
	//slog.SetLogLoggerLevel(slog.LevelInfo)
	AddJob("resolve domains", "*/1 * * * *", ResolveDomains)
	//AddJob("Simple Port Scan", "*/1 * * * *", SimplePortScan)
	AddJob("Simple Port Scan", "*/1 * * * *", SimplePortScanWorkers)
	AddJob("Check Certs", "*/1 * * * *", CheckCerts)

}

var jobsMap = map[string]time.Time{}

var mu = sync.Mutex{}

func AddJob(jobId string, cronExpr string, run func()) {
	app_state.GetApp().Cron().MustAdd(jobId, cronExpr, NewJob(jobId, run))
}

func NewJob(jobId string, run func()) func() {
	return func() {
		logger := app_state.GetApp().Logger().WithGroup("Jobs")
		//slog.SetLogLoggerLevel(slog.LevelInfo)

		if !lockJob(jobId) {
			logger.Warn("Job status update", "job", jobId, "status", "already running")
			return
		}

		defer unlockJob(jobId)
		run()
		duration := getJobDuration(jobId)
		logger.Info("Job status update", "job", jobId, "status", "complete", "duration", duration)
	}
}

func lockJob(jobId string) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := jobsMap[jobId]; ok {
		return false
	}
	jobsMap[jobId] = time.Now()
	return true
}

func getJobDuration(jobId string) time.Duration {
	mu.Lock()
	defer mu.Unlock()
	if started, ok := jobsMap[jobId]; ok {
		return time.Since(started)
	}
	return time.Duration(0)
}

func unlockJob(jobId string) {
	mu.Lock()
	defer mu.Unlock()
	delete(jobsMap, jobId)
}
