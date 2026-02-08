with tmp as (
    select num
    from MyNumbers
    group by num
    having COUNT(num) = 1
)
select max(num) as num
from tmp