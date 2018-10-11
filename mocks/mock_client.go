package mocks

import (
  "net"

 "github.com/joelevering/gomud/interfaces"
)

type MockPlayer struct {
  Defeated  interfaces.CharI
  DefeatedBy  interfaces.CharI
}

func (m *MockPlayer) StartWriter(conn net.Conn) {}
func (m *MockPlayer) List() {}
func (m *MockPlayer) Look() {}
func (m *MockPlayer) LookNPC(string) {}
func (m *MockPlayer) Status() {}
func (m *MockPlayer) AttackNPC(string) {}
func (m *MockPlayer) Move(string) {}
func (m *MockPlayer) Say(string) {}
func (m *MockPlayer) Yell(string) {}
func (m *MockPlayer) SendMsg(...string) {}
func (m *MockPlayer) LeaveRoom(string) {}
func (m *MockPlayer) EnterRoom(interfaces.RoomI) {}
func (m *MockPlayer) GetName() string { return "mock player" }
func (m *MockPlayer) SetName(name string) {}
func (m *MockPlayer) GetRoom() interfaces.RoomI { return nil }

func (m *MockPlayer) GetCombatCmd() []string { return []string{} }
func (m *MockPlayer) SetCombatCmd(cmd []string) {}

func (m *MockPlayer) LoseCombat(npc interfaces.CharI) {
  m.DefeatedBy = npc
}

func (m *MockPlayer) WinCombat(npc interfaces.CharI) {
  m.Defeated = npc
}
