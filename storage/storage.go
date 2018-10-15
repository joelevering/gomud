package storage

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "os"
)

func LoadStore(filename string) *Storage {
  f, err := os.OpenFile(filename, os.O_CREATE, 0644)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  b, err := ioutil.ReadAll(f)
  if err != nil {
    panic(err)
  }

  s := &Storage{
    PlayersClasses: make(map[string]map[string]ClassStats),
    Filename: filename,
  }

  if len(b) == 0 {
    s.PersistStore(filename)
    return s
  }

  err = json.Unmarshal(b, s)
  if err != nil {
    panic(err)
  }

  return s
}

// Stores player name to a map of class name to stats
type Storage struct {
  PlayersClasses map[string]map[string]ClassStats `json:"players_classes"`
  Filename       string
}

type ClassStats struct {
  Lvl        int `json:"level"`
  MaxDet     int `json:"max_health"`
  Exp        int `json:"experience"`
  NextLvlExp int `json:"next_level_exp"`
}

func (s *Storage) StoreExists(pID string) bool {
  return s.PlayersClasses[pID] != nil
}

func (s *Storage) InitStats(pID string) {
  if s.PlayersClasses[pID] == nil {
    s.PlayersClasses[pID] = make(map[string]ClassStats)
  }
}

func (s *Storage) PersistClass(pID, className string, stats ClassStats) {
  s.PlayersClasses[pID][className] = stats

  s.PersistStore(s.Filename)
}

func (s *Storage) LoadStats(pID, className string) ClassStats {
  return s.PlayersClasses[pID][className]
}

func (s *Storage) PersistStore(filename string) {
  j, _ := json.Marshal(s)
  err := ioutil.WriteFile(filename, j, 0644)
  if err != nil {
    log.Printf("storage:persist - %s", err)
  }
}
