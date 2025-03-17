package model

import "github.com/pocketbase/pocketbase"

var app *pocketbase.PocketBase

func SetApp(sapp *pocketbase.PocketBase) {
	app = sapp
}

func GetApp() *pocketbase.PocketBase {
	if app == nil {
		panic("app not initialized")
	}

	return app
}
