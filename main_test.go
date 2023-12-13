// main_test.go
package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr/request"
	"net/http"
	"net/http/httptest"
	"testing"
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
		{"get employees", http.MethodGet, "employee", http.StatusOK, nil},
		{"post employees", http.MethodPost, "employee", http.StatusCreated, []byte(`{
			"id":80,
			"name":"mahak",
			"email":"msjce",
			"phone":928902,
			"city":"kolkata"
		}`),
		}}

	for i, tc := range tests {
		req, _ := request.NewMock(tc.method, "http://localhost:3000/"+tc.endpoint, bytes.NewBuffer(tc.body))

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
