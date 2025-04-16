-- name: CreateUser :one
insert into users (name)
values (@name)
    returning id;

-- name: GetUser :one
select *
from users
where id = @id;

-- name: GetUsers :many
select *
from users;
