package game

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	tempFile := createTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)

	store, err := NewFileSystemPlayerStore(tempFile)
	if err != nil {
		t.Fatalf("could not create file system store, %v", err)
	}

	t.Run("get player score", func(t *testing.T) {
		got := store.GetPlayerScore("Cleo")
		want := 10
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("get league", func(t *testing.T) {
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}
		got := store.GetLeague()
		assertLeague(t, want, got)

		got = store.GetLeague()
		assertLeague(t, want, got)
	})

	t.Run("league is sorted by wins", func(t *testing.T) {
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}
		got := store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("record player score", func(t *testing.T) {
		store.RecordWin("Cleo")
		got := store.GetPlayerScore("Cleo")
		want := 11
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("record player score for new player", func(t *testing.T) {
		store.RecordWin("Bogdan")
		got := store.GetPlayerScore("Bogdan")
		want := 1
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

// createTempFile creates a temp file with the given data and registers cleanup with t.Cleanup
func createTempFile(t *testing.T, data string) *os.File {
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
