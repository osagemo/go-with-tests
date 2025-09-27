package game

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type Player struct {
	Name string `json:"Name"`
	Wins int    `json:"Wins"`
}
