with tmp as (
    select e.name as name, b.bonus as bonus
    from Employee e LEFT JOIN
    Bonus b ON e.empId = b.empId
)
select * 
from tmp
where bonus IS NULL OR bonus < 1000