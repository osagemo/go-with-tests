package game

import (
	"encoding/json"
	"io"
)

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func (l League) ToJson() ([]byte, error) {
	s, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func LeagueFromJson(rdr io.Reader) (League, error) {
	var league []Player

	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		return nil, err
	}

	return league, nil
}
