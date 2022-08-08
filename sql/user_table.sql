set search_path to "gin";

create table if not exists "user" (
    id integer not null,
    email character varying not null unique,
    password character varying not null,
    name character varying,
    updated_at integer,
    created_at integer,
    CONSTRAINT user_id PRIMARY KEY(id)
);

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE user_id_seq OWNED BY "user".id;

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);

SELECT pg_catalog.setval('user_id_seq', 1, false);

insert into gin.user(email, password, name) values ('wltjd9870@naver.com', 'test', 'test') returning id;
delete from gin.user where email = 'wltjd9870@naver.com';
select * from gin.user;
