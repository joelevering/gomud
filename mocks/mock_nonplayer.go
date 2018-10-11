package mocks

import "github.com/joelevering/gomud/interfaces"

type MockNP struct {
  IsDead bool
  DefeatedBy interfaces.CharI
}

func (m *MockNP) Init(room interfaces.RoomI, queue interfaces.QueueI) {}
func (m *MockNP) GetCharacter() interfaces.CharI { return &MockCharacter{} }

func (m *MockNP) GetName() string { return "mock np name" }
func (m *MockNP) GetDesc() string { return "mock np desc" }

func (m *MockNP) SetSpawn(spawn interfaces.RoomI) {}
func (m *MockNP) SetBehavior(interfaces.QueueI) {}

func (m *MockNP) Spawn() {}
func (m *MockNP) IsAlive() bool { return !m.IsDead }

func (m *MockNP) Say(string) {}
func (m *MockNP) Emote(string) {}

func (m *MockNP) LoseCombat(ch interfaces.CharI) {
  m.IsDead = true
  m.DefeatedBy = ch
}
