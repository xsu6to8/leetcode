# Write your MySQL query statement below
with tmp as (
    select salary, RANK() OVER(ORDER BY salary DESC) as rnk
    from Employee
    group by salary
)
select MAX(salary) as SecondHighestSalary
from tmp
where rnk = 2;