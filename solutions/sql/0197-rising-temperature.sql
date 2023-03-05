# Write your MySQL query statement below
SELECT w.id
FROM Weather w
INNER JOIN Weather t ON t.recordDate = DATE_SUB(w.recordDate, INTERVAL 1 DAY)
WHERE w.temperature > t.temperature;
