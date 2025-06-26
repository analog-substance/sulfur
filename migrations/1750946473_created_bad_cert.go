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
					"id": "_clone_Pr6P",
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
					"id": "_clone_iMkl",
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
					"id": "_clone_1f1E",
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
					"id": "_clone_hHxu",
					"max": "",
					"min": "",
					"name": "last_resolved",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "date"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "_clone_Vj8O",
					"max": 0,
					"min": 0,
					"name": "fingerprint",
					"pattern": "",
					"presentable": true,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "_clone_hG4e",
					"max": 0,
					"min": 0,
					"name": "subject",
					"pattern": "",
					"presentable": true,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "_clone_7UvQ",
					"max": 0,
					"min": 0,
					"name": "alternative_names",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				}
			],
			"id": "pbc_2213552472",
			"indexes": [],
			"listRule": null,
			"name": "bad_cert",
			"system": false,
			"type": "view",
			"updateRule": null,
			"viewQuery": "SELECT d.id, org.id org_id, org.name org_name, d.name domain, d.value ip_address, d.last_resolved, cr.fingerprint, cr.subject, cr.alternative_names\nFROM dns_records d\nINNER JOIN org_domains o ON o.root_domain=d.root_domain\nINNER JOIN organizations org on o.organization=org.id\nINNER JOIN ip_addresses a ON a.address = d.value\nINNER JOIN ip_ports p ON a.id = p.ip_address\nINNER JOIN ip_port_certificates c on p.id = c.ip_port\nINNER JOIN certificates cr on cr.id = c.certificate\nLEFT JOIN org_certificates oc ON oc.certificate=c.certificate\nWHERE d.last_resolved > datetime('now', '-24 hour')\nAND oc.id IS NULL;",
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2213552472")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
