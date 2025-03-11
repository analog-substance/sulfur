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
					"collectionId": "pbc_3032838457",
					"hidden": false,
					"id": "relation1358299459",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "ip_port",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_3669933913",
					"hidden": false,
					"id": "relation563927626",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "certificate",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
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
			"id": "pbc_1792527727",
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_QuzSGMA6p0` + "`" + ` ON ` + "`" + `ip_port_certificatres` + "`" + ` (\n  ` + "`" + `ip_port` + "`" + `,\n  ` + "`" + `certificate` + "`" + `\n)"
			],
			"listRule": null,
			"name": "ip_port_certificatres",
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
		collection, err := app.FindCollectionByNameOrId("pbc_1792527727")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
