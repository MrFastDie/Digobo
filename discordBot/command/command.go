package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"strings"
)

// TODO implement event system with states
type Command interface {
	// Which is the command that invokes this - no spaces allowed everything else goes into args
	Name() string

	// A brief title of this command
	Title() string

	// A short description for this command
	Description() string

	// Does this command use Events or is it just an ping - pong command?
	HasInteractions() bool

	// The main execution progress of the command
	Execute(string, *discordgo.Session, *discordgo.MessageCreate) error
}

type commandsStruct struct {
	commands map[string]Command
}

func (this *commandsStruct) LoadCommand(command Command) error {
	if this.commands[strings.ToLower(command.Name())] != nil {
		return errors.New(fmt.Sprintf("command %s already loaded", command.Name()))
	}

	this.commands[strings.ToLower(command.Name())] = command

	return nil
}

func (this *commandsStruct) GetCommand(name string) (Command, error) {
	command := this.commands[strings.ToLower(name)]

	if command == nil {
		return nil, errors.New(fmt.Sprintf("command %s not found", name))
	}

	return command, nil
}

func (this *commandsStruct) GetCommands() map[string]Command {
	return this.commands
}

var Commands commandsStruct

func init() {
	Commands = commandsStruct{commands: map[string]Command{}}
}
