package main

import (
	"github.com/analog-substance/sulfur/pkg/pb_util"
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
	log.Println("Starting sulfur...")
	app := pocketbase.New()

	log.Println("Init App")
	pb_util.InitApp(app)

	forceMigrate := os.Getenv("FORCE_AUTOMIGRATE")

	// loosely check if it was executed using "go run"
	isGoRun := forceMigrate == "1" || strings.HasPrefix(os.Args[0], os.TempDir())

	log.Println("Register automigrate: ", isGoRun)
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Dashboard
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	log.Println("bind public dir setup")
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		log.Println("public dir added to routes")
		return se.Next()
	})

	log.Println("Start")
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
