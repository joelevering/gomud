package mocks

import "github.com/joelevering/gomud/interfaces"

type MockCharacter struct {
  Spawn      interfaces.RoomI
  InCombat   bool

  SetMaxHealthArg int
  SetStrArg int
  SetEndArg int
  LeveledUp     bool
  Healed        bool
  ExpGained     int
  ShouldLevelUp bool
}

func (m *MockCharacter) GetName() string { return "Heide" }
func (m *MockCharacter) SetName(name string) {}
func (m *MockCharacter) GetHealth() int { return 199 }
func (m *MockCharacter) SetHealth(health int) {}
func (m *MockCharacter) GetMaxHealth() int { return 200 }
func (m *MockCharacter) SetMaxHealth(maxHealth int) { m.SetMaxHealthArg = maxHealth }
func (m *MockCharacter) GetStr() int { return 30 }
func (m *MockCharacter) SetStr(str int) { m.SetStrArg = str }
func (m *MockCharacter) GetEnd() int { return 50 }
func (m *MockCharacter) SetEnd(end int) { m.SetEndArg = end }
func (m *MockCharacter) GetLevel() int { return 2 }
func (m *MockCharacter) GetExp() int { return 666 }
func (m *MockCharacter) GetNextLvlExp() int { return 1000 }
func (m *MockCharacter) GetSpawn() interfaces.RoomI { return m.Spawn }
func (m *MockCharacter) SetSpawn(spawn interfaces.RoomI) {}

func (m *MockCharacter) EnterCombat() {}
func (m *MockCharacter) LeaveCombat() {}
func (m *MockCharacter) IsInCombat() bool { return m.InCombat }
func (m *MockCharacter) ExpToLvl() int { return 100 }

func (m *MockCharacter) Heal() {
  m.Healed = true
}

func (m *MockCharacter) GainExp(exp int) bool {
  m.ExpGained += exp
  return m.ShouldLevelUp
}
