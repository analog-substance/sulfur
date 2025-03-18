package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_960090968")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_7u5bAFB4uV` + "`" + ` ON ` + "`" + `org_asset_root_domains` + "`" + ` (` + "`" + `domain` + "`" + `)"
			],
			"name": "org_asset_root_domains"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation1176952354")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2873630990",
			"hidden": false,
			"id": "relation3253625724",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "organization",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_960090968")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_zKMNdjGcUd` + "`" + ` ON ` + "`" + `env_root_domains` + "`" + ` (\n  ` + "`" + `environment` + "`" + `,\n  ` + "`" + `domain` + "`" + `\n)"
			],
			"name": "env_root_domains"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2578508855",
			"hidden": false,
			"id": "relation1176952354",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "environment",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation3253625724")

		return app.Save(collection)
	})
}
