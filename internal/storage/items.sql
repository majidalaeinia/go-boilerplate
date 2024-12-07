-- name: CreateItem :one
insert into items (name)
values (@name)
returning id;

-- name: GetItem :one
select *
from items
where id = @id;
