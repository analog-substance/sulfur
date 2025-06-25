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
					"id": "relation4102257691",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "org_id",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "_clone_ikog",
					"max": 0,
					"min": 0,
					"name": "org_name",
					"pattern": "",
					"presentable": true,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "_clone_3Kpc",
					"max": 0,
					"min": 0,
					"name": "domain",
					"pattern": "",
					"presentable": true,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "_clone_Jmpt",
					"max": 0,
					"min": 0,
					"name": "ip_address",
					"pattern": "",
					"presentable": true,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "_clone_hYsW",
					"max": "",
					"min": "",
					"name": "last_resolved",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "date"
				}
			],
			"id": "pbc_2813928730",
			"indexes": [],
			"listRule": null,
			"name": "subdomain_takeovers",
			"system": false,
			"type": "view",
			"updateRule": null,
			"viewQuery": "SELECT d.id, org.id org_id, org.name org_name, d.name domain, d.value ip_address, d.last_resolved\nFROM dns_records d\nINNER JOIN org_domains o ON o.root_domain=d.root_domain\nINNER JOIN organizations org on o.organization=org.id\nINNER JOIN ip_addresses a ON a.address = d.value\nLEFT JOIN org_ip_addresses i ON a.id=i.ip_address and o.organization=i.organization\nWHERE d.last_resolved > datetime('now', '-24 hour')\nAND i.id IS NULL;\n",
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2813928730")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
