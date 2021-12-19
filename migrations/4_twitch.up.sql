CREATE TABLE public.twitch_watcher (
                                                 uuid uuid DEFAULT public.uuid_generate_v4() NOT NULL,
                                                 user_id TEXT NOT NULL,
                                                 channel_id TEXT NOT NULL
);

ALTER TABLE twitch_watcher ADD COLUMN online BOOLEAN NOT NULL DEFAULT false;