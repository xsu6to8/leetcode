WITH tmp as (
    SELECT p.product_id as product_id, price * units as own_sum, units
    FROM Prices p LEFT JOIN UnitsSold u 
        ON p.product_id = u.product_id 
        AND (start_date <= purchase_date AND purchase_date <= end_date) 
)
SELECT product_id, 
    IFNULL(ROUND(SUM(own_sum) / SUM(units), 2), 0) as average_price
FROM tmp
group by product_id