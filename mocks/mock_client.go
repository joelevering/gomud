package mocks

import (
  "net"

 "github.com/joelevering/gomud/interfaces"
)

type MockClient struct {
  Defeated  interfaces.CharI
  DefeatedBy  interfaces.CharI
}

func (m *MockClient) StartWriter(conn net.Conn) {}
func (m *MockClient) List() {}
func (m *MockClient) Look() {}
func (m *MockClient) LookNPC(string) {}
func (m *MockClient) Status() {}
func (m *MockClient) AttackNPC(string) {}
func (m *MockClient) Move(string) {}
func (m *MockClient) Say(string) {}
func (m *MockClient) Yell(string) {}
func (m *MockClient) SendMsg(...string) {}
func (m *MockClient) LeaveRoom(string) {}
func (m *MockClient) EnterRoom(interfaces.RoomI) {}
func (m *MockClient) GetName() string { return "mock client" }
func (m *MockClient) SetName(name string) {}
func (m *MockClient) GetRoom() interfaces.RoomI { return nil }

func (m *MockClient) GetCombatCmd() []string { return []string{} }
func (m *MockClient) SetCombatCmd(cmd []string) {}

func (m *MockClient) LoseCombat(npc interfaces.CharI) {
  m.DefeatedBy = npc
}

func (m *MockClient) WinCombat(npc interfaces.CharI) {
  m.Defeated = npc
}
