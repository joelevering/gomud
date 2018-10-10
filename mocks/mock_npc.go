package mocks

import (
 "github.com/joelevering/gomud/interfaces"
)

type MockNPC struct {
  Det    int
  Str    int
  Flo    int
  Exp    int
  IsDead bool
  Atk    int
  Def    int

  SetDetArg  int
  DefeatedBy interfaces.CharI
}

func (m *MockNPC) GetName() string { return "" }
func (m *MockNPC) GetDesc() string { return "" }
func (m *MockNPC) GetDet() int { return m.Det }
func (m *MockNPC) SetDet(det int) { m.SetDetArg = det }
func (m *MockNPC) GetMaxDet() int { return 100 }
func (m *MockNPC) GetStr() int { return m.Str }
func (m *MockNPC) GetExp() int { return 10 }
func (m *MockNPC) GetAtk() int { return m.Atk }
func (m *MockNPC) GetDef() int { return m.Def }

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
