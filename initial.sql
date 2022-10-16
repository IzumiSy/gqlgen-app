drop table if exists user;
drop table if exists todo;

create table user(id text, name text);
create table todo(id text, text text, done integer, user_id text);
