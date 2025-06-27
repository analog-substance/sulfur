SELECT
    (ROW_NUMBER() OVER()) as id, COUNT(*) total, root_domain, type
from dns_records
GROUP BY root_domain, type
ORDER BY total DESC