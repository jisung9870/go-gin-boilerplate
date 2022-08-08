-- User 조회
select * from PG_USER;

-- User 생성 
create user gin_admin with password 'ginadmin';
create user gin_client with password 'ginclient';

create database gin_db with owner=gin_admin;

revoke all on database gin_db from public;
revoke all on schema public from public;

create schema gin authorization gin_admin;

grant connect,temporary on database gin_db to gin_client;
grant usage on schema gin to gin_client;
alter role gin_client set search_path to gin;
grant select, insert, update, delete on all tables in schema gin to gin_client;
alter default privileges in schema gin grant select, insert, update, delete on tables to gin_client;
grant usage on all sequences in schema gin to gin_client;
alter default privileges in schema gin grant usage on sequences to gin_client;
