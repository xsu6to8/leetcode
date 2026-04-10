SELECT id, 
    CASE 
        WHEN t1.p_id IS null THEN 'Root'
        WHEN t1.p_id IS NOT null AND EXISTS (
            SELECT *
            FROM Tree t2
            WHERE t1.id IN (
                t2.p_id
            )
        ) THEN 'Inner'
        ELSE 'Leaf'
        END AS 'type'
FROM Tree t1
