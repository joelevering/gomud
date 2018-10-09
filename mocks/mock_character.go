package mocks

import "github.com/joelevering/gomud/interfaces"

type MockCharacter struct {
  Spawn      interfaces.RoomI
  InCombat   bool

  Atk           int
  Def           int
  SetDetArg     int
  SetMaxDetArg  int
  SetStrArg     int
  SetFloArg     int
  LeveledUp     bool
  Healed        bool
  ExpGained     int
  ShouldLevelUp bool
}

func (m *MockCharacter) GetClassName() string { return "Superstar" }
func (m *MockCharacter) GetName() string { return "Heide" }
func (m *MockCharacter) SetName(name string) {}
func (m *MockCharacter) GetDet() int { return 199 }
func (m *MockCharacter) SetDet(det int) { m.SetDetArg = det }
func (m *MockCharacter) GetMaxDet() int { return 200 }
func (m *MockCharacter) SetMaxDet(maxDet int) { m.SetMaxDetArg = maxDet }
func (m *MockCharacter) GetMaxStm() int { return 100 }
func (m *MockCharacter) SetMaxStm(maxStm int) {}
func (m *MockCharacter) GetMaxFoc() int { return 100 }
func (m *MockCharacter) SetMaxFoc(maxFoc int) {}
func (m *MockCharacter) GetStr() int { return 30 }
func (m *MockCharacter) SetStr(str int) { m.SetStrArg = str }
func (m *MockCharacter) GetFlo() int { return 50 }
func (m *MockCharacter) SetFlo(flo int) { m.SetFloArg = flo }
func (m *MockCharacter) GetIng() int { return 50 }
func (m *MockCharacter) SetIng(ing int) {}
func (m *MockCharacter) GetKno() int { return 50 }
func (m *MockCharacter) SetKno(kno int) {}
func (m *MockCharacter) GetSag() int { return 50 }
func (m *MockCharacter) SetSag(sag int) {}
func (m *MockCharacter) GetAtk() int { return m.Atk }
func (m *MockCharacter) GetDef() int { return m.Def }
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
