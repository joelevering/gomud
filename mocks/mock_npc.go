package mocks

import "github.com/joelevering/gomud/interfaces"

type MockNPC struct {
  IsDead bool
  DefeatedBy interfaces.CharI
}

func (m *MockNPC) Init(room interfaces.RoomI, queue interfaces.QueueI) {}
func (m *MockNPC) GetCharacter() interfaces.CharI { return &MockCharacter{} }

func (m *MockNPC) GetName() string { return "mock npc name" }
func (m *MockNPC) GetDesc() string { return "mock npc desc" }

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
