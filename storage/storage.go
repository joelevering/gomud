package storage

import "github.com/joelevering/gomud/interfaces"

// Stores player name to a map of class name to stats
var store = &Storage{
  PlayersClasses: make(map[string]map[string]ClassStats),
}

type ClassStats struct {
  Lvl        int
  MaxDet     int
  Exp        int
  NextLvlExp int
}

type Storage struct {
  PlayersClasses map[string]map[string]ClassStats
}

func CreateStore(pID string) {
  store.PlayersClasses[pID] = make(map[string]ClassStats)
}

func PersistClass(p interfaces.PlI, className string) {
  store.PlayersClasses[p.GetID()][className] = ClassStats{
    Lvl: p.GetLevel(),
    MaxDet: p.GetMaxDet(),
    Exp: p.GetExp(),
    NextLvlExp: p.GetNextLvlExp(),
  }
}

func LoadStats(pID, className string) ClassStats {
  return store.PlayersClasses[pID][className]
}
