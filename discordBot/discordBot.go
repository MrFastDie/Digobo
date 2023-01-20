package discordBot

import (
	"Digobo/config"
	"Digobo/discordBot/command"
	"Digobo/log"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var Whitespace = "⠀"

var instance *discordgo.Session

var commands []*discordgo.ApplicationCommand

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

func createCommands(command []*cobra.Command) []*discordgo.ApplicationCommand {
	var commands []*discordgo.ApplicationCommand

	for _, v := range command {
		var cmd = &discordgo.ApplicationCommand{
			Name:        v.Use,
			Description: v.Short,
		}

		if len(v.Commands()) > 0 {
			cmd.Options = createOptions(v.Commands())
		}

		commands = append(commands, cmd)
	}

	return commands
}

func createOptions(command []*cobra.Command) []*discordgo.ApplicationCommandOption {
	var options []*discordgo.ApplicationCommandOption

	for _, v := range command {
		var cmd = &discordgo.ApplicationCommandOption{
			Name:        v.Use,
			Description: v.Short,
			Type:        discordgo.ApplicationCommandOptionString,
		}

		if len(v.Commands()) > 0 {
			cmd.Type = discordgo.ApplicationCommandOptionSubCommand
			cmd.Options = createOptions(v.Commands())
		}

		options = append(options, cmd)
	}

	return options
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

	commands = createCommands(command.RootCommand.Commands())

	for i, v := range commands {
		if v == nil {
			continue
		}

		_, err := instance.ApplicationCommandCreate(instance.State.User.ID, "", commands[i])
		if err != nil {
			log.Error.Fatal(fmt.Sprintf("unable to add command %s", v.Name), err)
			return
		}
	}

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

func SendMessageFromInteraction(msg string, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	_, err := s.ChannelMessageSend(i.ChannelID, msg)
	if err != nil {
		log.Error.Println("can't send channel message", err)

		return err
	}

	return nil
}

func SendInteractionMessage(msg string, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
}

func SendEmbed(embed *discordgo.MessageEmbed, channelId string, s *discordgo.Session) error {
	_, err := s.ChannelMessageSendEmbed(channelId, embed)
	if err != nil {
		log.Error.Println("can't send embed", err)

		return err
	}

	return nil
}
