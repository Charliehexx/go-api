package main

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"gofr.dev/pkg/gofr"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
func TestGetcars(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Set up expected rows and columns
	expectedRows := sqlmock.NewRows([]string{"id", "license_plate", "color", "model", "repair_status", "entry_time"}).
		AddRow(1, "ABC123", "Blue", "Sedan", "OK", time.Now()).
		AddRow(2, "XYZ789", "Red", "SUV", "Needs Repair", time.Now())

	// Expect the SQL query and return the mock rows
	mock.ExpectQuery("SELECT * FROM cars").WillReturnRows(expectedRows)

	// Create a gofr context with the mock database
	ctx := gofr.NewTestContextWithDB(db)

	// Call the Getcars function
	result, err := Getcars(ctx)

	// Check for errors
	if err != nil {
		t.Fatalf("Error during Getcars: %v", err)
	}

	// Assert the result type
	cars, ok := result.([]Car)
	if !ok {
		t.Fatalf("Unexpected result type. Expected: []Car, Got: %T", result)
	}

	// Assert the length of the cars slice
	if len(cars) != 2 {
		t.Fatalf("Unexpected number of cars. Expected: 2, Got: %d", len(cars))
	}

	// Add more specific assertions based on your requirements
	// For example, check the values of individual cars in the slice

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations: %s", err)
	}
}
