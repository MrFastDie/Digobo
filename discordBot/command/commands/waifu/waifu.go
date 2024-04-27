package waifu

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"Digobo/config"
	"Digobo/discordBot/command"
	"github.com/bwmarrin/discordgo"
)

var Command = command.Command{
	Name:        "waifu",
	Description: "Get a random waifu picture",
	Execute:     executeWaifuCommand,
}

func executeWaifuCommand(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	// Send a random waifu picture
	if err := sendRandomWaifu(s, i.ChannelID); err != nil {
		log.Println("Error sending waifu picture:", err)
	}
	return nil
}

func sendRandomWaifu(s *discordgo.Session, channelID string) error {
	// Send a request to the Waifu.im API to fetch a random waifu image
	response, err := http.Get("https://api.waifu.pics/sfw/neko")
	if err != nil {
		log.Println("Error fetching waifu image:", err)
		return err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return err
	}

	log.Println("Response body:", string(body)) // Print the response body for debugging

	// Decode the response JSON
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error decoding waifu image response:", err)
		return err
	}

	// Extract the image URL from the response
	imageURL := data["url"].(string)

	// Create an embed with the waifu image
	embed := &discordgo.MessageEmbed{
		Title:     "Random Waifu Picture",
		Image:     &discordgo.MessageEmbedImage{URL: imageURL},
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     config.Config.Bot.DefaultEmbedColor,
	}

	// Send the embed as a message in the channel
	_, err = s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		log.Println("Error sending waifu image:", err)
		return err
	}

	return nil
}

func WaifuSender(s *discordgo.Session, channelID string) error {
	return sendRandomWaifu(s, channelID)
}

func init() {
	command.AddCommand(Command)
}
