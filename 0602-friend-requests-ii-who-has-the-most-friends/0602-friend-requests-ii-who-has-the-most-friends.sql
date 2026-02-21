with ids as (
    select requester_id as id
    from RequestAccepted

    UNION ALL

    select accepter_id as id
    from RequestAccepted
)
select id, count(*) as num
from ids
group by id
order by count(*) DESC
LIMIT 1