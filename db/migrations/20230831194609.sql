-- migrate:up
create table users
(
    id         integer primary key autoincrement,
    name       text not null,
    email      text not null,
    password   text not null,
    role       text not null,
    created_at datetime default current_timestamp
);

create table user_invites
(
    id         integer primary key autoincrement,
    email      text not null,
    token      text not null,
    created_at datetime default current_timestamp
);

create table projects
(
    id          integer primary key autoincrement,
    name        text    not null,
    description text,
    created_by  integer not null references users (id),
    created_at  datetime default current_timestamp
);

create table project_members
(
    id         integer primary key autoincrement,
    project_id integer not null references projects (id),
    user_id    integer not null references users (id),
    role       text    not null,
    created_at timestamp default current_timestamp
);

create table services
(
    id             integer primary key autoincrement,
    name           text    not null,
    environment_id integer not null references environments (id),
    created_at     timestamp default current_timestamp
);

-- migrate:down
drop table users;
drop table user_invites;
drop table projects;
drop table project_members;
drop table services;
