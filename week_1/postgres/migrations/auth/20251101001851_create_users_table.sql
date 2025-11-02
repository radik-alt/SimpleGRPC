-- +goose Up
create table auth(
    id serial primary key,
    name text not null,
    password text not null,
    created_at timestamp not null default now()
);

-- +goose Down
drop table auth;
