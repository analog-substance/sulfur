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
				"CREATE INDEX ` + "`" + `idx_nmGpDHBWK2` + "`" + ` ON ` + "`" + `dns_records` + "`" + ` (\n  ` + "`" + `last_seen` + "`" + `,\n  ` + "`" + `last_resolved` + "`" + `,\n  ` + "`" + `resolve_error_count` + "`" + `\n)"
			],
			"name": "dns_records"
		}`), &collection); err != nil {
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
				"CREATE UNIQUE INDEX ` + "`" + `idx_MNinmJjbyT` + "`" + ` ON ` + "`" + `dns_resolutions` + "`" + ` (\n  ` + "`" + `name` + "`" + `,\n  ` + "`" + `type` + "`" + `,\n  ` + "`" + `value` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_nmGpDHBWK2` + "`" + ` ON ` + "`" + `dns_resolutions` + "`" + ` (\n  ` + "`" + `last_seen` + "`" + `,\n  ` + "`" + `last_resolved` + "`" + `,\n  ` + "`" + `resolve_error_count` + "`" + `\n)"
			],
			"name": "dns_resolutions"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
