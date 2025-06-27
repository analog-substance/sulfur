package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2213552472")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT d.id, d.name domain, d.value ip_address, d.last_resolved, cr.fingerprint, cr.subject, cr.alternative_names, org.id organization, cr.id certificate, d.root_domain\nFROM dns_records d\nINNER JOIN org_domains o ON o.root_domain=d.root_domain\nINNER JOIN organizations org on o.organization=org.id\nINNER JOIN ip_addresses a ON a.address = d.value\nINNER JOIN ip_ports p ON a.id = p.ip_address\nINNER JOIN ip_port_certificates c on p.id = c.ip_port\nINNER JOIN certificates cr on cr.id = c.certificate\nLEFT JOIN org_certificates oc ON oc.certificate=c.certificate\nWHERE d.last_resolved > datetime('now', '-24 hour')\nAND oc.id IS NULL;"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation4102257691")

		// remove field
		collection.Fields.RemoveById("_clone_Vdmt")

		// remove field
		collection.Fields.RemoveById("_clone_4B9g")

		// remove field
		collection.Fields.RemoveById("_clone_nloT")

		// remove field
		collection.Fields.RemoveById("_clone_sBft")

		// remove field
		collection.Fields.RemoveById("_clone_D9ds")

		// remove field
		collection.Fields.RemoveById("_clone_an3H")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_zKIO",
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
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_86RK",
			"max": 0,
			"min": 0,
			"name": "ip_address",
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
			"id": "_clone_mZf0",
			"max": "",
			"min": "",
			"name": "last_resolved",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "date"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_px7f",
			"max": 0,
			"min": 0,
			"name": "fingerprint",
			"pattern": "",
			"presentable": true,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_SRRu",
			"max": 0,
			"min": 0,
			"name": "subject",
			"pattern": "",
			"presentable": true,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_tmi0",
			"max": 0,
			"min": 0,
			"name": "alternative_names",
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
		if err := collection.Fields.AddMarshaledJSONAt(7, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2873630990",
			"hidden": false,
			"id": "relation3253625724",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "organization",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(9, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_933409617",
			"hidden": false,
			"id": "_clone_FM6d",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "root_domain",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2213552472")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT d.id, org.id org_id, cr.id certificate, d.name domain, d.value ip_address, d.last_resolved, cr.fingerprint, cr.subject, cr.alternative_names\nFROM dns_records d\nINNER JOIN org_domains o ON o.root_domain=d.root_domain\nINNER JOIN organizations org on o.organization=org.id\nINNER JOIN ip_addresses a ON a.address = d.value\nINNER JOIN ip_ports p ON a.id = p.ip_address\nINNER JOIN ip_port_certificates c on p.id = c.ip_port\nINNER JOIN certificates cr on cr.id = c.certificate\nLEFT JOIN org_certificates oc ON oc.certificate=c.certificate\nWHERE d.last_resolved > datetime('now', '-24 hour')\nAND oc.id IS NULL;"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
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
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_Vdmt",
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
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_4B9g",
			"max": 0,
			"min": 0,
			"name": "ip_address",
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
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"hidden": false,
			"id": "_clone_nloT",
			"max": "",
			"min": "",
			"name": "last_resolved",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "date"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_sBft",
			"max": 0,
			"min": 0,
			"name": "fingerprint",
			"pattern": "",
			"presentable": true,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(7, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_D9ds",
			"max": 0,
			"min": 0,
			"name": "subject",
			"pattern": "",
			"presentable": true,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(8, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_an3H",
			"max": 0,
			"min": 0,
			"name": "alternative_names",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_zKIO")

		// remove field
		collection.Fields.RemoveById("_clone_86RK")

		// remove field
		collection.Fields.RemoveById("_clone_mZf0")

		// remove field
		collection.Fields.RemoveById("_clone_px7f")

		// remove field
		collection.Fields.RemoveById("_clone_SRRu")

		// remove field
		collection.Fields.RemoveById("_clone_tmi0")

		// remove field
		collection.Fields.RemoveById("relation3253625724")

		// remove field
		collection.Fields.RemoveById("_clone_FM6d")

		return app.Save(collection)
	})
}
