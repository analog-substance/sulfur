package model

import (
	"database/sql"
	"errors"
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func FirstOrCreateByFilter(nameOrID string, filter string, params ...dbx.Params) (*core.Record, error) {

	record := &core.Record{}

	err := app_state.GetApp().RecordQuery(nameOrID).
		AndWhere(dbx.NewExp(filter, params...)).
		Limit(1).
		One(record)

	//
	//record, err := GetApp().FindFirstRecordByFilter(
	//	nameOrID,
	//	filter,
	//	params...,
	//)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			// an error that isn't empty results
			return nil, err
		}
		collection, err := app_state.GetApp().FindCollectionByNameOrId(nameOrID)
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
