with tmp as (
    select sell_date, count(distinct product) as num_sold, group_concat(distinct product order by product ASC) as products     
    from Activities
    group by sell_date
)
select *
from tmp