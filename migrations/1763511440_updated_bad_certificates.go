package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3199248551")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT d.id, artifacts.id artifact_id, d.name domain, d.value ip_address,  cr.subject, cr.alternative_names, cr.serial, cr.issuer, cr.not_before, external_references.value external_id, external_references.name external_name, org.id organization, cr.id certificate, d.root_domain\nFROM dns_records d\n    INNER JOIN org_domains o ON o.root_domain=d.root_domain\n    INNER JOIN organizations org on o.organization=org.id\n    INNER JOIN ip_addresses a ON a.address = d.value\n    INNER JOIN ip_ports p ON a.id = p.ip_address AND p.last_seen > datetime('now', '-8 hour')\n    INNER JOIN ip_port_certificates c on p.id = c.ip_port AND c.last_seen > datetime('now', '-8 hour')\n    INNER JOIN certificates cr on cr.id = c.certificate\n    LEFT JOIN org_certificates oc ON oc.serial=cr.serial AND oc.issuer=cr.issuer AND oc.not_before=cr.not_before\n    LEFT JOIN artifacts on artifacts.dns_record=d.id\n    LEFT JOIN external_references on external_references.id=d.external_reference\nWHERE d.last_resolved > datetime('now', '-8 hour')\n  AND oc.id IS NULL;"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_7HGH")

		// remove field
		collection.Fields.RemoveById("_clone_2fZp")

		// remove field
		collection.Fields.RemoveById("_clone_1lQM")

		// remove field
		collection.Fields.RemoveById("_clone_ugt6")

		// remove field
		collection.Fields.RemoveById("_clone_qDLr")

		// remove field
		collection.Fields.RemoveById("_clone_0rXK")

		// remove field
		collection.Fields.RemoveById("_clone_TbKa")

		// remove field
		collection.Fields.RemoveById("_clone_Q1hF")

		// remove field
		collection.Fields.RemoveById("_clone_qPkL")

		// remove field
		collection.Fields.RemoveById("_clone_Plyr")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_8aJs",
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
			"id": "_clone_7p0D",
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
			"id": "_clone_Ay2A",
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
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_mMdM",
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
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_yQrW",
			"max": 0,
			"min": 0,
			"name": "serial",
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
			"id": "_clone_Hx0k",
			"max": 0,
			"min": 0,
			"name": "issuer",
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
			"hidden": false,
			"id": "_clone_bbZL",
			"max": "",
			"min": "",
			"name": "not_before",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "date"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(9, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_puRj",
			"max": 0,
			"min": 0,
			"name": "external_id",
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
		if err := collection.Fields.AddMarshaledJSONAt(10, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_pQ3Y",
			"max": 0,
			"min": 0,
			"name": "external_name",
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
		if err := collection.Fields.AddMarshaledJSONAt(13, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_933409617",
			"hidden": false,
			"id": "_clone_y5F1",
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
		collection, err := app.FindCollectionByNameOrId("pbc_3199248551")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT d.id, artifacts.id artifact_id, d.name domain, d.value ip_address,  cr.subject, cr.alternative_names, cr.serial, cr.issuer, cr.not_before, external_references.value external_id, external_references.name external_name, org.id organization, cr.id certificate, d.root_domain\nFROM dns_records d\n    INNER JOIN org_domains o ON o.root_domain=d.root_domain\n    INNER JOIN organizations org on o.organization=org.id\n    INNER JOIN ip_addresses a ON a.address = d.value\n    INNER JOIN ip_ports p ON a.id = p.ip_address AND p.last_seen > datetime('now', '-8 hour')\n    INNER JOIN ip_port_certificates c on p.id = c.ip_port AND c.last_seen > datetime('now', '-8 hour')\n    INNER JOIN certificates cr on cr.id = c.certificate\n    LEFT JOIN org_certificates oc ON oc.serial=cr.serial\n    LEFT JOIN artifacts on artifacts.dns_record=d.id\n    LEFT JOIN external_references on external_references.id=d.external_reference\nWHERE d.last_resolved > datetime('now', '-8 hour')\n  AND oc.id IS NULL;"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_7HGH",
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
			"id": "_clone_2fZp",
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
			"id": "_clone_1lQM",
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
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_ugt6",
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
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_qDLr",
			"max": 0,
			"min": 0,
			"name": "serial",
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
			"id": "_clone_0rXK",
			"max": 0,
			"min": 0,
			"name": "issuer",
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
			"hidden": false,
			"id": "_clone_TbKa",
			"max": "",
			"min": "",
			"name": "not_before",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "date"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(9, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_Q1hF",
			"max": 0,
			"min": 0,
			"name": "external_id",
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
		if err := collection.Fields.AddMarshaledJSONAt(10, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_qPkL",
			"max": 0,
			"min": 0,
			"name": "external_name",
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
		if err := collection.Fields.AddMarshaledJSONAt(13, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_933409617",
			"hidden": false,
			"id": "_clone_Plyr",
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
		collection.Fields.RemoveById("_clone_8aJs")

		// remove field
		collection.Fields.RemoveById("_clone_7p0D")

		// remove field
		collection.Fields.RemoveById("_clone_Ay2A")

		// remove field
		collection.Fields.RemoveById("_clone_mMdM")

		// remove field
		collection.Fields.RemoveById("_clone_yQrW")

		// remove field
		collection.Fields.RemoveById("_clone_Hx0k")

		// remove field
		collection.Fields.RemoveById("_clone_bbZL")

		// remove field
		collection.Fields.RemoveById("_clone_puRj")

		// remove field
		collection.Fields.RemoveById("_clone_pQ3Y")

		// remove field
		collection.Fields.RemoveById("_clone_y5F1")

		return app.Save(collection)
	})
}
