create table random_answer_list
(
    uuid       uuid default uuid_generate_v4() not null
        constraint random_answer_list_pk
            primary key,
    command    text                            not null,
    value      text                            not null,
    discord_id text default ''::text           not null,
    constraint random_answer_list_pk_2
        unique (command, value)
);