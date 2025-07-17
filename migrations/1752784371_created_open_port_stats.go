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
					"cascadeDelete": false,
					"collectionId": "pbc_2873630990",
					"hidden": false,
					"id": "_clone_8F1t",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "organization",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"hidden": false,
					"id": "_clone_48l1",
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
					"id": "number3257917790",
					"max": null,
					"min": null,
					"name": "total",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				}
			],
			"id": "pbc_2941913654",
			"indexes": [],
			"listRule": null,
			"name": "open_port_stats",
			"system": false,
			"type": "view",
			"updateRule": null,
			"viewQuery": "SELECT (ROW_NUMBER() OVER()) as id, org_ip_addresses.organization, ip_ports.port, count(*) total\nFROM org_ip_addresses\n         INNER JOIN ip_ports ON ip_ports.ip_address = org_ip_addresses.ip_address and ip_ports.last_seen > datetime('now', '-24 hours')\nWHERE org_ip_addresses.last_seen > datetime('now', '-24 hours')\nGROUP BY org_ip_addresses.organization, ip_ports.port\nORDER BY TOTAL desc;",
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2941913654")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
