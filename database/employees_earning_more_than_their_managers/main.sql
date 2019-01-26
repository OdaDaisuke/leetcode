# Write your MySQL query statement below

SELECT first.Name As Employee FROM Employee as first
LEFT JOIN Employee as second ON second.Id = first.ManagerID
WHERE first.Salary > second.Salary AND first.ManagerId IS NOT NULL;
