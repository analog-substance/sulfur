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
				"CREATE INDEX ` + "`" + `idx_ZWSHfwMZyu` + "`" + ` ON ` + "`" + `org_ip_addresses` + "`" + ` (` + "`" + `last_seen` + "`" + `)"
			]
		}`), &collection); err != nil {
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
				"CREATE UNIQUE INDEX ` + "`" + `idx_shA142XKhm` + "`" + ` ON ` + "`" + `org_ip_addresses` + "`" + ` (\n  ` + "`" + `organization` + "`" + `,\n  ` + "`" + `ip_address` + "`" + `\n)"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
