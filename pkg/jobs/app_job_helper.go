package jobs

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"log"
	"sync"
	"time"
)

func AddJobs() {
	//slog.SetLogLoggerLevel(slog.LevelInfo)
	AddJob("resolve domains", "*/1 * * * *", ResolveDomains)
	AddJob("Simple Port Scan", "*/1 * * * *", SimplePortScan)
}

var jobsMap = map[string]time.Time{}

var mu = sync.Mutex{}

func AddJob(jobId string, cronExpr string, run func()) {
	app_state.GetApp().Cron().MustAdd(jobId, cronExpr, NewJob(jobId, run))
}

func NewJob(jobId string, run func()) func() {
	return func() {
		//slog.SetLogLoggerLevel(slog.LevelInfo)

		if _, ok := jobsMap[jobId]; ok {
			app_state.GetApp().Logger().Info("Job already running", "job", jobId)
			log.Println("Job already running", "job", jobId)
			return
		}

		mu.Lock()
		jobsMap[jobId] = time.Now()
		mu.Unlock()
		defer func() {
			mu.Lock()
			delete(jobsMap, jobId)
			mu.Unlock()
			app_state.GetApp().Logger().Info("Job complete", "job", jobId, "duration", time.Since(jobsMap[jobId]))
			log.Println("Job complete", "job", jobId)
		}()
		app_state.GetApp().Logger().Info("Job starting", "job", jobId)
		log.Println("Job starting", "job", jobId)

		run()
	}
}
