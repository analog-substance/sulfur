package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2578508855")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_gVwxT2VASK` + "`" + ` ON ` + "`" + `org_environments` + "`" + ` (\n  ` + "`" + `organization` + "`" + `,\n  ` + "`" + `name` + "`" + `\n)"
			],
			"name": "org_environments"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2578508855")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_gVwxT2VASK` + "`" + ` ON ` + "`" + `environments` + "`" + ` (\n  ` + "`" + `organization` + "`" + `,\n  ` + "`" + `name` + "`" + `\n)"
			],
			"name": "environments"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
