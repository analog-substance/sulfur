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
			"viewQuery": "SELECT d.id, artifacts.id artifact_id, d.name domain, d.value ip_address, cr.fingerprint, cr.subject, cr.alternative_names, external_references.value external_id, external_references.name external_ref_name, org.id organization, cr.id certificate, d.root_domain\nFROM dns_records d\n         INNER JOIN org_domains o ON o.root_domain=d.root_domain\n         INNER JOIN organizations org on o.organization=org.id\n         INNER JOIN ip_addresses a ON a.address = d.value\n         INNER JOIN ip_ports p ON a.id = p.ip_address AND p.last_seen > datetime('now', '-8 hour')\n         INNER JOIN ip_port_certificates c on p.id = c.ip_port AND c.last_seen > datetime('now', '-8 hour')\n         INNER JOIN certificates cr on cr.id = c.certificate\n         LEFT JOIN org_certificates oc ON oc.certificate=c.certificate\n         LEFT JOIN artifacts on artifacts.dns_record=d.id\n         LEFT JOIN external_references on external_references.id=d.external_reference\nWHERE d.last_resolved > datetime('now', '-8 hour')\n  AND oc.id IS NULL;"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_gJoB")

		// remove field
		collection.Fields.RemoveById("_clone_2rR6")

		// remove field
		collection.Fields.RemoveById("_clone_hMsz")

		// remove field
		collection.Fields.RemoveById("_clone_MOLJ")

		// remove field
		collection.Fields.RemoveById("_clone_gmYX")

		// remove field
		collection.Fields.RemoveById("_clone_yP7e")

		// remove field
		collection.Fields.RemoveById("_clone_XKH4")

		// remove field
		collection.Fields.RemoveById("_clone_UdEV")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_64ua",
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
			"id": "_clone_7eLI",
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
			"id": "_clone_ghtp",
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
			"id": "_clone_Zi7B",
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
			"id": "_clone_99GT",
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
			"id": "_clone_qxGk",
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
		if err := collection.Fields.AddMarshaledJSONAt(8, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_iRBV",
			"max": 0,
			"min": 0,
			"name": "external_ref_name",
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
			"id": "_clone_r6Ht",
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
			"viewQuery": "SELECT d.id, artifacts.id artifact_id, d.name domain, d.value ip_address, cr.fingerprint, cr.subject, cr.alternative_names, external_references.value external_ref_id, external_references.name external_ref_name, org.id organization, cr.id certificate, d.root_domain\nFROM dns_records d\n         INNER JOIN org_domains o ON o.root_domain=d.root_domain\n         INNER JOIN organizations org on o.organization=org.id\n         INNER JOIN ip_addresses a ON a.address = d.value\n         INNER JOIN ip_ports p ON a.id = p.ip_address AND p.last_seen > datetime('now', '-8 hour')\n         INNER JOIN ip_port_certificates c on p.id = c.ip_port AND c.last_seen > datetime('now', '-8 hour')\n         INNER JOIN certificates cr on cr.id = c.certificate\n         LEFT JOIN org_certificates oc ON oc.certificate=c.certificate\n         LEFT JOIN artifacts on artifacts.dns_record=d.id\n         LEFT JOIN external_references on external_references.id=d.external_reference\nWHERE d.last_resolved > datetime('now', '-8 hour')\n  AND oc.id IS NULL;"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_gJoB",
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
			"id": "_clone_2rR6",
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
			"id": "_clone_hMsz",
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
			"id": "_clone_MOLJ",
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
			"id": "_clone_gmYX",
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
			"id": "_clone_yP7e",
			"max": 0,
			"min": 0,
			"name": "external_ref_id",
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
			"id": "_clone_XKH4",
			"max": 0,
			"min": 0,
			"name": "external_ref_name",
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
			"id": "_clone_UdEV",
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
		collection.Fields.RemoveById("_clone_64ua")

		// remove field
		collection.Fields.RemoveById("_clone_7eLI")

		// remove field
		collection.Fields.RemoveById("_clone_ghtp")

		// remove field
		collection.Fields.RemoveById("_clone_Zi7B")

		// remove field
		collection.Fields.RemoveById("_clone_99GT")

		// remove field
		collection.Fields.RemoveById("_clone_qxGk")

		// remove field
		collection.Fields.RemoveById("_clone_iRBV")

		// remove field
		collection.Fields.RemoveById("_clone_r6Ht")

		return app.Save(collection)
	})
}
