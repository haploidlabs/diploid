-- name: CreateProject :one
insert into projects (name, description, created_by)
values (?, ?, ?)
returning *;

-- name: GetProjectByID :one
select *
from projects
where id = ?;

-- name: GetProjectsByUser :many
select *
from projects
where created_by = ?;

-- name: GetProjectByName :many
select *
from projects
where name = ?
  and created_by = ?;

-- name: UpdateProject :one
update projects
set name        = ?,
    description = ?
where id = ?
returning *;
