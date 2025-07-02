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
			"viewQuery": "SELECT d.id, artifacts.id artifact_id, d.name domain, d.value ip_address, cr.fingerprint, cr.subject, cr.alternative_names, external_references.value, external_references.name, org.id organization, cr.id certificate, d.root_domain\nFROM dns_records d\n         INNER JOIN org_domains o ON o.root_domain=d.root_domain\n         INNER JOIN organizations org on o.organization=org.id\n         INNER JOIN ip_addresses a ON a.address = d.value\n         INNER JOIN ip_ports p ON a.id = p.ip_address AND p.last_seen > datetime('now', '-24 hour')\n         INNER JOIN ip_port_certificates c on p.id = c.ip_port\n         INNER JOIN certificates cr on cr.id = c.certificate\n         LEFT JOIN org_certificates oc ON oc.certificate=c.certificate\n         LEFT JOIN artifacts on artifacts.dns_record=d.id\n         LEFT JOIN external_references on external_references.id=d.external_reference\nWHERE d.last_resolved > datetime('now', '-24 hour')\n  AND oc.id IS NULL;"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_PVw6")

		// remove field
		collection.Fields.RemoveById("_clone_Ebh1")

		// remove field
		collection.Fields.RemoveById("_clone_TcrL")

		// remove field
		collection.Fields.RemoveById("_clone_VcE0")

		// remove field
		collection.Fields.RemoveById("_clone_aZyL")

		// remove field
		collection.Fields.RemoveById("_clone_KdAs")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_lKIR",
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
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_Tfcu",
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
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_d8fx",
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
			"id": "_clone_G99H",
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
			"id": "_clone_RyBH",
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
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_bnmX",
			"max": 0,
			"min": 0,
			"name": "value",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": true,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(8, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_ov9A",
			"max": 0,
			"min": 0,
			"name": "name",
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
		if err := collection.Fields.AddMarshaledJSONAt(11, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_933409617",
			"hidden": false,
			"id": "_clone_XLQD",
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
			"viewQuery": "SELECT d.id, artifacts.id artifact_id, d.name domain, d.value ip_address, cr.fingerprint, cr.subject, cr.alternative_names, org.id organization, cr.id certificate, d.root_domain\nFROM dns_records d\n         INNER JOIN org_domains o ON o.root_domain=d.root_domain\n         INNER JOIN organizations org on o.organization=org.id\n         INNER JOIN ip_addresses a ON a.address = d.value\n         INNER JOIN ip_ports p ON a.id = p.ip_address AND p.last_seen > datetime('now', '-24 hour')\n         INNER JOIN ip_port_certificates c on p.id = c.ip_port\n         INNER JOIN certificates cr on cr.id = c.certificate\n         LEFT JOIN org_certificates oc ON oc.certificate=c.certificate\n         LEFT JOIN artifacts on artifacts.dns_record=d.id\nWHERE d.last_resolved > datetime('now', '-24 hour')\n  AND oc.id IS NULL;\n"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_PVw6",
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
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_Ebh1",
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
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_TcrL",
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
			"id": "_clone_VcE0",
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
			"id": "_clone_aZyL",
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
		if err := collection.Fields.AddMarshaledJSONAt(9, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_933409617",
			"hidden": false,
			"id": "_clone_KdAs",
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

		// remove field
		collection.Fields.RemoveById("_clone_lKIR")

		// remove field
		collection.Fields.RemoveById("_clone_Tfcu")

		// remove field
		collection.Fields.RemoveById("_clone_d8fx")

		// remove field
		collection.Fields.RemoveById("_clone_G99H")

		// remove field
		collection.Fields.RemoveById("_clone_RyBH")

		// remove field
		collection.Fields.RemoveById("_clone_bnmX")

		// remove field
		collection.Fields.RemoveById("_clone_ov9A")

		// remove field
		collection.Fields.RemoveById("_clone_XLQD")

		return app.Save(collection)
	})
}
