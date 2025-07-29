package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_778186020")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_shA142XKhm` + "`" + ` ON ` + "`" + `org_ip_addresses` + "`" + ` (\n  ` + "`" + `organization` + "`" + `,\n  ` + "`" + `ip_address` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_ZWSHfwMZyu` + "`" + ` ON ` + "`" + `org_ip_addresses` + "`" + ` (` + "`" + `last_seen` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_D5beHGX7sv` + "`" + ` ON ` + "`" + `org_ip_addresses` + "`" + ` (` + "`" + `external_reference` + "`" + `)"
			]
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
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
		collection, err := app.FindCollectionByNameOrId("pbc_778186020")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_shA142XKhm` + "`" + ` ON ` + "`" + `org_ip_addresses` + "`" + ` (\n  ` + "`" + `organization` + "`" + `,\n  ` + "`" + `ip_address` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_ZWSHfwMZyu` + "`" + ` ON ` + "`" + `org_ip_addresses` + "`" + ` (` + "`" + `last_seen` + "`" + `)"
			]
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation2331567623")

		return app.Save(collection)
	})
}
