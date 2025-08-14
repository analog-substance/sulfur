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
					"id": "_clone_VmB8",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "organization",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "_clone_WerB",
					"max": 0,
					"min": 0,
					"name": "name",
					"pattern": "",
					"presentable": true,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "_clone_y1Gb",
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
					"autogeneratePattern": "",
					"hidden": false,
					"id": "_clone_9oCe",
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
					"hidden": false,
					"id": "_clone_jDrN",
					"max": "",
					"min": "",
					"name": "last_seen",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "date"
				}
			],
			"id": "pbc_1229338967",
			"indexes": [],
			"listRule": null,
			"name": "org_domain_ports",
			"system": false,
			"type": "view",
			"updateRule": null,
			"viewQuery": "SELECT DISTINCT dns_records.id, org_domains.organization, dns_records.name, ip_ports.port, ip_ports.service, ip_ports.last_seen\nFROM ip_ports\nINNER JOIN ip_addresses on ip_ports.ip_address=ip_addresses.id\nINNER JOIN dns_records on dns_records.value=ip_addresses.address  AND dns_records.last_resolved > DATETIME('now', '-24 hours')\nINNER JOIN org_domains on org_domains.root_domain=dns_records.root_domain\nWHERE ip_ports.last_seen > DATETIME('now', '-24 hours')",
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1229338967")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
