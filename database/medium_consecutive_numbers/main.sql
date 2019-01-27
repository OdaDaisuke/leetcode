/* Write your T-SQL query statement below */

SELECT DISTINCT l1.Num as ConsecutiveNums
FROM Logs as l1, Logs as l2, Logs as l3
WHERE
    (l1.Num = l2.Num AND l1.id = (l2.Id + 1)) AND
    (l2.Num = l3.Num AND l2.Id = (l3.Id + 1))
