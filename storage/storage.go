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
    PlayersData: make(map[string]*PlayerData),
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

type Storage struct {
  PlayersData map[string]*PlayerData `json:"players_data"`
  Filename    string
}

type PlayerData struct {
  Classes map[string]ClassStats `json:classes`
}

type ClassStats struct {
  Lvl        int `json:"level"`
  MaxDet     int `json:"max_health"`
  Exp        int `json:"experience"`
  NextLvlExp int `json:"next_level_exp"`
}

func (s *Storage) StoreExists(pID string) bool {
  return s.PlayersData[pID] != nil
}

func (s *Storage) InitPlayerData(pID string) {
  if s.PlayersData[pID] == nil {
    s.PlayersData[pID] = &PlayerData{
      Classes: make(map[string]ClassStats),
    }
  }
}

func (s *Storage) PersistClass(pID, className string, stats ClassStats) {
  s.PlayersData[pID].Classes[className] = stats

  s.PersistStore(s.Filename)
}

func (s *Storage) LoadStats(pID, className string) ClassStats {
  return s.PlayersData[pID].Classes[className]
}

func (s *Storage) PersistStore(filename string) {
  j, _ := json.Marshal(s)
  err := ioutil.WriteFile(filename, j, 0644)
  if err != nil {
    log.Printf("storage:persist - %s", err)
  }
}
