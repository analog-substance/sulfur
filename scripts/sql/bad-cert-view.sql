SELECT d.id, artifacts.id artifact_id, d.name domain, d.value ip_address,  cr.subject, cr.alternative_names, cr.serial, cr.issuer, cr.not_before, external_references.value external_id, external_references.name external_name, org.id organization, cr.id certificate, d.root_domain
FROM dns_records d
    INNER JOIN org_domains o ON o.root_domain=d.root_domain
    INNER JOIN organizations org on o.organization=org.id
    INNER JOIN ip_addresses a ON a.address = d.value
    INNER JOIN ip_ports p ON a.id = p.ip_address AND p.last_seen > datetime('now', '-8 hour')
    INNER JOIN ip_port_certificates c on p.id = c.ip_port AND c.last_seen > datetime('now', '-8 hour')
    INNER JOIN certificates cr on cr.id = c.certificate
    LEFT JOIN org_certificates oc ON oc.serial=cr.serial AND oc.issuer=cr.issuer AND oc.not_before=cr.not_before
    LEFT JOIN artifacts on artifacts.dns_record=d.id
    LEFT JOIN external_references on external_references.id=d.external_reference
WHERE d.last_resolved > datetime('now', '-8 hour')
  AND oc.id IS NULL;