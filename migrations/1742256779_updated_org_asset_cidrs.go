package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_548419197")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_Lka7RTqNRw` + "`" + ` ON ` + "`" + `asset_cidrs` + "`" + ` (\n  ` + "`" + `cidr` + "`" + `,\n  ` + "`" + `organization` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_LzLzy5UctW` + "`" + ` ON ` + "`" + `asset_cidrs` + "`" + ` (\n  ` + "`" + `is_private` + "`" + `,\n  ` + "`" + `is_shared` + "`" + `,\n  ` + "`" + `is_loopback` + "`" + `,\n  ` + "`" + `is_multicast` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_2TYcEw19Ij` + "`" + ` ON ` + "`" + `asset_cidrs` + "`" + ` (\n  ` + "`" + `is_ephemeral` + "`" + `,\n  ` + "`" + `last_seen` + "`" + `\n)"
			],
			"name": "asset_cidrs"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_548419197")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_Lka7RTqNRw` + "`" + ` ON ` + "`" + `org_asset_cidrs` + "`" + ` (\n  ` + "`" + `cidr` + "`" + `,\n  ` + "`" + `organization` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_LzLzy5UctW` + "`" + ` ON ` + "`" + `org_asset_cidrs` + "`" + ` (\n  ` + "`" + `is_private` + "`" + `,\n  ` + "`" + `is_shared` + "`" + `,\n  ` + "`" + `is_loopback` + "`" + `,\n  ` + "`" + `is_multicast` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_2TYcEw19Ij` + "`" + ` ON ` + "`" + `org_asset_cidrs` + "`" + ` (\n  ` + "`" + `is_ephemeral` + "`" + `,\n  ` + "`" + `last_seen` + "`" + `\n)"
			],
			"name": "org_asset_cidrs"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
