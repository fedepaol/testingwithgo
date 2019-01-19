package httptest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// START OMIT
//we want to test
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

//our test
func TestHealthCheckHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/health-check", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(rr, req)

	if rr.Body.String() != `{"alive": true}` {
		t.Errorf("Unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// END OMIT
