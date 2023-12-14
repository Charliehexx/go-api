package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/request"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {
	go main()
	time.Sleep(3 * time.Second)

	tests := []struct {
		desc       string
		method     string
		endpoint   string
		statusCode int
		body       []byte
	}{
		{"get employees", http.MethodGet, "car", http.StatusOK, nil},
		{"post employees", http.MethodPost, "car/enter", http.StatusCreated, []byte(`{
			"id":1,
			"model":"mahak",
			"color":"msjce",
			"repair_status":"928902"
		}`),
		}}

	for i, tc := range tests {
		req, _ := request.NewMock(tc.method, "http://localhost:9000/"+tc.endpoint, bytes.NewBuffer(tc.body))

		c := http.Client{}

		resp, err := c.Do(req)
		if err != nil {
			t.Errorf("TEST[%v] Failed.\tHTTP request encountered Err: %v\n%s", i, err, tc.desc)
			continue
		}

		if resp.StatusCode != tc.statusCode {
			t.Errorf("TEST[%v] Failed.\tExpected %v\tGot %v\n%s", i, tc.statusCode, resp.StatusCode, tc.desc)
		}

		_ = resp.Body.Close()
	}
}
func TestCreateCar(t *testing.T) {
	// Create a new instance of your application
	go
	app := gofr.New()

	// Create a test server
	server := httptest.NewServer(app.Handler())
	defer server.Close()

	// Prepare a JSON payload for the POST request
	payload := `{"license_plate": "ABC123", "model": "Sedan", "color": "Blue", "repair_status": "Pending"}`

	// Make a POST request to the /car/enter endpoint
	resp, err := http.Post(server.URL+"/car/enter", "application/json", bytes.NewBufferString(payload))
	if err != nil {
		t.Fatalf("Error making POST request: %v", err)
	}
	defer resp.Body.Close()

	// Verify the response status code is as expected
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code %d, got %d", http.StatusOK, resp.StatusCode)

	// Decode the response body
	var response interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	// Verify the response content, adjust based on your application logic
	expectedResponse := "Car entered the garage successfully"
	assert.Equal(t, expectedResponse, response, "Unexpected response content")

	// You can add more assertions based on your application logic
}
