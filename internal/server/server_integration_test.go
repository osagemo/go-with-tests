package server

import (
	"github.com/osagemo/go-with-tests/internal/game"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	file := game.CreateTempFile(t, "")

	store, err := game.NewFileSystemPlayerStore(file)
	if err != nil {
		t.Fatalf("could not create file system store, %v", err)
	}

	server := NewPlayerServer(store)
	player := "ostron"

	// Send 3 POST requests to /players/ostron
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertResponseStatus(t, response, http.StatusOK)
		assertResponseBody(t, response, "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetLeagueRequest())
		assertResponseStatus(t, response, http.StatusOK)

		want := []game.Player{
			{"ostron", 3},
		}
		got := getLeagueFromResponse(t, response.Body)
		game.AssertLeague(t, got, want)
	})
}
