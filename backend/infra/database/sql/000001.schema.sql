create table t_material
(
    id         serial primary key,
    created_at timestamp    not null default current_timestamp,
    updated_at timestamp    not null default current_timestamp,

    key        varchar(255) not null,
    name       varchar(255) not null,
    extend     text         not null
);
