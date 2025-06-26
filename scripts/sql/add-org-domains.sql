INSERT INTO org_domains (root_domain, organization, description, created, updated)
SELECT root_domains.id root_domain, "changeme" org,  "auto imported", datetime('now'), datetime('now')
FROM root_domains
LEFT JOIN org_domains ON org_domains.root_domain=root_domains.id
WHERE domain LIKE '%changeme%'
AND org_domains.id IS NULL