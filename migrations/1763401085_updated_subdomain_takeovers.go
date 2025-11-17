package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2813928730")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT DISTINCT d.id, org.id org_id, org.name org_name, artifacts.id artifact_id, d.name domain, d.value ip_address, d.last_resolved\nFROM dns_records d\n         INNER JOIN org_domains o ON o.root_domain=d.root_domain\n         INNER JOIN organizations org on o.organization=org.id\n         INNER JOIN ip_addresses a ON a.address = d.value\n\n\nLEFT JOIN org_ip_addresses on a.id=org_ip_addresses.ip_address AND org.id=org_ip_addresses.organization and org_ip_addresses.last_seen > datetime('now', '-24 hour')\n\nLEFT JOIN ignored_subdomain_takeovers i ON d.name=i.domain and org.id = i.org\n    LEFT JOIN artifacts on artifacts.dns_record=d.id\nWHERE d.last_resolved > datetime('now', '-7 days')\n    AND a.is_private = false\n    AND a.is_loopback = false\n    AND org_ip_addresses.id IS NULL\n  AND i.id IS NULL\n\n;"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_xalM")

		// remove field
		collection.Fields.RemoveById("relation955505121")

		// remove field
		collection.Fields.RemoveById("_clone_n4B3")

		// remove field
		collection.Fields.RemoveById("_clone_HbNA")

		// remove field
		collection.Fields.RemoveById("_clone_u3J1")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_FGFh",
			"max": 0,
			"min": 0,
			"name": "org_name",
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
			"cascadeDelete": false,
			"collectionId": "pbc_1303748624",
			"hidden": false,
			"id": "relation3800762284",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "artifact_id",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_Uq3F",
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
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_afYh",
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
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"hidden": false,
			"id": "_clone_fBHp",
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

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2813928730")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT DISTINCT d.id, org.id org_id, org.name org_name, artifacts.id artifacts_id, d.name domain, d.value ip_address, d.last_resolved\nFROM dns_records d\n         INNER JOIN org_domains o ON o.root_domain=d.root_domain\n         INNER JOIN organizations org on o.organization=org.id\n         INNER JOIN ip_addresses a ON a.address = d.value\n\n\nLEFT JOIN org_ip_addresses on a.id=org_ip_addresses.ip_address AND org.id=org_ip_addresses.organization and org_ip_addresses.last_seen > datetime('now', '-24 hour')\n\nLEFT JOIN ignored_subdomain_takeovers i ON d.name=i.domain and org.id = i.org\n    LEFT JOIN artifacts on artifacts.dns_record=d.id\nWHERE d.last_resolved > datetime('now', '-7 days')\n    AND a.is_private = false\n    AND a.is_loopback = false\n    AND org_ip_addresses.id IS NULL\n  AND i.id IS NULL\n\n;"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_xalM",
			"max": 0,
			"min": 0,
			"name": "org_name",
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
			"cascadeDelete": false,
			"collectionId": "pbc_1303748624",
			"hidden": false,
			"id": "relation955505121",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "artifacts_id",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_n4B3",
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
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_HbNA",
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
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"hidden": false,
			"id": "_clone_u3J1",
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

		// remove field
		collection.Fields.RemoveById("_clone_FGFh")

		// remove field
		collection.Fields.RemoveById("relation3800762284")

		// remove field
		collection.Fields.RemoveById("_clone_Uq3F")

		// remove field
		collection.Fields.RemoveById("_clone_afYh")

		// remove field
		collection.Fields.RemoveById("_clone_fBHp")

		return app.Save(collection)
	})
}
