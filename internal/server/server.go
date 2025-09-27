package server

import (
	"encoding/json"
	"fmt"
	"github.com/osagemo/go-with-tests/internal/game"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store game.PlayerStore
	http.Handler
}

const jsonContentType = "application/json"

func NewPlayerServer(store game.PlayerStore) *PlayerServer {
	server := new(PlayerServer)
	server.store = store

	router := http.NewServeMux()
	router.Handle("/players/", http.HandlerFunc(server.playersHandler))
	router.Handle("/league", http.HandlerFunc(server.leagueHandler))

	server.Handler = router

	return server
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		p.getScore(w, player)
	case http.MethodPost:
		p.storeWin(w, player)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (p *PlayerServer) getScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) storeWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
