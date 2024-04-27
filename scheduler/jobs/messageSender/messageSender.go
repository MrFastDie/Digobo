package messageSender

import (
	"Digobo/discordBot"
	"Digobo/discordBot/command/commands/waifu"
	"Digobo/log"
	"Digobo/scheduler"
	"encoding/json"
	"time"
)

type Data struct {
	ChannelId string
}

type MessageSender struct{}

func (this *MessageSender) Execute(rawData string) error {
	var jobData Data

	err := json.Unmarshal([]byte(rawData), &jobData)
	if err != nil {
		log.Error.Println("Cant unmarshal data for scheduled job", err)
		return err
	}

	waifu.WaifuSender(discordBot.GetInstance(), jobData.ChannelId)

	job := scheduler.Job{
		ExecutionTime: time.Now().Add(10 * time.Second),
		ExecutionFunc: &MessageSender{},
		Data:          rawData,
	}

	scheduler.GetScheduler().AddThresholdJob(job)

	return nil
}
