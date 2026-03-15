-- name: Get :one
select * from tbusers where id = ?;

-- name: List :many
select id, name, email, birthdate
from tbusers
order by name;

-- name: Create :execresult
insert into tbusers (
    id, name, email, birthdate
) 
values(
    ?, ?, ?, ?
);

-- name: Delete :exec
delete from tbusers 
where id = ?;

-- name: Update :exec
update tbusers 
set name = ?, email = ?, birthdate = ?
where id = ?;

-- name: CheckEmail :one
select 1
  from tbusers
 where email = ?;