WITH tmp as (
    SELECT account, SUM(amount) as balance
    FROM Transactions
    GROUP BY account
    HAVING SUM(amount) > 10000
)
SELECT name, balance
FROM tmp t JOIN Users u ON t.account = u.account
