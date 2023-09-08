-- migrate:up
create table users
(
    id         integer primary key autoincrement,
    name       text not null,
    email      text not null,
    password   text not null,
    created_at datetime default current_timestamp
);

-- migrate:down
drop table users
