package command

var mapC = map[string]Command{}

func GetMap() map[string]Command {
	return mapC
}

func AddCommand(command Command) {
	mapC[command.Name] = command
}
