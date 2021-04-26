create table calendar
(
    uuid               uuid default uuid_generate_v4()
        constraint calendar_pk
            primary key,
    name               text not null,
    description        text,
    creator_discord_id text not null
);

create table event
(
    uuid               uuid default uuid_generate_v4()
        constraint event_pk
            primary key,
    calendar_uuid      uuid not null,
    rrule              text,
    title              text not null,
    description        text,
    type               int  not null,
    data               json,
    creator_discord_id text,
    parent_event       uuid
);

alter table event
    add constraint event_event_uuid_fk
        foreign key (parent_event) references event
            on delete cascade;

alter table event
    add start_date timestamp with time zone;

alter table event
    alter column start_date set not null;

alter table event
    alter column rrule set not null;