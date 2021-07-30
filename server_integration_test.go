package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	request := newPostWinRequest(player)

	server.ServeHTTP(httptest.NewRecorder(), request)
	server.ServeHTTP(httptest.NewRecorder(), request)
	server.ServeHTTP(httptest.NewRecorder(), request)

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	if response.Code != http.StatusOK {
		t.Errorf("got status %d want status %d", response.Code, http.StatusOK)
	}

	rBody := response.Body.String()
	if rBody != "3" {
		t.Errorf("response body did not have the right score. Got '%s' want '%s'", rBody, "3")
	}
}
