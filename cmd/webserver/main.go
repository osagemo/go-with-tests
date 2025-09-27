package main

import (
	"github.com/osagemo/go-with-tests/internal/game"
	"github.com/osagemo/go-with-tests/internal/server"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := game.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := server.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
