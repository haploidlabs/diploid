-- name: CreateUser :one
insert into users (name, email, password, role)
values (?, ?, ?, ?)
returning *;

-- name: GetUserByID :one
select *
from users
where id = ?;

-- name: GetUserByEmail :one
select *
from users
where email = ?;

-- name: CountUsers :one
select count(*) from users;
