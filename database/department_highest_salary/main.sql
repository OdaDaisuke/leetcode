# Write your MySQL query statement below

SELECT Department.Name AS Department, Employee.Name AS Employee, Employee.Salary AS Salary
FROM Employee
LEFT JOIN Department ON Department.Id = Employee.DepartmentID
WHERE (Department.Id, Salary)
IN (
    SELECT DepartmentId, MAX(Salary)
    FROM Employee
    GROUP BY DepartmentId
)
