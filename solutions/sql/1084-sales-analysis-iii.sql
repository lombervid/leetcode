# Write your MySQL query statement below
SELECT
    p.product_id,
    p.product_name
FROM Product p
INNER JOIN Sales s ON s.product_id = p.product_id
GROUP BY p.product_id
HAVING MIN(sale_date) >= '2019-01-01' AND MAX(sale_date) <= '2019-03-31';

-- Another solution
SELECT
    DISTINCT p.product_id,
    product_name
FROM Product p
INNER JOIN Sales s ON s.product_id = p.product_id
WHERE p.product_id NOT IN(
    SELECT DISTINCT product_id
    FROM Sales
    WHERE sale_date NOT BETWEEN '2019-01-01' AND '2019-03-31'
);
