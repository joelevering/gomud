package mocks

import (
  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/structs"
)

type MockNP struct {
  *MockCharacter

  IsDead bool
  DefeatedBy interfaces.Combatant
}

func NewMockNP() *MockNP {
  return &MockNP{
    MockCharacter: &MockCharacter{},
  }
}

func (m *MockNP) Init(room interfaces.RoomI, queue interfaces.QueueI) {}
func (m *MockCharacter) SetClass() {}
func (m *MockCharacter) ResetStats() {}

func (m *MockNP) GetName() string { return "mock np name" }
func (m *MockNP) GetDesc() string { return "mock np desc" }

func (m *MockNP) SetSpawn(spawn interfaces.RoomI) {}
func (m *MockNP) SetBehavior(interfaces.QueueI) {}

func (m *MockNP) Spawn() {}
func (m *MockNP) IsAlive() bool { return !m.IsDead }

func (m *MockNP) Say(string) {}
func (m *MockNP) Emote(string) {}

func (m *MockNP) EnterCombat(opp interfaces.Combatant) { m.EnteredCombat = true }
func (m *MockNP) AtkFx() structs.CmbFx {
  return structs.CmbFx{}
}
func (m *MockNP) ResistAtk(fx structs.CmbFx) structs.CmbFx {
  return fx
}
func (m *MockNP) IsDefeated() bool { return false }
func (m *MockNP) ReportAtk(opp interfaces.Combatant, fx structs.CmbFx) {}
func (m *MockNP) ReportDef(opp interfaces.Combatant, fx structs.CmbFx) {}
func (m *MockNP) LoseCombat(opp interfaces.Combatant) {
  m.IsDead = true
  m.DefeatedBy = opp
}
func (m *MockNP) WinCombat(opp interfaces.Combatant) {}

