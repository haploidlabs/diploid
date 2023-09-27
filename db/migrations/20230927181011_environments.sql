-- migrate:up
create table environments
(
    id          integer not null primary key autoincrement,
    project_id  integer not null references projects (id),
    name        text    not null unique,
    description text
);

-- migrate:down
drop table environments;
