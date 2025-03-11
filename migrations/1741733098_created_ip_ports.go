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
					"autogeneratePattern": "[a-z0-9]{15}",
					"hidden": false,
					"id": "text3208210256",
					"max": 15,
					"min": 15,
					"name": "id",
					"pattern": "^[a-z0-9]+$",
					"presentable": false,
					"primaryKey": true,
					"required": true,
					"system": true,
					"type": "text"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_2789751422",
					"hidden": false,
					"id": "relation587191692",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "ip_address",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text1872607463",
					"max": 0,
					"min": 0,
					"name": "banner",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text3785202386",
					"max": 0,
					"min": 0,
					"name": "service",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text3225476321",
					"max": 0,
					"min": 0,
					"name": "app_protocol",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text3368074316",
					"max": 0,
					"min": 0,
					"name": "protocol",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "number1133600204",
					"max": null,
					"min": null,
					"name": "port",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				},
				{
					"hidden": false,
					"id": "date846843460",
					"max": "",
					"min": "",
					"name": "last_seen",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "date"
				},
				{
					"hidden": false,
					"id": "autodate2990389176",
					"name": "created",
					"onCreate": true,
					"onUpdate": false,
					"presentable": false,
					"system": false,
					"type": "autodate"
				},
				{
					"hidden": false,
					"id": "autodate3332085495",
					"name": "updated",
					"onCreate": true,
					"onUpdate": true,
					"presentable": false,
					"system": false,
					"type": "autodate"
				}
			],
			"id": "pbc_3032838457",
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_xhKYMuxKIV` + "`" + ` ON ` + "`" + `ip_ports` + "`" + ` (\n  ` + "`" + `ip_address` + "`" + `,\n  ` + "`" + `port` + "`" + `,\n  ` + "`" + `protocol` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_JgG6lkIJtj` + "`" + ` ON ` + "`" + `ip_ports` + "`" + ` (` + "`" + `last_seen` + "`" + `)"
			],
			"listRule": null,
			"name": "ip_ports",
			"system": false,
			"type": "base",
			"updateRule": null,
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3032838457")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
