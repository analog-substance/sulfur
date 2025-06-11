package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_4107852781")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT 1 as id,\n  (SELECT COUNT(*) FROM dns_records WHERE last_resolved > DATETIME('now', '-8 hours')) active_dns_records,\n (SELECT COUNT(*) FROM ip_ports WHERE last_seen > DATETIME('now', '-8 hours')) active_ports\n"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_4107852781")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT 1 as id,\n  (SELECT COUNT(*) FROM dns_records WHERE last_seen > DATETIME('now', '-8 hours')) active_dns_records,\n (SELECT COUNT(*) FROM ip_ports WHERE last_seen > DATETIME('now', '-8 hours')) active_ports\n"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
