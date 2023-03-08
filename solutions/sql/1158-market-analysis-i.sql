# Write your MySQL query statement below
SELECT
    user_id AS buyer_id,
    join_date,
    COUNT(order_id) AS orders_in_2019
FROM Users u
LEFT JOIN Orders o ON o.buyer_id = user_id AND YEAR(order_date) = '2019'
GROUP BY user_id;
