package mocks

import (
 "github.com/joelevering/gomud/interfaces"
)

type MockNPC struct {
  MaxHealth int
  Health    int
  Str       int
  End       int
  Exp       int
  IsDead    bool

  DefeatedBy  interfaces.CharI
}

func (m *MockNPC) GetName() string { return "" }
func (m *MockNPC) GetDesc() string { return "" }
func (m *MockNPC) GetHealth() int { return m.Health }
func (m *MockNPC) SetHealth(health int) { m.Health = health }
func (m *MockNPC) GetMaxHealth() int { return m.MaxHealth }
func (m *MockNPC) GetStr() int { return m.Str }
func (m *MockNPC) GetEnd() int { return m.End }
func (m *MockNPC) GetExp() int { return m.Exp }

func (m *MockNPC) SetSpawn(spawn interfaces.RoomI) {}
func (m *MockNPC) SetBehavior(interfaces.QueueI) {}

func (m *MockNPC) Spawn() {}
func (m *MockNPC) IsAlive() bool { return !m.IsDead }

func (m *MockNPC) Say(string) {}
func (m *MockNPC) Emote(string) {}

func (m *MockNPC) LoseCombat(ch interfaces.CharI) {
  m.IsDead = true
  m.DefeatedBy = ch
}
