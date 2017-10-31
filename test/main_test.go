package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	//Create a request to pass into the handler. We don't have any query parameters.
	req, err := http.NewRequest("GET", "/members", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	//ask for attep request with response writer and request
	handler.ServeHTTP(rr, req)

	//check if status ok
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler did not work got  %v wanted %v", status, http.StatusOK)
	}

	//check response value
	expected := `{"alive":true}`
	if rr.Body.String() != expected {
		t.Errorf("handler got unknown %v wanted %v", rr.Body.String(), expected)
	}
}
