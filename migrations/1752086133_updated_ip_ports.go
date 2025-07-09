package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3032838457")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_xhKYMuxKIV` + "`" + ` ON ` + "`" + `ip_ports` + "`" + ` (\n  ` + "`" + `ip_address` + "`" + `,\n  ` + "`" + `port` + "`" + `,\n  ` + "`" + `protocol` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_JgG6lkIJtj` + "`" + ` ON ` + "`" + `ip_ports` + "`" + ` (` + "`" + `last_seen` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_5ViGIBqzs5` + "`" + ` ON ` + "`" + `ip_ports` + "`" + ` (` + "`" + `port` + "`" + `)"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3032838457")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_xhKYMuxKIV` + "`" + ` ON ` + "`" + `ip_ports` + "`" + ` (\n  ` + "`" + `ip_address` + "`" + `,\n  ` + "`" + `port` + "`" + `,\n  ` + "`" + `protocol` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_JgG6lkIJtj` + "`" + ` ON ` + "`" + `ip_ports` + "`" + ` (` + "`" + `last_seen` + "`" + `)"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
