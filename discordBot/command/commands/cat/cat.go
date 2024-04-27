package cat

import (
	"encoding/json"
	"log" // Import the log package
	"net/http"
	"time"

	"Digobo/config"
	"Digobo/discordBot/command"
	"github.com/bwmarrin/discordgo"
)

var Command = command.Command{
	Name:        "cat",
	Description: "Get a random cat picture",
	Execute:     executeCatCommand,
}

func executeCatCommand(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	// Send a random cat picture
	if err := sendRandomCat(s, i.ChannelID); err != nil {
		log.Println("Error sending cat picture:", err)
	}
	return nil
}

func sendRandomCat(s *discordgo.Session, channelID string) error {
	// Send a request to The Cat API to fetch a random cat image
	response, err := http.Get("https://api.thecatapi.com/v1/images/search")
	if err != nil {
		log.Println("Error fetching cat image:", err)
		return err

	}
	defer func() {
		// Check for error when closing the response body
		if err := response.Body.Close(); err != nil {
			log.Println("Error closing response body:", err)
		}
	}()

	// Decode the response JSON
	var data []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil || len(data) == 0 {
		log.Println("Error decoding cat image response:", err)
		return err
	}

	// Extract the image URL from the response
	imageURL := data[0]["url"].(string)

	// Create an embed with the cat image
	embed := &discordgo.MessageEmbed{
		Title:     "Random Cat Picture",
		Image:     &discordgo.MessageEmbedImage{URL: imageURL},
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     config.Config.Bot.DefaultEmbedColor,
	}

	// Send the embed as a message in the channel
	_, err = s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		log.Println("Error sending cat image:", err)
		return err
	}

	return nil
}

func init() {
	command.AddCommand(Command)
}
