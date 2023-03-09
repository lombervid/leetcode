# Write your MySQL query statement below
SELECT
    name,
    SUM(amount) AS balance
FROM Users u
INNER JOIN Transactions t ON t.account = u.account
GROUP BY u.account
HAVING balance > 10000;
