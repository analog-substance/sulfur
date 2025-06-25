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
					"id": "number3257917790",
					"max": null,
					"min": null,
					"name": "total",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_933409617",
					"hidden": false,
					"id": "_clone_vbeL",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "root_domain",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
				}
			],
			"id": "pbc_105466805",
			"indexes": [],
			"listRule": null,
			"name": "root_domain_stats",
			"system": false,
			"type": "view",
			"updateRule": null,
			"viewQuery": "SELECT (ROW_NUMBER() OVER()) as id, COUNT(*) total, root_domain from dns_records GROUP BY root_domain",
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_105466805")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
