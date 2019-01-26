# Write your MySQL query statement below

DELETE first FROM Person as first
JOIN Person as second
ON first.Email = second.Email WHERE second.Id < first.Id;
