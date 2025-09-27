package server

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/osagemo/go-with-tests/internal/game"
)

func TestGETPlayers(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"king": 20,
			"sven": 10,
		},
	}
	server := NewPlayerServer(store)

	t.Run("returns score for king", func(t *testing.T) {
		request := newGetScoreRequest("king")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseStatus(t, response, http.StatusOK)
		assertResponseBody(t, response, "20")
	})

	t.Run("returns score for Sven", func(t *testing.T) {
		request := newGetScoreRequest("sven")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response, "10")
		assertResponseStatus(t, response, http.StatusOK)
	})

	t.Run("returns 404 for missing player", func(t *testing.T) {
		request := newGetScoreRequest("Ann")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseStatus(t, response, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseStatus(t, response, http.StatusAccepted)
		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}
		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"king": 20,
			"sven": 10,
		},
	}
	server := NewPlayerServer(store)

	t.Run("it returns 200", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/league", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseStatus(t, response, http.StatusOK)
	})

	t.Run("it returns expected JSON", func(t *testing.T) {
		request := newGetLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseContentType(t, response, "application/json")
		assertResponseStatus(t, response, http.StatusOK)

		want := []game.Player{
			{"king", 20},
			{"sven", 10},
		}
		got := getLeagueFromResponse(t, response.Body)
		game.AssertLeague(t, got, want)
	})
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest("GET", fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newGetLeagueRequest() *http.Request {
	request, _ := http.NewRequest("GET", "/league", nil)
	return request
}

func getLeagueFromResponse(t *testing.T, body *bytes.Buffer) []game.Player {
	got, err := game.LeagueFromJson(body)
	if err != nil {
		t.Fatalf("Unable to decode league request body: %v", err)
	}

	return got
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest("POST", fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	got := response.Body.String()
	if got != want {
		t.Fatalf("Expected player score to be %q, got %q", want, got)
	}
}

func assertResponseStatus(t testing.TB, response *httptest.ResponseRecorder, want int) {
	t.Helper()

	got := response.Code
	if got != want {
		t.Fatalf("Expected status code to be %v, got %v", want, got)
	}
}

func assertResponseContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	got := response.Result().Header.Get("content-type")
	if got != want {
		t.Fatalf("Expected content-type %v, got %v", want, got)
	}
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() game.League {
	league := []game.Player{}
	for name, wins := range s.scores {
		league = append(league, game.Player{Name: name, Wins: wins})
	}
	return league
}
