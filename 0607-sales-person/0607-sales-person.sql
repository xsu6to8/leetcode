select name
from SalesPerson 
where sales_id NOT IN (
    select sales_id   
    from Company NATURAL JOIN Orders
    where name = 'RED'
)
