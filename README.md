# Digobo
A Discord bot written in go

All commands get added as discord `/` commands

The bot needs a postgreSQL database

### Features
 - Schedule messages like a reminder in a calender (for this rrules are used)
 - Get a notification if a streamer comes online in twitch
 - Get a notification if a user in osu! gets a new highscore, achievement or uploads a new track
 
### Requirements
You need to create a postgres database with the `uuid_ossp` extension enabled

    CREATE EXTENSION "uuid_ossp";

You can the compiled binary with the argument `migrate` to get the DB filled after creating a `config.yaml`

To start the bot simply run the binary with the argument `start`