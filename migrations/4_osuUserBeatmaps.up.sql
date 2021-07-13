create table osu_user_beatmaps
(
    uuid uuid default uuid_generate_v4()
        constraint osu_user_beatmaps_pk
            primary key,
    "user" int not null,
    user_name text,
    beatmap_type int not null,
    beatmap_data json not null
);

create table osu_user_recent_activity
(
    uuid uuid default uuid_generate_v4()
        constraint osu_user_recent_activity_pk
            primary key,
    user_id int not null,
    last_activity_id int not null
);

create unique index osu_user_recent_activity_user_id_uindex
	on osu_user_recent_activity (user_id);