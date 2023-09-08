CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE users
(
    id         integer primary key autoincrement,
    name       text not null,
    email      text not null,
    password   text not null,
    created_at datetime default current_timestamp
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20230831194609');
