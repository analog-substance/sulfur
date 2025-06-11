package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1792527727")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_QuzSGMA6p0` + "`" + ` ON ` + "`" + `ip_port_certificates` + "`" + ` (\n  ` + "`" + `ip_port` + "`" + `,\n  ` + "`" + `certificate` + "`" + `\n)"
			],
			"name": "ip_port_certificates"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1792527727")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_QuzSGMA6p0` + "`" + ` ON ` + "`" + `ip_port_certificatres` + "`" + ` (\n  ` + "`" + `ip_port` + "`" + `,\n  ` + "`" + `certificate` + "`" + `\n)"
			],
			"name": "ip_port_certificatres"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
