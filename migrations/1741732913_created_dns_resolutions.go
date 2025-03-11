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
					"collectionId": "pbc_960090968",
					"hidden": false,
					"id": "relation2232677066",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "root_domain",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text1579384326",
					"max": 0,
					"min": 0,
					"name": "name",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text2363381545",
					"max": 0,
					"min": 0,
					"name": "type",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text494360628",
					"max": 0,
					"min": 0,
					"name": "value",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "number2750318623",
					"max": null,
					"min": null,
					"name": "ttl",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text1304463345",
					"max": 0,
					"min": 0,
					"name": "resolve_error",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "number3706125909",
					"max": null,
					"min": null,
					"name": "resolve_error_count",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				},
				{
					"hidden": false,
					"id": "date3484784209",
					"max": "",
					"min": "",
					"name": "last_resolved",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "date"
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
			"id": "pbc_38995645",
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_MNinmJjbyT` + "`" + ` ON ` + "`" + `dns_resolutions` + "`" + ` (\n  ` + "`" + `name` + "`" + `,\n  ` + "`" + `type` + "`" + `,\n  ` + "`" + `value` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_nmGpDHBWK2` + "`" + ` ON ` + "`" + `dns_resolutions` + "`" + ` (\n  ` + "`" + `last_seen` + "`" + `,\n  ` + "`" + `last_resolved` + "`" + `,\n  ` + "`" + `resolve_error_count` + "`" + `\n)"
			],
			"listRule": null,
			"name": "dns_resolutions",
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
		collection, err := app.FindCollectionByNameOrId("pbc_38995645")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
