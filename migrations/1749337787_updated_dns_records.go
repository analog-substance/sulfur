package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_38995645")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_MNinmJjbyT` + "`" + ` ON ` + "`" + `dns_records` + "`" + ` (\n  ` + "`" + `name` + "`" + `,\n  ` + "`" + `type` + "`" + `,\n  ` + "`" + `value` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_nmGpDHBWK2` + "`" + ` ON ` + "`" + `dns_records` + "`" + ` (\n  ` + "`" + `last_seen` + "`" + `,\n  ` + "`" + `last_resolved` + "`" + `,\n  ` + "`" + `resolve_error_count` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_uy6vJmVLRG` + "`" + ` ON ` + "`" + `dns_records` + "`" + ` (` + "`" + `root_domain` + "`" + `)"
			]
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(9, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_933409617",
			"hidden": false,
			"id": "relation2232677066",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "root_domain",
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

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_MNinmJjbyT` + "`" + ` ON ` + "`" + `dns_records` + "`" + ` (\n  ` + "`" + `name` + "`" + `,\n  ` + "`" + `type` + "`" + `,\n  ` + "`" + `value` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_nmGpDHBWK2` + "`" + ` ON ` + "`" + `dns_records` + "`" + ` (\n  ` + "`" + `last_seen` + "`" + `,\n  ` + "`" + `last_resolved` + "`" + `,\n  ` + "`" + `resolve_error_count` + "`" + `\n)"
			]
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation2232677066")

		return app.Save(collection)
	})
}
