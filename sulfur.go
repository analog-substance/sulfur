package main

import (
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/analog-substance/sulfur/pkg/router"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"log"
	"os"
	"strings"

	_ "github.com/analog-substance/sulfur/migrations"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()
	model.SetApp(app)

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Dashboard
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		return se.Next()
	})

	app.OnServe().BindFunc(router.AttachRoutes)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
