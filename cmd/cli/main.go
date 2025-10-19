package main

import (
	"fmt"
	"log"
	"os"

	"github.com/osagemo/go-with-tests/internal/commandline"
	"github.com/osagemo/go-with-tests/internal/game"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := game.FileSystemPlayerStoreFromFile(dbFileName)
	defer close()

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	cli, err := commandline.NewCli(store, os.Stdin)
	if err != nil {
		log.Fatalf("problem creating cli instance, %v ", err)
	}

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	cli.PlayPoker()
}
