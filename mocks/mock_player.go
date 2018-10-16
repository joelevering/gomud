package mocks

import (
  "net"

 "github.com/joelevering/gomud/interfaces"
)

type MockPlayer struct {
  *MockCharacter

  Defeated  interfaces.CharI
  DefeatedBy  interfaces.CharI
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

func (m *MockPlayer) LoseCombat(npc interfaces.CharI) {
  m.DefeatedBy = npc
}

func (m *MockPlayer) WinCombat(npc interfaces.CharI) {
  m.Defeated = npc
}
