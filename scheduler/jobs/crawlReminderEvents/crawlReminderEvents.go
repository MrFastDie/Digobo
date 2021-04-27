package crawlReminderEvents

import (
	"Digobo/calendar"
	"Digobo/database"
	"Digobo/json"
	"Digobo/log"
	"Digobo/scheduler"
	"Digobo/scheduler/jobs/notifyEventParticipants"
	"time"
)

type Data struct {
	NotifyChannels []string `json:"notify_channels"`
	NotifyUsers    []string `json:"notify_users"`
}

type CrawlReminderEvents struct{}

func (this *CrawlReminderEvents) Execute(data string) error {
	events, err := database.GetNextTwoHourEventsByType(calendar.CalendarTypeReminder)
	if err != nil {
		return err
	}

	if data != "" {
		var prevRunEvents []database.Event
		err = json.Default.Unmarshal([]byte(data), &prevRunEvents)
		if err != nil {
			log.Error.Fatal("cant unmarshal CrawlReminderEvents []byte into []database.Event", err)
			return err
		}

		// compare arrays and delete occurences from array prevRunEvents out of the actual events array
		for i := 0; i < len(events); {
			exists := false
			for _, b := range prevRunEvents {
				if b == events[i] {
					exists = true
					break
				}
			}
			if !exists {
				events = append(events[:i], events[i+1:]...)
			} else {
				i++
			}
		}
	}

	for _, notifyEvent := range events {
		var eventData Data
		err = json.Default.Unmarshal([]byte(notifyEvent.Data), &eventData)
		if err != nil {
			log.Warning.Println("can't unmarshal into notifyEvent.Data", err)
			continue
		}

		var notify notifyEventParticipants.Data
		notify.ParentEventUuid = notifyEvent.ParentEvent
		notify.NotifyChannels = eventData.NotifyChannels
		notify.NotifyUsers = eventData.NotifyUsers
		notify.OriginDate = notifyEvent.Occurrences

		notifyByteData, err := json.Default.Marshal(notify)
		if err != nil {
			log.Warning.Println("can't marshal notifyEvent.data to []byte", err)
			continue
		}

		job := scheduler.Job{
			ExecutionTime: notifyEvent.Occurrences.Time,
			ExecutionFunc: &notifyEventParticipants.NotifyEventParticipants{},
			Data:          string(notifyByteData),
		}

		scheduler.GetScheduler().AddThresholdJob(job)
	}

	// when this job is done start next job in 1h30
	dataStr, err := json.Default.Marshal(events)
	if err != nil {
		log.Error.Fatal("cant marshal CrawlReminderEvents events into []byte", err)
		return err
	}

	CrawlReminderEventJobStart(time.Now().Add(time.Hour*1+time.Minute*30), string(dataStr))

	return nil
}

func CrawlReminderEventJobStart(runTime time.Time, data string) {
	job := scheduler.Job{
		ExecutionTime: runTime,
		ExecutionFunc: &CrawlReminderEvents{},
		Data:          data,
	}

	scheduler.GetScheduler().AddThresholdJob(job)
}
