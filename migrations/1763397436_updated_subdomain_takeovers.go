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
			"viewQuery": "SELECT DISTINCT d.id, org.id org_id, org.name org_name, d.name domain, d.value ip_address, d.last_resolved\nFROM dns_records d\n         INNER JOIN org_domains o ON o.root_domain=d.root_domain\n         INNER JOIN organizations org on o.organization=org.id\n         INNER JOIN ip_addresses a ON a.address = d.value\n\n\nLEFT JOIN org_ip_addresses on a.id=org_ip_addresses.ip_address AND org.id=org_ip_addresses.organization and org_ip_addresses.last_seen > datetime('now', '-24 hour')\n\nWHERE d.last_resolved > datetime('now', '-7 days')\n    AND a.is_private = false\n    AND a.is_loopback = false\n    AND org_ip_addresses.id IS NULL\n\n;"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_h37A")

		// remove field
		collection.Fields.RemoveById("_clone_FB9h")

		// remove field
		collection.Fields.RemoveById("_clone_AOsp")

		// remove field
		collection.Fields.RemoveById("_clone_3nTG")

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

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2813928730")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT DISTINCT d.id, org.id org_id, org.name org_name, d.name domain, d.value ip_address, d.last_resolved\nFROM dns_records d\n         INNER JOIN org_domains o ON o.root_domain=d.root_domain\n         INNER JOIN organizations org on o.organization=org.id\n         INNER JOIN ip_addresses a ON a.address = d.value\n\n\nLEFT JOIN ip_ports on ip_ports.ip_address=a.id and ip_ports.last_seen > datetime('now', '-24 hour') AND ip_ports.port=443\nLEFT JOIN ip_port_certificates on ip_port_certificates.ip_port=ip_ports.id\nLEFT JOIN org_ip_addresses on org.id=org_ip_addresses.organization and org_ip_addresses.last_seen > datetime('now', '-24 hour') and a.id=org_ip_addresses.ip_address\nLEFT JOIN org_certificates on org.id=org_certificates.organization and ip_port_certificates.certificate=org_certificates.certificate\n\n\nWHERE d.last_resolved > datetime('now', '-24 hour')\n    AND a.is_private = false\n    AND a.is_loopback = false\n    AND org_ip_addresses.id IS NULL\n    AND ip_ports.id IS NOT NULL\n    AND ip_port_certificates.id IS NOT NULL\n    AND org_certificates.id IS NULL\n\n;"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_h37A",
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
			"id": "_clone_FB9h",
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
			"id": "_clone_AOsp",
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
			"id": "_clone_3nTG",
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
		collection.Fields.RemoveById("_clone_IyZD")

		// remove field
		collection.Fields.RemoveById("_clone_qmeD")

		// remove field
		collection.Fields.RemoveById("_clone_twsE")

		// remove field
		collection.Fields.RemoveById("_clone_abII")

		return app.Save(collection)
	})
}
