package json

import (
	"bytes"
	"strings"
	"testing"
)

func TestJSON(t *testing.T) {
	// Define test data
	var byteStr = []byte(`{"example_example":"Hello World"}`)
	var testCase = struct {
		ExampleExample string `db:"example_example"`
	}{
		ExampleExample: "Hello World",
	}

	// Test DB configuration
	t.Run("TestMarshalDB", func(t *testing.T) {
		marshalled, err := DB.Marshal(&testCase)
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(marshalled, byteStr) {
			t.Errorf("Marshalled bytes are not equal")
		}
	})

	t.Run("TestUnmarshalDB", func(t *testing.T) {
		var unmarshalled = struct {
			ExampleExample string `db:"example_example"`
		}{}

		err := DB.Unmarshal(byteStr, &unmarshalled)
		if err != nil {
			t.Error(err)
		}

		if unmarshalled.ExampleExample != "Hello World" {
			t.Errorf("Unmarshalled bytes are not equal")
		}
	})

	// Test Unmarshal function
	t.Run("TestUnmarshal", func(t *testing.T) {
		invalidSrc := "invalid data" // Invalid data type (not []byte)
		var out interface{}

		err := Unmarshal(invalidSrc, &out)
		if err == nil {
			t.Error("Expected error for invalid type, got nil")
		}

		expectedErrSubstring := "invalid type"
		if !strings.Contains(err.Error(), expectedErrSubstring) {
			t.Errorf("Expected error message to contain substring '%s', got '%s'", expectedErrSubstring, err.Error())
		}
	})

	// Test Unmarshal function with valid input
	t.Run("TestUnmarshalValid", func(t *testing.T) {
		// Valid JSON data
		validSrc := []byte(`{"example_example":"Valid JSON"}`)
		var out struct {
			ExampleExample string `db:"example_example"`
		}

		err := Unmarshal(validSrc, &out)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if out.ExampleExample != "Valid JSON" {
			t.Errorf("Unexpected unmarshalled data: %v", out.ExampleExample)
		}
	})
}
