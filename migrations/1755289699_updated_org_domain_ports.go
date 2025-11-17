package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1229338967")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT DISTINCT dns_records.id, org_domains.organization, dns_records.name domain, ip_ports.port, ip_ports.service, ip_ports.last_seen\nFROM ip_ports\nINNER JOIN ip_addresses on ip_ports.ip_address=ip_addresses.id\nINNER JOIN dns_records on dns_records.value=ip_addresses.address  AND dns_records.last_resolved > DATETIME('now', '-24 hours')\nINNER JOIN org_domains on org_domains.root_domain=dns_records.root_domain\nWHERE ip_ports.last_seen > DATETIME('now', '-24 hours')\nGROUP BY dns_records.id, ip_ports.port"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_Rfkv")

		// remove field
		collection.Fields.RemoveById("_clone_NUv5")

		// remove field
		collection.Fields.RemoveById("_clone_E2dB")

		// remove field
		collection.Fields.RemoveById("_clone_z3B0")

		// remove field
		collection.Fields.RemoveById("_clone_QAp5")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2873630990",
			"hidden": false,
			"id": "_clone_Oyds",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "organization",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_t45T",
			"max": 0,
			"min": 0,
			"name": "domain",
			"pattern": "",
			"presentable": true,
			"primaryKey": false,
			"required": true,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"hidden": false,
			"id": "_clone_pXu2",
			"max": null,
			"min": null,
			"name": "port",
			"onlyInt": false,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_muam",
			"max": 0,
			"min": 0,
			"name": "service",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"hidden": false,
			"id": "_clone_H7zl",
			"max": "",
			"min": "",
			"name": "last_seen",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "date"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1229338967")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT DISTINCT dns_records.id, org_domains.organization, dns_records.name domain, ip_ports.port, ip_ports.service, ip_ports.last_seen\nFROM ip_ports\nINNER JOIN ip_addresses on ip_ports.ip_address=ip_addresses.id\nINNER JOIN dns_records on dns_records.value=ip_addresses.address  AND dns_records.last_resolved > DATETIME('now', '-24 hours')\nINNER JOIN org_domains on org_domains.root_domain=dns_records.root_domain\nWHERE ip_ports.last_seen > DATETIME('now', '-24 hours')"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2873630990",
			"hidden": false,
			"id": "_clone_Rfkv",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "organization",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_NUv5",
			"max": 0,
			"min": 0,
			"name": "domain",
			"pattern": "",
			"presentable": true,
			"primaryKey": false,
			"required": true,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"hidden": false,
			"id": "_clone_E2dB",
			"max": null,
			"min": null,
			"name": "port",
			"onlyInt": false,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_z3B0",
			"max": 0,
			"min": 0,
			"name": "service",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"hidden": false,
			"id": "_clone_QAp5",
			"max": "",
			"min": "",
			"name": "last_seen",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "date"
		}`)); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_Oyds")

		// remove field
		collection.Fields.RemoveById("_clone_t45T")

		// remove field
		collection.Fields.RemoveById("_clone_pXu2")

		// remove field
		collection.Fields.RemoveById("_clone_muam")

		// remove field
		collection.Fields.RemoveById("_clone_H7zl")

		return app.Save(collection)
	})
}
