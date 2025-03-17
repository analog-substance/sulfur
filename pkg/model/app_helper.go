package model

import (
	"database/sql"
	"errors"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

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

func FirstOrCreateByFilter(nameOrID string, filter string, params ...dbx.Params) (*core.Record, error) {

	record, err := GetApp().FindFirstRecordByFilter(
		nameOrID,
		filter,
		params...,
	)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			// an error that isn't empty results
			return nil, err
		}
		collection, err := app.FindCollectionByNameOrId(nameOrID)
		if err != nil {
			return nil, err
		}
		record = core.NewRecord(collection)

		for _, param := range params {
			for attr, val := range param {
				record.Set(attr, val)
			}
		}
	}
	return record, nil
}
