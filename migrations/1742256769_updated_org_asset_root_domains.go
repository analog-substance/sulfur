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
				"CREATE UNIQUE INDEX ` + "`" + `idx_7u5bAFB4uV` + "`" + ` ON ` + "`" + `asset_root_domains` + "`" + ` (` + "`" + `domain` + "`" + `)"
			],
			"name": "asset_root_domains"
		}`), &collection); err != nil {
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
				"CREATE UNIQUE INDEX ` + "`" + `idx_7u5bAFB4uV` + "`" + ` ON ` + "`" + `org_asset_root_domains` + "`" + ` (` + "`" + `domain` + "`" + `)"
			],
			"name": "org_asset_root_domains"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
