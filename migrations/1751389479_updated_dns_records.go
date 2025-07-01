package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_38995645")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(10, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_3310716100",
			"hidden": false,
			"id": "relation2331567623",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "external_reference",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_38995645")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation2331567623")

		return app.Save(collection)
	})
}
