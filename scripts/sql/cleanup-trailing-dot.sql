SELECT COUNT(*) FROM dns_records
WHERE name LIKE '%.';

UPDATE dns_records
SET name = rtrim(name, ".")
WHERE name LIKE '%.';

DELETE FROM dns_records
WHERE name LIKE '%.';