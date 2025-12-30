# Write your MySQL query statement below
select E.name as Employee
from Employee as E JOIN Employee as M 
  ON E.managerId = M.id
where E.salary > M.salary