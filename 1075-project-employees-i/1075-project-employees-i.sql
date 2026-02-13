select project_id, ROUND(AVG(experience_years), 2) average_years
from Project NATURAL JOIN Employee
group by project_id