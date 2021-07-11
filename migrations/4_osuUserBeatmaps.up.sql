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