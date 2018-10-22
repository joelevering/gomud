package mocks

import (
  "net"

 "github.com/joelevering/gomud/interfaces"
 "github.com/joelevering/gomud/structs"
)

type MockPlayer struct {
  *MockCharacter

  Defeated  interfaces.Combatant
  DefeatedBy  interfaces.Combatant
}

func NewMockPlayer() *MockPlayer {
  return &MockPlayer{
    MockCharacter: &MockCharacter{},
  }
}

func (m *MockPlayer) StartWriter(conn net.Conn) {}
func (m *MockPlayer) Init() {}
func (m *MockPlayer) Save() {}
func (m *MockPlayer) List() {}
func (m *MockPlayer) Look() {}
func (m *MockPlayer) LookNP(string) {}
func (m *MockPlayer) Status() {}
func (m *MockPlayer) AttackNP(string) {}
func (m *MockPlayer) Move(string) {}
func (m *MockPlayer) Say(string) {}
func (m *MockPlayer) Yell(string) {}
func (m *MockPlayer) SendMsg(...string) {}
func (m *MockPlayer) LeaveRoom(string) {}
func (m *MockPlayer) EnterRoom(interfaces.RoomI) {}
func (m *MockPlayer) GetName() string { return "mock player" }
func (m *MockPlayer) SetName(name string) {}
func (m *MockPlayer) GetID() string { return "mock ID" }
func (m *MockPlayer) GetRoom() interfaces.RoomI { return nil }

func (m *MockPlayer) EnterCombat(opp interfaces.Combatant) {}
func (m *MockPlayer) IsDefeated() bool { return false }
func (m *MockPlayer) LoseCombat(winner interfaces.Combatant) {
  m.DefeatedBy = winner
}

func (m *MockPlayer) WinCombat(loser interfaces.Combatant) {
  m.Defeated = loser
}

func (m *MockPlayer) AtkFx() structs.CmbFx {
  return structs.CmbFx{}
}

func (m *MockPlayer) ResistAtk(fx structs.CmbFx) structs.CmbFx {
  return fx
}

func (m *MockPlayer) ReportAtk(opp interfaces.Combatant, fx structs.CmbFx) {}
func (m *MockPlayer) ReportDef(opp interfaces.Combatant, fx structs.CmbFx) {}
