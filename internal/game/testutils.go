package game

import (
	"os"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.Scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	league := []Player{}
	for name, wins := range s.Scores {
		league = append(league, Player{Name: name, Wins: wins})
	}
	return league
}

// CreateTempFile creates a temp file with the given data and registers cleanup with t.Cleanup
func CreateTempFile(t *testing.T, data string) *os.File {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}

	if _, err := tmpFile.Write([]byte(data)); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		t.Fatalf("could not write to temp file: %v", err)
	}

	t.Cleanup(func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	})

	return tmpFile
}

func AssertLeague(t *testing.T, got, want []Player) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}
	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.WinCalls[0], winner)
	}
}
