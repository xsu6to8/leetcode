with tmp as (
    select query_name, rating / position as own_q, rating
    from Queries
)
select query_name, 
    ROUND(sum(own_q)/count(*), 2) as quality,
    ROUND((select count(*) from tmp t2 where t1.query_name = t2.query_name AND rating < 3)/count(*) * 100, 2) as poor_query_percentage
from tmp t1
group by query_name