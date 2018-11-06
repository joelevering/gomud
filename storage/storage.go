package storage

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "os"

  "github.com/joelevering/gomud/character"
)

type StorageI interface {
  StoreExists(string) bool
  InitPlayerData(string)
  PersistClass(string, string, ClassStats)
  PersistChar(string, *character.Character)
  LoadStats(string, string) ClassStats
  LoadClasses(string) map[string]ClassStats
  LoadChar(string) CharStats
  PersistStore(string)
}

type ClassStats struct {
  Lvl        int `json:"level"`
  MaxDet     int `json:"max_health"`
  Exp        int `json:"experience"`
  NextLvlExp int `json:"next_level_exp"`
}

type CharStats struct {
  Det    int `json:"health"`
  MaxStm int `json:"max_stamina"`
  Stm    int `json:"stamina"`
  MaxFoc int `json:"max_focus"`
  Foc    int `json:"focus"`
  Str    int `json:"strength"`
  Flo    int `json:"flow"`
  Ing    int `json:"ingenuity"`
  Kno    int `json:"knowledge"`
  Sag    int `json:"sagacity"`

  Room   int `json:"room"`
  Spawn  int `json:"spawn"`
}

type Storage struct {
  PlayersData map[string]*PlayerData `json:"players_data"`
  Filename    string
}

type PlayerData struct {
  Classes   map[string]ClassStats `json:classes`
  Character CharStats             `json:character`
}

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

func (s *Storage) StoreExists(pID string) bool {
  return s.PlayersData[pID] != nil
}

func (s *Storage) InitPlayerData(pID string) {
  if s.PlayersData[pID] == nil {
    s.PlayersData[pID] = &PlayerData{
      Classes: make(map[string]ClassStats),
      Character: CharStats{
        Room: -1,
        Spawn: -1,
      },
    }
  }
}

func (s *Storage) PersistClass(pID, className string, stats ClassStats) {
  s.PlayersData[pID].Classes[className] = stats

  s.PersistStore(s.Filename)
}

func (s *Storage) PersistChar(pID string, ch *character.Character) {
  s.PlayersData[pID].Character = CharStats{
    Det: ch.Det,
    MaxStm: ch.MaxStm,
    Stm: ch.Stm,
    MaxFoc: ch.MaxFoc,
    Foc: ch.Foc,
    Str: ch.Str,
    Flo: ch.Flo,
    Ing: ch.Ing,
    Kno: ch.Kno,
    Sag: ch.Sag,
    Room: ch.Room.GetID(),
    Spawn: ch.GetSpawn().GetID(),
  }

  s.PersistStore(s.Filename)
}

func (s *Storage) LoadStats(pID, className string) ClassStats {
  return s.PlayersData[pID].Classes[className]
}

func (s *Storage) LoadClasses(pID string) map[string]ClassStats {
  return s.PlayersData[pID].Classes
}

func (s *Storage) LoadChar(pID string) CharStats {
  return s.PlayersData[pID].Character
}

func (s *Storage) PersistStore(filename string) {
  j, err := json.Marshal(s)
  if err != nil {
    log.Printf("storage:persist:marshal - %s", err)
  }
  err = ioutil.WriteFile(filename, j, 0644)
  if err != nil {
    log.Printf("storage:persist:write - %s", err)
  }
}
