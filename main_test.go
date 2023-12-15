package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gofr.dev/pkg/gofr"
)

func TestCreatecar(t *testing.T) {
	// Create a request body with valid JSON data
	requestBody := []byte(`{"license_plate": "ABC123", "model": "Sedan", "color": "Blue", "repair_status": "OK"}`)

	// Create a request with the given request body
	request, err := http.NewRequest("POST", "/car/enter", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	responseRecorder := httptest.NewRecorder()

	// Create a gofr Context for testing
	ctx := gofr.NewTestContext(request)

	// Call the Createcar function, passing the response recorder and context
	Createcar(ctx)

	// Check the HTTP status code
	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("Unexpected status code. Expected: %d, Got: %d", http.StatusOK, responseRecorder.Code)
	}

	// Parse the response body
	var responseBody map[string]interface{}
	err = json.NewDecoder(responseRecorder.Body).Decode(&responseBody)
	if err != nil {
		t.Fatal(err)
	}

	// Add your assertions based on the expected response
	// For example, you might check the content of the response JSON
	expectedMessage := "Car entered the garage successfully"
	if responseMessage, ok := responseBody["message"].(string); !ok || responseMessage != expectedMessage {
		t.Fatalf("Unexpected response. Expected: %s, Got: %s", expectedMessage, responseMessage)
	}

}
