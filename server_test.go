package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

var PlayerScores = map[string]int{
	"Pepper": 20,
	"Floyd": 10,
}

func TestGETPlayerScore(t *testing.T) {
	testCases := []struct {
		desc string
		player string
		score int
		status int
	} {
		{
			desc: "returns Pepper's score",
			player: "Pepper",
			score: 20,
			status: http.StatusOK,
		},
		{
			desc: "returns Floyd's score",
			player: "Floyd",
			score: 10,
			status: http.StatusOK,
		},
		{
			desc: "doesn't return Apollo's score (unknown)",
			player: "Apollo",
			score: 0,
			status: http.StatusNotFound,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			request := newGetScoreRequest(tt.player)
			response := httptest.NewRecorder()

			server := &PlayerServer{&StubPlayerStore{scores: PlayerScores}}
			server.ServeHTTP(response, request)

			got := response.Body.String()

			if got != strconv.Itoa(tt.score) {
				t.Errorf("got %s, want %d", got, tt.score)
			}

			if response.Code != tt.status {
				t.Errorf("got status code %d, want status code %d", response.Code, tt.status)
			}
		})
	}
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{scores: PlayerScores}
	server := &PlayerServer{&store}

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusAccepted {
			t.Errorf("got status code %d, want status code %d", response.Code, http.StatusAccepted)
		}

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner \ngot %q, want %q", store.winCalls[0], player)
		}
	})
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}
