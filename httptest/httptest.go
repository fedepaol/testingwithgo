package httptest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// START OMIT

func TestHealthCheckHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/health-check", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if rr.Code != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if rr.Body.String() != `{"alive": true}` {
		t.Errorf("Unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// END OMIT
