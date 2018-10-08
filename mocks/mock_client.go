package mocks

import (
  "net"

 "github.com/joelevering/gomud/interfaces"
)

type MockClient struct {
  MaxHealth int
  Health    int
  Str       int
  End       int
  CombatCmd []string

  Defeated  interfaces.NPCI
  KilledBy  interfaces.NPCI
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
func (m *MockClient) GetName() string { return "" }
func (m *MockClient) GetRoom() interfaces.RoomI { return nil }

func (m *MockClient) GetHealth() int { return m.Health }
func (m *MockClient) SetHealth(health int) { m.Health = health }
func (m *MockClient) GetMaxHealth() int { return m.MaxHealth }
func (m *MockClient) SetMaxHealth(maxHealth int) { m.MaxHealth = maxHealth }
func (m *MockClient) GetStr() int { return m.Str }
func (m *MockClient) SetStr(str int) { m.Str = str }
func (m *MockClient) GetEnd() int { return m.End }
func (m *MockClient) SetEnd(end int) { m.End = end }
func (m *MockClient) GetCombatCmd() []string { return m.CombatCmd }

func (m *MockClient) Die(npc interfaces.NPCI) {
  m.KilledBy = npc
}

func (m *MockClient) Defeat(npc interfaces.NPCI) {
  m.Defeated = npc
}
