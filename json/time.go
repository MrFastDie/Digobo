package json

import (
	"time"
)

type TimestampWithoutTimezone struct {
	time.Time
}

//type TimestampWithoutTimezone time.Time

func (t *TimestampWithoutTimezone) UnmarshalJSON(data []byte) error {
	ts, err := time.Parse("\"2006-01-02T15:04:05\"", string(data))
	if err == nil {
		t.Time = ts.UTC()
	}
	return err
}