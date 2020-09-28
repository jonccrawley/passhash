package handler

import (
	"github.com/jonccrawley/passhash/definition"
	"github.com/jonccrawley/passhash/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatisticsHandler(t *testing.T) {

	definition.StatisticsRepo = &repository.StatisticsRepository{}

	req, err := http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatisticsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestStatisticsHandlerBadRequestMethod(t *testing.T) {

	definition.StatisticsRepo = &repository.StatisticsRepository{}

	req, err := http.NewRequest("POST", "/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatisticsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestStatisticsHandlerResultSingle(t *testing.T) {

	definition.StatisticsRepo = &repository.StatisticsRepository{}
	definition.StatisticsRepo.Add(100)
	definition.StatisticsRepo.Add(50)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatisticsHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status !=  http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status,  http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"total":2,"average":75}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


