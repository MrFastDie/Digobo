CREATE TABLE public.calendar (
                                 uuid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
                                 name text NOT NULL,
                                 description text,
                                 creator_discord_id text NOT NULL
);

CREATE TABLE public.event (
                              uuid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
                              calendar_uuid uuid NOT NULL,
                              rrule text NOT NULL,
                              title text NOT NULL,
                              description text,
                              type integer NOT NULL,
                              data json,
                              creator_discord_id text,
                              parent_event uuid,
                              start_date timestamp with time zone NOT NULL
);

CREATE TABLE public.osu_user_beatmaps (
                                          uuid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
                                          "user" integer NOT NULL,
                                          beatmap_type integer NOT NULL,
                                          beatmap_data json NOT NULL,
                                          user_name text
);

CREATE TABLE public.osu_user_recent_activity (
                                                 uuid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
                                                 user_id integer NOT NULL,
                                                 last_activity_id integer NOT NULL
);

CREATE TABLE public.osu_user_watcher (
                                         user_id integer NOT NULL,
                                         user_name text NOT NULL,
                                        color integer NULL
);

CREATE TABLE public.osu_user_watcher_channel (
                                                 user_id integer NOT NULL,
                                                 channel_id text NOT NULL,
                                                 color integer NULL
);

CREATE TABLE public.random_answer_list (
                                           uuid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
                                           command text NOT NULL,
                                           value text NOT NULL,
                                           discord_id text DEFAULT ''::text NOT NULL
);

ALTER TABLE ONLY public.calendar
    ADD CONSTRAINT calendar_pk PRIMARY KEY (uuid);


ALTER TABLE ONLY public.event
    ADD CONSTRAINT event_pk PRIMARY KEY (uuid);


ALTER TABLE ONLY public.osu_user_beatmaps
    ADD CONSTRAINT osu_user_beatmaps_pk PRIMARY KEY (uuid);


ALTER TABLE ONLY public.osu_user_recent_activity
    ADD CONSTRAINT osu_user_recent_activity_pk PRIMARY KEY (uuid);


ALTER TABLE ONLY public.osu_user_watcher_channel
    ADD CONSTRAINT osu_user_watcher_channel_pk PRIMARY KEY (user_id, channel_id);

ALTER TABLE ONLY public.osu_user_watcher
    ADD CONSTRAINT osu_user_watcher_pk PRIMARY KEY (user_id);

ALTER TABLE ONLY public.random_answer_list
    ADD CONSTRAINT random_answer_list_pk PRIMARY KEY (uuid);

ALTER TABLE ONLY public.random_answer_list
    ADD CONSTRAINT random_answer_list_pk_2 UNIQUE (command, value);

CREATE UNIQUE INDEX osu_user_recent_activity_user_id_uindex ON public.osu_user_recent_activity USING btree (user_id);

ALTER TABLE ONLY public.event
    ADD CONSTRAINT event_event_uuid_fk FOREIGN KEY (parent_event) REFERENCES public.event(uuid) ON DELETE CASCADE;

ALTER TABLE ONLY public.osu_user_recent_activity
    ADD CONSTRAINT osu_user_recent_activity_osu_user_watcher_user_id_fk FOREIGN KEY (user_id) REFERENCES public.osu_user_watcher(user_id);

ALTER TABLE ONLY public.osu_user_watcher_channel
    ADD CONSTRAINT osu_user_watcher_channel_osu_user_watcher_user_id_fk FOREIGN KEY (user_id) REFERENCES public.osu_user_watcher(user_id);


--
-- PostgreSQL database dump complete
--

