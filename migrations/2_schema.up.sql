CREATE TABLE public.calendar (
                                 uuid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
                                 name text NOT NULL,
                                 description text,
                                 creator_discord_id text NOT NULL
);


ALTER TABLE public.calendar OWNER TO digibo;

--
-- Name: event; Type: TABLE; Schema: public; Owner: digibo
--

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


ALTER TABLE public.event OWNER TO digibo;

--
-- Name: osu_user_beatmaps; Type: TABLE; Schema: public; Owner: digibo
--

CREATE TABLE public.osu_user_beatmaps (
                                          uuid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
                                          "user" integer NOT NULL,
                                          beatmap_type integer NOT NULL,
                                          beatmap_data json NOT NULL,
                                          user_name text
);


ALTER TABLE public.osu_user_beatmaps OWNER TO digibo;

--
-- Name: osu_user_recent_activity; Type: TABLE; Schema: public; Owner: digibo
--

CREATE TABLE public.osu_user_recent_activity (
                                                 uuid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
                                                 user_id integer NOT NULL,
                                                 last_activity_id integer NOT NULL
);


ALTER TABLE public.osu_user_recent_activity OWNER TO digibo;

--
-- Name: osu_user_watcher; Type: TABLE; Schema: public; Owner: digibo
--

CREATE TABLE public.osu_user_watcher (
                                         user_id integer NOT NULL,
                                         user_name text NOT NULL
);


ALTER TABLE public.osu_user_watcher OWNER TO digibo;

--
-- Name: osu_user_watcher_channel; Type: TABLE; Schema: public; Owner: digibo
--

CREATE TABLE public.osu_user_watcher_channel (
                                                 user_id integer NOT NULL,
                                                 channel_id text NOT NULL
);


ALTER TABLE public.osu_user_watcher_channel OWNER TO digibo;

--
-- Name: random_answer_list; Type: TABLE; Schema: public; Owner: digibo
--

CREATE TABLE public.random_answer_list (
                                           uuid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
                                           command text NOT NULL,
                                           value text NOT NULL,
                                           discord_id text DEFAULT ''::text NOT NULL
);


ALTER TABLE public.random_answer_list OWNER TO digibo;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: digibo
--

CREATE TABLE public.schema_migrations (
                                          version bigint NOT NULL,
                                          dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO digibo;

--
-- Name: calendar calendar_pk; Type: CONSTRAINT; Schema: public; Owner: digibo
--

ALTER TABLE ONLY public.calendar
    ADD CONSTRAINT calendar_pk PRIMARY KEY (uuid);


--
-- Name: event event_pk; Type: CONSTRAINT; Schema: public; Owner: digibo
--

ALTER TABLE ONLY public.event
    ADD CONSTRAINT event_pk PRIMARY KEY (uuid);


--
-- Name: osu_user_beatmaps osu_user_beatmaps_pk; Type: CONSTRAINT; Schema: public; Owner: digibo
--

ALTER TABLE ONLY public.osu_user_beatmaps
    ADD CONSTRAINT osu_user_beatmaps_pk PRIMARY KEY (uuid);


--
-- Name: osu_user_recent_activity osu_user_recent_activity_pk; Type: CONSTRAINT; Schema: public; Owner: digibo
--

ALTER TABLE ONLY public.osu_user_recent_activity
    ADD CONSTRAINT osu_user_recent_activity_pk PRIMARY KEY (uuid);


--
-- Name: osu_user_watcher_channel osu_user_watcher_channel_pk; Type: CONSTRAINT; Schema: public; Owner: digibo
--

ALTER TABLE ONLY public.osu_user_watcher_channel
    ADD CONSTRAINT osu_user_watcher_channel_pk PRIMARY KEY (user_id, channel_id);


--
-- Name: osu_user_watcher osu_user_watcher_pk; Type: CONSTRAINT; Schema: public; Owner: digibo
--

ALTER TABLE ONLY public.osu_user_watcher
    ADD CONSTRAINT osu_user_watcher_pk PRIMARY KEY (user_id);


--
-- Name: random_answer_list random_answer_list_pk; Type: CONSTRAINT; Schema: public; Owner: digibo
--

ALTER TABLE ONLY public.random_answer_list
    ADD CONSTRAINT random_answer_list_pk PRIMARY KEY (uuid);


--
-- Name: random_answer_list random_answer_list_pk_2; Type: CONSTRAINT; Schema: public; Owner: digibo
--

ALTER TABLE ONLY public.random_answer_list
    ADD CONSTRAINT random_answer_list_pk_2 UNIQUE (command, value);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: digibo
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: osu_user_recent_activity_user_id_uindex; Type: INDEX; Schema: public; Owner: digibo
--

CREATE UNIQUE INDEX osu_user_recent_activity_user_id_uindex ON public.osu_user_recent_activity USING btree (user_id);


--
-- Name: event event_event_uuid_fk; Type: FK CONSTRAINT; Schema: public; Owner: digibo
--

ALTER TABLE ONLY public.event
    ADD CONSTRAINT event_event_uuid_fk FOREIGN KEY (parent_event) REFERENCES public.event(uuid) ON DELETE CASCADE;


--
-- Name: osu_user_recent_activity osu_user_recent_activity_osu_user_watcher_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: digibo
--

ALTER TABLE ONLY public.osu_user_recent_activity
    ADD CONSTRAINT osu_user_recent_activity_osu_user_watcher_user_id_fk FOREIGN KEY (user_id) REFERENCES public.osu_user_watcher(user_id);


--
-- Name: osu_user_watcher_channel osu_user_watcher_channel_osu_user_watcher_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: digibo
--

ALTER TABLE ONLY public.osu_user_watcher_channel
    ADD CONSTRAINT osu_user_watcher_channel_osu_user_watcher_user_id_fk FOREIGN KEY (user_id) REFERENCES public.osu_user_watcher(user_id);


--
-- PostgreSQL database dump complete
--

