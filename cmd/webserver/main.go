package main

import (
	"github.com/osagemo/go-with-tests/internal/game"
	"github.com/osagemo/go-with-tests/internal/server"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := game.FileSystemPlayerStoreFromFile(dbFileName)
	defer close()

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := server.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
