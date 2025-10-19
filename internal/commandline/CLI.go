package commandline

import (
	"bufio"
	"io"
	"strings"

	"github.com/osagemo/go-with-tests/internal/game"
)

type CLI struct {
	store   game.PlayerStore
	scanner *bufio.Scanner
}

func NewCli(store game.PlayerStore, in io.Reader) (*CLI, error) {
	return &CLI{store, bufio.NewScanner(in)}, nil
}

func (c *CLI) PlayPoker() {
	if err := c.scanner.Err(); err != nil {
		panic("omg scan error")
	}

	if c.scanner.Scan() {
		line := c.scanner.Text()
		winner := strings.Split(line, " ")
		c.store.RecordWin(winner[0])
	}
}
