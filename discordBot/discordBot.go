package discordBot

import (
	"Digobo/config"
	"Digobo/log"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + config.Config.Discord.Token)
	if err != nil {
		log.Error.Fatal("error creating Discord session,", err)
		return
	}

	// Register the mainLoop func as a callback for MessageCreate events.
	dg.AddHandler(mainLoop)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentsDirectMessageReactions | discordgo.IntentsGuildMessageReactions

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Error.Fatal("error opening connection,", err)
		return
	}

	err = dg.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			&discordgo.Activity{
				Name: "\U0001F970",
				Type: 5,
			},
		},
		Status:     "online",
	})

	if err != nil {
		log.Warning.Println("can't update Bot's status", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Info.Println("Discord bot Digobo is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
