-- name: GetEnvironmentsForProject :many
select *
from environments
where project_id = ?;

-- name: GetEnvironmentById :one
select *
from environments
where id = ?;

-- name: GetEnvironmentByName :one
select *
from environments
where name = ?
  and project_id = ?;

-- name: CreateEnvironment :one
insert into environments (project_id, name, description)
values (?, ?, ?)
returning *;

-- name: UpdateEnvironment :one
update environments
set name        = ?,
    description = ?
where id = ?
returning *;
