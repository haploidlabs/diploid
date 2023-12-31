CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE users
(
    id         integer primary key autoincrement,
    name       text not null,
    email      text not null,
    password   text not null,
    role       text not null,
    created_at datetime default current_timestamp
);
CREATE TABLE user_invites
(
    id         integer primary key autoincrement,
    email      text not null,
    token      text not null,
    created_at datetime default current_timestamp
);
CREATE TABLE projects
(
    id          integer primary key autoincrement,
    name        text    not null,
    description text,
    created_by  integer not null references users (id),
    created_at  datetime default current_timestamp
);
CREATE TABLE project_members
(
    id         integer primary key autoincrement,
    project_id integer not null references projects (id),
    user_id    integer not null references users (id),
    role       text    not null,
    created_at timestamp default current_timestamp
);
CREATE TABLE services
(
    id             integer primary key autoincrement,
    name           text    not null,
    environment_id integer not null references environments (id),
    created_at     timestamp default current_timestamp
);
CREATE TABLE environments
(
    id          integer not null primary key autoincrement,
    project_id  integer not null references projects (id),
    name        text    not null unique,
    description text
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20230831194609'),
  ('20230927181011');
