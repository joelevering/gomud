package mocks

import "github.com/joelevering/gomud/interfaces"

type MockClass struct {
  LeveledUpChar interfaces.CharI
}

func (m *MockClass) GetName() string { return "Mock Class" }
func (m *MockClass) LevelUp(ch interfaces.CharI) { m.LeveledUpChar = ch }
