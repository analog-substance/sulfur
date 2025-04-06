package pb_util

import (
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/jobs"
	"github.com/analog-substance/sulfur/pkg/router"
	"github.com/pocketbase/pocketbase"
)

func InitApp(app *pocketbase.PocketBase) {
	app_state.SetApp(app)
	jobs.AddJobs()
	app.OnServe().BindFunc(router.AttachRoutes)
}
