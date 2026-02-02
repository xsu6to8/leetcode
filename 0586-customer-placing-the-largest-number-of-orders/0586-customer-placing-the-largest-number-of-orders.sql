with tmp as (
    select customer_number, count(customer_number) as cnt
    from Orders
    group by customer_number 
)
select customer_number 
from tmp
WHERE cnt = (select MAX(cnt) from tmp);