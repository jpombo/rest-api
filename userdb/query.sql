-- name: GetUser :one
select * from tbusers where id = ?;

-- name: ListUsers :many
select id, name, email, birthdate
from tbusers
order by name;

-- name: CreateUser :execresult
insert into tbusers (
    id, name, email, birthdate
) 
values(
    ?, ?, ?, ?
);

-- name: DeleteUser :execrows
delete from tbusers 
where id = ?;

-- name: UpdateUser :execrows
update tbusers 
set name = ?, birthdate = ?
where id = ?;

-- name: CheckEmail :one
select 1
  from tbusers
 where email = ?;

-- name: GetProduct :one
select * 
  from tbproducts 
 where id = ?;

-- name: ListProducts :many
select id, descricao, categoria
from tbproducts
order by descricao;

-- name: CreateProduc :execresult
insert into tbproducts (
    id, descricao, categoria
) 
values(
    ?, ?, ?
);

-- name: DeleteProduct :execrows
delete from tbproducts 
where id = ?;

-- name: UpdateProduct :execrows
update tbproducts 
set categoria = ?
where id = ?;

-- name: CheckDescricao :one
select 1
  from tbproducts
 where descricao = ?;