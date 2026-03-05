WITH tmp as (
    SELECT p.product_name as product_name, unit
    FROM Products p JOIN Orders o 
        ON p.product_id = o.product_id 
    where order_date BETWEEN '2020-02-01' AND '2020-02-29'
)
select product_name, sum(unit) as unit
from tmp
group by product_name
having sum(unit) >= 100