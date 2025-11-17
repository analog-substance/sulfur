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
			"viewQuery": "SELECT ip_ports.id, org_domains.organization, dns_records.name domain, ip_ports.port, ip_ports.service, ip_ports.last_seen, ip_ports.created, count(*) total\nFROM ip_ports\nINNER JOIN ip_addresses on ip_ports.ip_address=ip_addresses.id\nINNER JOIN dns_records on dns_records.value=ip_addresses.address  AND dns_records.last_resolved > DATETIME('now', '-24 hours')\nINNER JOIN org_domains on org_domains.root_domain=dns_records.root_domain\nWHERE ip_ports.last_seen > DATETIME('now', '-24 hours')\nGROUP BY dns_records.name, ip_ports.port"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_IBoP")

		// remove field
		collection.Fields.RemoveById("_clone_X9wM")

		// remove field
		collection.Fields.RemoveById("_clone_6EqD")

		// remove field
		collection.Fields.RemoveById("_clone_Aghn")

		// remove field
		collection.Fields.RemoveById("_clone_f2kM")

		// remove field
		collection.Fields.RemoveById("_clone_Jqp6")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2873630990",
			"hidden": false,
			"id": "_clone_9cBG",
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
			"id": "_clone_kqM5",
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
			"id": "_clone_Niby",
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
			"id": "_clone_iXUD",
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
			"id": "_clone_JtKf",
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

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"hidden": false,
			"id": "_clone_Sgj2",
			"name": "created",
			"onCreate": true,
			"onUpdate": false,
			"presentable": false,
			"system": false,
			"type": "autodate"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(7, []byte(`{
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
			"viewQuery": "SELECT DISTINCT dns_records.id, org_domains.organization, dns_records.name domain,  ip_ports.port, ip_ports.service, ip_ports.last_seen, ip_ports.created\nFROM ip_ports\nINNER JOIN ip_addresses on ip_ports.ip_address=ip_addresses.id\nINNER JOIN dns_records on dns_records.value=ip_addresses.address  AND dns_records.last_resolved > DATETIME('now', '-24 hours')\nINNER JOIN org_domains on org_domains.root_domain=dns_records.root_domain\nWHERE ip_ports.last_seen > DATETIME('now', '-24 hours')\nGROUP BY dns_records.id, ip_ports.port"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2873630990",
			"hidden": false,
			"id": "_clone_IBoP",
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
			"id": "_clone_X9wM",
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
			"id": "_clone_6EqD",
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
			"id": "_clone_Aghn",
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
			"id": "_clone_f2kM",
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

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"hidden": false,
			"id": "_clone_Jqp6",
			"name": "created",
			"onCreate": true,
			"onUpdate": false,
			"presentable": false,
			"system": false,
			"type": "autodate"
		}`)); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_9cBG")

		// remove field
		collection.Fields.RemoveById("_clone_kqM5")

		// remove field
		collection.Fields.RemoveById("_clone_Niby")

		// remove field
		collection.Fields.RemoveById("_clone_iXUD")

		// remove field
		collection.Fields.RemoveById("_clone_JtKf")

		// remove field
		collection.Fields.RemoveById("_clone_Sgj2")

		// remove field
		collection.Fields.RemoveById("number3257917790")

		return app.Save(collection)
	})
}
