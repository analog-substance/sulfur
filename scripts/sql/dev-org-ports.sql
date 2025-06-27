SELECT organizations.name, ip_addresses.address, ip_ports.port, ip_ports.protocol
FROM ip_ports
INNER JOIN ip_addresses on ip_ports.ip_address=ip_addresses.id
INNER JOIN org_ip_addresses ON org_ip_addresses.ip_address=ip_addresses.id
INNER JOIN organizations on organizations.id=org_ip_addresses.organization
WHERE ip_ports.last_seen > datetime('now', '-1 day');


SELECT organizations.name, ip_addresses.address, ip_ports.port, ip_ports.protocol, group_concat(distinct dns_records.name)
FROM ip_ports
         INNER JOIN ip_addresses on ip_ports.ip_address = ip_addresses.id
         INNER JOIN dns_records ON dns_records.value = ip_addresses.address
         LEFT JOIN dns_records d2 ON instr(d2.name, dns_records.name) > 0
         LEFT JOIN org_domains  ON org_domains.root_domain=dns_records.root_domain
         LEFT JOIN organizations on organizations.id=org_domains.organization
WHERE ip_ports.last_seen > datetime('now', '-1 day')
  AND (organizations.id IS NOT NULL )
GROUP BY  ip_addresses.address;



SELECT DISTINCT organizations.name, ip_addresses.address, ip_ports.port, ip_ports.protocol
FROM ip_ports
         INNER JOIN ip_addresses on ip_ports.ip_address = ip_addresses.id
         INNER JOIN dns_records ON dns_records.value = ip_addresses.address
         LEFT JOIN dns_records d2 ON instr(d2.name, dns_records.name) > 0
         LEFT JOIN org_domains  ON org_domains.root_domain=dns_records.root_domain
         LEFT JOIN organizations on organizations.id=org_domains.organization
WHERE ip_ports.last_seen > datetime('now', '-1 day')
  AND (organizations.id IS NOT NULL )
;


SELECT DISTINCT organizations.name, ip_ports.port, count(distinct ip_ports.id)
FROM ip_ports
         INNER JOIN ip_addresses on ip_ports.ip_address = ip_addresses.id
         INNER JOIN dns_records ON dns_records.value = ip_addresses.address
         LEFT JOIN dns_records d2 ON instr(d2.name, dns_records.name) > 0
         LEFT JOIN org_domains  ON org_domains.root_domain=dns_records.root_domain
         LEFT JOIN organizations on organizations.id=org_domains.organization
WHERE ip_ports.last_seen > datetime('now', '-1 day')
  AND (organizations.id IS NOT NULL )
GROUP BY organizations.name, ip_ports.port
;
