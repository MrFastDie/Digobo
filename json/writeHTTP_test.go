package json

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestJson(t *testing.T) {
	var byteStr = []byte(`{"example_example":"Hello World"}`)
	var testCase = struct {
		ExampleExample string `json:"example_example"`
	}{
		ExampleExample: "Hello World",
	}

	t.Run("Test Marshal", func(t *testing.T) {
		marshalled, err := Json.Marshal(&testCase)
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(marshalled, byteStr) {
			t.Errorf("Marshalled bytes are not equal")
		}
	})

	t.Run("Test Unmarshal", func(t *testing.T) {
		var unmarshalled = struct {
			ExampleExample string `json:"example_example"`
		}{}

		err := Json.Unmarshal(byteStr, &unmarshalled)
		if err != nil {
			t.Error(err)
		}

		if unmarshalled.ExampleExample != "Hello World" {
			t.Errorf("Unmarshalled bytes are not equal")
		}
	})

	t.Run("Test WriteHTTP", func(t *testing.T) {
		type TestData struct {
			Name  string
			Value int
		}

		// Test data
		testData := TestData{"example", 42}

		// Create a new HTTP request recorder
		w := httptest.NewRecorder()

		// Call the WriteHTTP function
		err := WriteHTTP(w, testData)
		if err != nil {
			t.Errorf("WriteHTTP returned an error: %v", err)
		}

		// Check the HTTP response status code
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		// Decode the response body
		var result TestData
		err = jsoniter.Unmarshal(w.Body.Bytes(), &result)
		if err != nil {
			t.Errorf("Failed to decode response body: %v", err)
		}

		// Compare the decoded result with the original test data
		if !reflect.DeepEqual(testData, result) {
			t.Errorf("Expected response data %v, got %v", testData, result)
		}
	})
}
