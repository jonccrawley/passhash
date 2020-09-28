package handler

import (
	"github.com/jonccrawley/passhash/definition"
	"github.com/jonccrawley/passhash/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSaveHashHandler(t *testing.T) {

	definition.HashRepo = &repository.HashRepository{}
	definition.StatisticsRepo = &repository.StatisticsRepository{}
	definition.ExecutionRepo = &repository.ExecutionRepository{}


	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/hash", strings.NewReader("password=123"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SaveHashHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body is what we expect.
	expected := `1`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSaveHashHandlerBadRequestMethod(t *testing.T) {

	definition.HashRepo = &repository.HashRepository{}
	definition.StatisticsRepo = &repository.StatisticsRepository{}
	definition.ExecutionRepo = &repository.ExecutionRepository{}


	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/hash", nil)

	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SaveHashHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestSaveHashHandlerEmptyPassword(t *testing.T) {

	definition.HashRepo = &repository.HashRepository{}
	definition.StatisticsRepo = &repository.StatisticsRepository{}
	definition.ExecutionRepo = &repository.ExecutionRepository{}


	req, err := http.NewRequest("POST", "/hash", strings.NewReader("password="))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SaveHashHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestSaveHashHandlerPasswordToLong(t *testing.T) {

	definition.HashRepo = &repository.HashRepository{}
	definition.StatisticsRepo = &repository.StatisticsRepository{}
	definition.ExecutionRepo = &repository.ExecutionRepository{}


	req, err := http.NewRequest("POST", "/hash", strings.NewReader("password=1111111111111111111111111111111"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SaveHashHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}