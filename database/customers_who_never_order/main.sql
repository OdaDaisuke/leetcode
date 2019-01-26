# Write your MySQL query statement below

SELECT Name as Customers FROM Customers
WHERE Customers.Id NOT IN(
    SELECT Customers.Id FROM Customers RIGHT JOIN Orders ON Orders.CustomerId = Customers.ID)
