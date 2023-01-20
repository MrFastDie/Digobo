package discordBot

import "github.com/bwmarrin/discordgo"

func ComesFromDM(s *discordgo.Session, i *discordgo.InteractionCreate) (bool, error) {
	channel, err := s.State.Channel(i.ChannelID)
	if err != nil {
		if channel, err = s.Channel(i.ChannelID); err != nil {
			return false, err
		}
	}

	return channel.Type == discordgo.ChannelTypeDM, nil
}
