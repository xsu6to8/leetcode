with temp as (
    select email
    from Person
    group by email
    having count(*) > 1
)
select email as Email
from temp
; 