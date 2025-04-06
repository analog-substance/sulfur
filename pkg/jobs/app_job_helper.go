package jobs

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"log/slog"
	"sync"
	"time"
)

func AddJobs() {
	AddJob("resolve domains", "*/1 * * * *", ResolveDomains)
}

var jobsMap = map[string]time.Time{}

var mu = sync.Mutex{}

func AddJob(jobId string, cronExpr string, run func()) {
	app_state.GetApp().Cron().MustAdd(jobId, cronExpr, NewJob(jobId, run))
}

func NewJob(jobId string, run func()) func() {
	return func() {
		slog.SetLogLoggerLevel(slog.LevelInfo)

		if _, ok := jobsMap[jobId]; ok {
			app_state.GetApp().Logger().Info("Job already running", "job", jobId)
		}

		mu.Lock()
		jobsMap[jobId] = time.Now()
		mu.Unlock()
		defer func() {
			mu.Lock()
			delete(jobsMap, jobId)
			mu.Unlock()
			app_state.GetApp().Logger().Info("Job complete", "job", jobId, "duration", time.Since(jobsMap[jobId]))
		}()
		app_state.GetApp().Logger().Info("Job starting", "job", jobId)

		run()
	}
}
