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
func (m *MockPlayer) LookTarget(string) {}
func (m *MockPlayer) Status() {}
func (m *MockPlayer) AttackNP(_, _ string) {}
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

func (m *MockPlayer) ReportAtk(_ interfaces.Combatant, _ structs.CmbRep) {}
func (m *MockPlayer) ReportDef(_ interfaces.Combatant, _ structs.CmbRep) {}
