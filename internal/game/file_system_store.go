package game

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

type FileSystemPlayerStore struct {
	file   *os.File
	league League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initFile(file)
	if err != nil {
		return nil, err
	}
	league, err := LeagueFromJson(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file: %v", err)
	}
	return &FileSystemPlayerStore{file, league}, nil
}

func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		closeFunc()
		return nil, nil, fmt.Errorf("problem creating file system player store: %v", err)
	}

	return store, closeFunc, nil
}

func initFile(file *os.File) error {
	info, err := file.Stat()
	if info.Size() == 0 {
		file.Write([]byte("[]"))
	}
	file.Seek(0, io.SeekStart)

	if err != nil {
		return fmt.Errorf("problem reading file: %v", err)
	}
	return nil
}

func (s *FileSystemPlayerStore) GetPlayerScore(player string) int {
	l := s.GetLeague()
	p := l.Find(player)
	if p != nil {
		return p.Wins
	}

	return 0
}

func (s *FileSystemPlayerStore) RecordWin(player string) {
	l := s.GetLeague()

	p := l.Find(player)
	if p != nil {
		p.Wins++
	} else {
		l = append(l, Player{player, 1})
	}
	s.Save(l)
}

func (s *FileSystemPlayerStore) GetLeague() League {
	leagueCopy := make(League, len(s.league))
	copy(leagueCopy, s.league)
	slices.SortFunc(leagueCopy, func(x, y Player) int {
		return y.Wins - x.Wins
	})
	return leagueCopy
}

func (s *FileSystemPlayerStore) Save(league League) {
	json, err := league.ToJson()
	if err != nil {
		panic(err)
	}

	s.file.Truncate(0)
	s.file.Seek(0, io.SeekStart)
	s.file.Write(json)
	s.league = league
}
