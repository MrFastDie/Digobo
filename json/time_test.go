package json

import (
	"bytes"
	"testing"
	"time"
)

func TestTimestampWithoutTimezoneJSON(t *testing.T) {
	t.Run("TestMarshalJSON", func(t *testing.T) {
		// Create a timestamp without timezone
		timestamp := time.Date(2024, time.April, 26, 15, 30, 0, 0, time.UTC)
		twt := TimestampWithoutTimezone{timestamp}

		// Marshal the timestamp to JSON
		data, err := twt.MarshalJSON()
		if err != nil {
			t.Fatalf("error marshalling timestamp: %v", err)
		}

		// Expected JSON string
		expected := []byte(`"2024-04-26T15:30:00"`)

		// Compare the marshalled JSON with the expected JSON
		if !bytes.Equal(data, expected) {
			t.Errorf("unexpected JSON output: got %s, want %s", data, expected)
		}
	})

	t.Run("TestUnmarshalJSON", func(t *testing.T) {
		// Input JSON data
		inputJSON := []byte(`"2024-04-26T15:30:00"`)

		// Create a timestamp without timezone
		twt := &TimestampWithoutTimezone{}

		// Unmarshal the JSON data
		err := twt.UnmarshalJSON(inputJSON)
		if err != nil {
			t.Fatalf("error unmarshalling JSON: %v", err)
		}

		// Expected timestamp
		expected := time.Date(2024, time.April, 26, 15, 30, 0, 0, time.UTC)

		// Compare the unmarshalled timestamp with the expected timestamp
		if !twt.Equal(expected) {
			t.Errorf("unexpected timestamp: got %v, want %v", twt, expected)
		}
	})
}
