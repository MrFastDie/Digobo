package discordBot

import (
	"Digobo/config"
	"Digobo/discordBot/command"
	"Digobo/log"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/exp/maps"
	"os"
	"os/signal"
	"syscall"
)

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

func createCommands(command []command.Command) []*discordgo.ApplicationCommand {
	var commands []*discordgo.ApplicationCommand

	for _, v := range command {
		var cmd = &discordgo.ApplicationCommand{
			Name:        v.Name,
			Description: v.Description,
		}

		if "" != config.Config.Bot.CommandPrefix {
			cmd.Name = fmt.Sprintf("%s-%s", config.Config.Bot.CommandPrefix, v.Name)
		}

		if len(v.SubCommands) > 0 {
			cmd.Options = createOptions(maps.Values(v.SubCommands))
		}

		commands = append(commands, cmd)
	}

	return commands
}

func createOptions(command []command.SubCommand) []*discordgo.ApplicationCommandOption {
	var options []*discordgo.ApplicationCommandOption

	for _, v := range command {
		var cmd = &discordgo.ApplicationCommandOption{
			Name:        v.Name,
			Description: v.Description,
			Type:        v.Type,
		}

		if len(v.SubCommands) > 0 {
			cmd.Type = discordgo.ApplicationCommandOptionSubCommand
			cmd.Options = createOptions(maps.Values(v.SubCommands))
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

	// TODO get a version for a command, if that command changes, update it here - if not dont
	// TODO make guild specific commands
	commands = createCommands(maps.Values(command.GetMap()))

	oldCommands, err := instance.ApplicationCommands(instance.State.User.ID, "")
	if err != nil {
		log.Error.Println("unable to fetch all application commands")
	}

	desiredMap := make(map[string]*discordgo.ApplicationCommand)
	for _, cmd := range commands {
		desiredMap[cmd.Name] = cmd
	}

	existingMap := make(map[string]*discordgo.ApplicationCommand)
	for _, cmd := range oldCommands {
		existingMap[cmd.Name] = cmd
	}

	// Delete commands not in the desired list
	for _, cmd := range oldCommands {
		if _, found := desiredMap[cmd.Name]; !found {
			err := instance.ApplicationCommandDelete(instance.State.User.ID, "", cmd.ID)
			if err != nil {
				log.Error.Printf("Failed to delete command %s (%s) in guild %s: %v\n", cmd.Name, cmd.ID, "", err)
			} else {
				log.Error.Printf("Successfully deleted command %s (%s) in guild %s\n", cmd.Name, cmd.ID, "")
			}
		}
	}

	// Create or update existing commands
	for _, cmd := range commands {
		if existingCmd, found := existingMap[cmd.Name]; found {
			// Edit existing command
			_, err := instance.ApplicationCommandEdit(instance.State.User.ID, "", existingCmd.ID, cmd)
			if err != nil {
				log.Error.Printf("Failed to edit command %s (%s) in guild %s: %v\n", cmd.Name, cmd.ID, "", err)
			} else {
				log.Error.Printf("Successfully edited command %s (%s) in guild %s\n", cmd.Name, cmd.ID, "")
			}
		} else {
			// Create new command
			_, err := instance.ApplicationCommandCreate(instance.State.User.ID, "", cmd)
			if err != nil {
				log.Error.Printf("Failed to create command %s in guild %s: %v\n", cmd.Name, "", err)
			} else {
				log.Error.Printf("Successfully created command %s in guild %s\n", cmd.Name, "")
			}
		}
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
