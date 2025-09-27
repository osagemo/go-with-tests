package game

type InMemoryPlayerStore struct {
	score map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (s *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return s.score[player]
}

func (s *InMemoryPlayerStore) RecordWin(player string) {
	s.score[player]++
}

func (s *InMemoryPlayerStore) GetLeague() League {
	league := []Player{}
	for name, wins := range s.score {
		league = append(league, Player{Name: name, Wins: wins})
	}
	return league
}
