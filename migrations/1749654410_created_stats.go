package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"createRule": null,
			"deleteRule": null,
			"fields": [
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text3208210256",
					"max": 0,
					"min": 0,
					"name": "id",
					"pattern": "^[a-z0-9]+$",
					"presentable": false,
					"primaryKey": true,
					"required": true,
					"system": true,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "json1417809052",
					"maxSize": 1,
					"name": "active_dns_records",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "json"
				}
			],
			"id": "pbc_4107852781",
			"indexes": [],
			"listRule": null,
			"name": "stats",
			"system": false,
			"type": "view",
			"updateRule": null,
			"viewQuery": "SELECT 1 as id, (\n  SELECT COUNT(*) FROM dns_records WHERE last_seen > DATETIME('now', '-8 hours')\n  ) active_dns_records\n",
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_4107852781")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
