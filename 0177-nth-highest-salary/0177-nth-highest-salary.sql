CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
      # Write your MySQL query statement below.
    with tmp as (
        select salary, RANK() OVER(ORDER BY salary DESC) as rnk
        from Employee
        group by salary
    )
    select max(salary) as 'getNthHighestSalary(n)'
    from tmp
    where rnk = n
  );
END