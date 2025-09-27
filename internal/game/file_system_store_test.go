package game

import (
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	tempFile := CreateTempFile(t, `[
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
		AssertLeague(t, want, got)

		got = store.GetLeague()
		AssertLeague(t, want, got)
	})

	t.Run("league is sorted by wins", func(t *testing.T) {
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}
		got := store.GetLeague()
		AssertLeague(t, got, want)
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
