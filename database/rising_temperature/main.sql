# Write your MySQL query statement below

SELECT Weather.Id FROM Weather
JOIN Weather as second
WHERE DATEDIFF(Weather.RecordDate, second.RecordDate) = 1
AND Weather.Temperature > second.Temperature;
