package game

import (
	"os"
	"reflect"
	"testing"
)

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
