WITH tmp AS (
    SELECT player_id, MIN(event_date) AS mDate
    FROM Activity
    GROUP BY player_id
)
SELECT 
    ROUND(
        COUNT(a.player_id) / (SELECT COUNT(DISTINCT player_id) FROM Activity), 
        2
    ) AS fraction
FROM tmp t JOIN Activity a 
    ON t.player_id = a.player_id 
WHERE DATEDIFF(a.event_date, t.mDate) = 1;