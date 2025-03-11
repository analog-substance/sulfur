package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"createRule": null,
			"deleteRule": null,
			"fields": [
				{
					"autogeneratePattern": "[a-z0-9]{15}",
					"hidden": false,
					"id": "text3208210256",
					"max": 15,
					"min": 15,
					"name": "id",
					"pattern": "^[a-z0-9]+$",
					"presentable": false,
					"primaryKey": true,
					"required": true,
					"system": true,
					"type": "text"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_548419197",
					"hidden": false,
					"id": "relation2425873977",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "env_cidr",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text223244161",
					"max": 0,
					"min": 0,
					"name": "address",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "bool3276039582",
					"name": "is_6",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "bool"
				},
				{
					"hidden": false,
					"id": "bool4213212674",
					"name": "is_global_unicast",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "bool"
				},
				{
					"hidden": false,
					"id": "bool437070178",
					"name": "is_interface_local_multicast",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "bool"
				},
				{
					"hidden": false,
					"id": "bool95852846",
					"name": "is_link_local_multicast",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "bool"
				},
				{
					"hidden": false,
					"id": "bool3706817414",
					"name": "is_link_local_unicast",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "bool"
				},
				{
					"hidden": false,
					"id": "bool1187734526",
					"name": "is_loopback",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "bool"
				},
				{
					"hidden": false,
					"id": "bool3541473878",
					"name": "is_multicast",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "bool"
				},
				{
					"hidden": false,
					"id": "bool3292153625",
					"name": "is_private",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "bool"
				},
				{
					"hidden": false,
					"id": "bool1461063986",
					"name": "is_unspecified",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "bool"
				},
				{
					"hidden": false,
					"id": "bool1137015140",
					"name": "is_shared",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "bool"
				},
				{
					"hidden": false,
					"id": "bool3120091063",
					"name": "is_ephemeral",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "bool"
				},
				{
					"hidden": false,
					"id": "date846843460",
					"max": "",
					"min": "",
					"name": "last_seen",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "date"
				},
				{
					"hidden": false,
					"id": "autodate2990389176",
					"name": "created",
					"onCreate": true,
					"onUpdate": false,
					"presentable": false,
					"system": false,
					"type": "autodate"
				},
				{
					"hidden": false,
					"id": "autodate3332085495",
					"name": "updated",
					"onCreate": true,
					"onUpdate": true,
					"presentable": false,
					"system": false,
					"type": "autodate"
				}
			],
			"id": "pbc_2789751422",
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_5Ny6Y9gYHc` + "`" + ` ON ` + "`" + `ip_addresses` + "`" + ` (\n  ` + "`" + `is_private` + "`" + `,\n  ` + "`" + `is_shared` + "`" + `,\n  ` + "`" + `is_loopback` + "`" + `,\n  ` + "`" + `is_multicast` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_jAtgJ7olEN` + "`" + ` ON ` + "`" + `ip_addresses` + "`" + ` (\n  ` + "`" + `is_ephemeral` + "`" + `,\n  ` + "`" + `last_seen` + "`" + `\n)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_kiyJRtxwRC` + "`" + ` ON ` + "`" + `ip_addresses` + "`" + ` (` + "`" + `address` + "`" + `)"
			],
			"listRule": null,
			"name": "ip_addresses",
			"system": false,
			"type": "base",
			"updateRule": null,
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2789751422")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
