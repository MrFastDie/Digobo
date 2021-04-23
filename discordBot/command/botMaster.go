package command

func IsBotMaster(id string) bool {
	switch id {
	case "152874021795397633", //MrFastDie
	"463652687288991756": //Shozilina
		return true
	default:
		return false
	}
}
