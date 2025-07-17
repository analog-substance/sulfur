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
					"id": "_clone_PsDQ",
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
					"id": "_clone_V9EI",
					"max": 0,
					"min": 0,
					"name": "address",
					"pattern": "",
					"presentable": true,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "_clone_OAUS",
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
					"id": "_clone_Q57K",
					"max": "",
					"min": "",
					"name": "last_seen",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "date"
				},
				{
					"hidden": false,
					"id": "json917420102",
					"maxSize": 1,
					"name": "dns_names",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "json"
				}
			],
			"id": "pbc_1745364752",
			"indexes": [],
			"listRule": null,
			"name": "org_ip_ports",
			"system": false,
			"type": "view",
			"updateRule": null,
			"viewQuery": "SELECT (ROW_NUMBER() OVER()) as id, org_ip_addresses.organization, ip_addresses.address, ip_ports.port, org_ip_addresses.last_seen, group_concat(dns_records.name) dns_names\nFROM org_ip_addresses\n         INNER JOIN ip_addresses ON ip_addresses.id = org_ip_addresses.ip_address\n         INNER JOIN ip_ports ON ip_ports.ip_address = org_ip_addresses.ip_address and ip_ports.last_seen > datetime('now', '-24 hours')\n        LEFT JOIN dns_records ON dns_records.value=ip_addresses.address AND dns_records.type=\"A\"\nWHERE org_ip_addresses.last_seen > datetime('now', '-24 hours')\nGROUP BY org_ip_addresses.organization, ip_addresses.id, ip_ports.port\nORDER BY org_ip_addresses.last_seen desc, ip_ports.port asc",
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1745364752")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
