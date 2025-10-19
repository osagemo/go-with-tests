package commandline_test

import (
	"strings"
	"testing"

	"github.com/osagemo/go-with-tests/internal/commandline"
	"github.com/osagemo/go-with-tests/internal/game"
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win", func(t *testing.T) {

		in := strings.NewReader("Chris wins\n")
		store := &game.StubPlayerStore{}
		cli, _ := commandline.NewCli(store, in)

		cli.PlayPoker()
		game.AssertPlayerWin(t, store, "Chris")
	})

	t.Run("record Stuart win", func(t *testing.T) {

		in := strings.NewReader("Stuart wins\n")
		store := &game.StubPlayerStore{}
		cli, _ := commandline.NewCli(store, in)

		cli.PlayPoker()
		game.AssertPlayerWin(t, store, "Stuart")
	})
}
