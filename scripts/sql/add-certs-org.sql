INSERT INTO org_certificates (organization, certificate, created, updated)
SELECT org_domains.organization, certificates.id, datetime('now'), datetime('now')
FROM certificates
         INNER JOIN root_domains ON instr(certificates.subject, root_domains.domain) > 0 OR instr(certificates.alternative_names, root_domains.domain) > 0
         INNER JOIN org_domains ON root_domains.id = org_domains.root_domain
        LEFT JOIN org_certificates ON org_certificates.certificate=certificates.id
WHERE org_certificates.id IS NULL
GROUP BY certificates.id
