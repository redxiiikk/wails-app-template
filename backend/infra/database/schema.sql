create table t_config
(
    id         serial primary key,
    created_at timestamp    not null default current_timestamp,
    updated_at timestamp    not null default current_timestamp,

    key        varchar(255) not null,
    type       varchar(255) not null,
    value      text         not null
);
