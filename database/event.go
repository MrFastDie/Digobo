package database

import (
	"Digobo/json"
	"Digobo/log"
)

type Event struct {
	Uuid             string                        `db:"uuid" json:"uuid"`
	CalendarUuid     string                        `db:"calendar_uuid" json:"calendar_uuid"`
	RRule            string                        `db:"rrule" json:"rrule"`
	Title            string                        `db:"title" json:"title"`
	Description      string                        `db:"description" json:"description"`
	Type             int                           `db:"type" json:"type"`
	Data             string                        `db:"data" json:"data"`
	CreatorDiscordId string                        `db:"creator_discord_id" json:"creator_discord_id"`
	ParentEvent      string                        `db:"parent_event" json:"parent_event"`
	StartDate        json.TimestampWithoutTimezone `db:"start_date" json:"start_date"`
	Occurrences      json.TimestampWithoutTimezone `db:"occurrences" json:"occurrences"`
}

func (this *Event) Scan(src interface{}) error {
	return json.Unmarshal(src, this)
}

func GetEventByUuid(eventUuid string) (Event, error) {
	var ret Event
	err := db.Get(&ret, `SELECT row_to_json(t) FROM (SELECT uuid, calendar_uuid, rrule, title, description, type, data::TEXT, creator_discord_id, parent_event, start_date::TIMESTAMP FROM event WHERE uuid = $1) t;`, eventUuid)
	if err != nil {
		log.Warning.Printf("cant fetch event by uuid %s %v\n", eventUuid, err)
		return Event{}, err
	}

	return ret, nil
}

func GetNextTwoHourEventsByType(eventType int) ([]Event, error) {
	var ret []Event

	err := db.Select(&ret, `SELECT row_to_json(t) as jsontext FROM (SELECT uuid, calendar_uuid, rrule, title, description, type, data::TEXT, creator_discord_id, parent_event, start_date::TIMESTAMP,
       (_rrule.occurrences(
               rrule,
               start_date::TIMESTAMP,
               tsrange(now()::timestamp, now()::timestamp + (interval '2h'))
           ))
FROM event
WHERE type = $1) t`, eventType)
	if err != nil {
		log.Warning.Println("cant fetch next two hour events by type", err)
		return nil, err
	}

	return ret, nil
}
