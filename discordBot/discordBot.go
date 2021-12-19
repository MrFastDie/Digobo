package discordBot

import (
	"Digobo/config"
	"Digobo/log"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var Whitespace = "⠀"

var instance *discordgo.Session

func GetInstance() *discordgo.Session {
	var err error

	if instance == nil {
		instance, err = discordgo.New("Bot " + config.Config.Discord.Token)
		if err != nil {
			log.Error.Fatal("error creating Discord session,", err)
		}
	}

	return instance
}

func Close() {
	err := instance.Close()
	if err != nil {
		log.Error.Fatal("error closing Discord session,", err)
		return
	}

	instance = nil
}

func Run() {
	instance = GetInstance()

	// Register the mainLoop func as a callback for MessageCreate events.
	instance.AddHandler(mainLoop)

	// In this example, we only care about receiving message events.
	instance.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentsDirectMessageReactions | discordgo.IntentsGuildMessageReactions

	// Open a websocket connection to Discord and begin listening.
	err := instance.Open()
	if err != nil {
		log.Error.Fatal("error opening connection,", err)
		return
	}

	err = instance.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			&discordgo.Activity{
				Name: "\U0001F970",
				Type: 5,
			},
		},
		Status: "online",
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
	Close()
}

func SendMessage(msg string, channelId string, s *discordgo.Session) error {
	_, err := s.ChannelMessageSend(channelId, msg)
	if err != nil {
		log.Error.Println("can't send channel message", err)

		return err
	}

	return nil
}

func SendEmbed(embed *discordgo.MessageEmbed, channelId string, s *discordgo.Session) error {
	_, err := s.ChannelMessageSendEmbed(channelId, embed)
	if err != nil {
		log.Error.Println("can't send embed", err)

		return err
	}

	return nil
}
