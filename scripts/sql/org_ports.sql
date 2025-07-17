SELECT org_ip_addresses.organization, ip_ports.port, count(*) total
FROM org_ip_addresses
         INNER JOIN ip_ports ON ip_ports.ip_address = org_ip_addresses.ip_address and ip_ports.last_seen > datetime('now', '-24 hours')
WHERE org_ip_addresses.last_seen > datetime('now', '-24 hours')
GROUP BY org_ip_addresses.organization, ip_ports.port
ORDER BY TOTAL desc;


SELECT org_ip_addresses.organization, ip_addresses.address, ip_ports.port, org_ip_addresses.last_seen, group_concat(dns_records.name)
FROM org_ip_addresses
         INNER JOIN ip_addresses ON ip_addresses.id = org_ip_addresses.ip_address
         INNER JOIN ip_ports ON ip_ports.ip_address = org_ip_addresses.ip_address and ip_ports.last_seen > datetime('now', '-24 hours')
        LEFT JOIN dns_records ON dns_records.value=ip_addresses.address AND dns_records.type="A"
WHERE org_ip_addresses.last_seen > datetime('now', '-24 hours')
GROUP BY org_ip_addresses.organization, ip_addresses.id, ip_ports.port
ORDER BY org_ip_addresses.last_seen desc, ip_ports.port asc;