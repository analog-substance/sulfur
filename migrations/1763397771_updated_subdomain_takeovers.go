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
			"viewQuery": "SELECT DISTINCT d.id, org.id org_id, org.name org_name, d.name domain, d.value ip_address, d.last_resolved\nFROM dns_records d\n         INNER JOIN org_domains o ON o.root_domain=d.root_domain\n         INNER JOIN organizations org on o.organization=org.id\n         INNER JOIN ip_addresses a ON a.address = d.value\n\n\nLEFT JOIN org_ip_addresses on a.id=org_ip_addresses.ip_address AND org.id=org_ip_addresses.organization and org_ip_addresses.last_seen > datetime('now', '-24 hour')\n\nLEFT JOIN ignored_subdomain_takeovers i ON d.name=i.domain and org.id = i.org\n  \nWHERE d.last_resolved > datetime('now', '-7 days')\n    AND a.is_private = false\n    AND a.is_loopback = false\n    AND org_ip_addresses.id IS NULL\n  AND i.id IS NULL\n\n;"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_IyZD")

		// remove field
		collection.Fields.RemoveById("_clone_qmeD")

		// remove field
		collection.Fields.RemoveById("_clone_twsE")

		// remove field
		collection.Fields.RemoveById("_clone_abII")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_Yh4r",
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
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_wshi",
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
			"id": "_clone_eqea",
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
			"id": "_clone_wmY9",
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
			"viewQuery": "SELECT DISTINCT d.id, org.id org_id, org.name org_name, d.name domain, d.value ip_address, d.last_resolved\nFROM dns_records d\n         INNER JOIN org_domains o ON o.root_domain=d.root_domain\n         INNER JOIN organizations org on o.organization=org.id\n         INNER JOIN ip_addresses a ON a.address = d.value\n\n\nLEFT JOIN org_ip_addresses on a.id=org_ip_addresses.ip_address AND org.id=org_ip_addresses.organization and org_ip_addresses.last_seen > datetime('now', '-24 hour')\n\nWHERE d.last_resolved > datetime('now', '-7 days')\n    AND a.is_private = false\n    AND a.is_loopback = false\n    AND org_ip_addresses.id IS NULL\n\n;"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_IyZD",
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
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_qmeD",
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
			"id": "_clone_twsE",
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
			"id": "_clone_abII",
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
		collection.Fields.RemoveById("_clone_Yh4r")

		// remove field
		collection.Fields.RemoveById("_clone_wshi")

		// remove field
		collection.Fields.RemoveById("_clone_eqea")

		// remove field
		collection.Fields.RemoveById("_clone_wmY9")

		return app.Save(collection)
	})
}
