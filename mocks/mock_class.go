package mocks

import (
  "github.com/joelevering/gomud/classes"
  "github.com/joelevering/gomud/interfaces"
)

type MockClass struct {
  LeveledUpChar interfaces.CharI
}

func (m *MockClass) GetName() string { return "Mock Class" }
func (m *MockClass) GetStatGrowth() classes.StatGrowth {
  return classes.StatGrowth{
    Health: 10,
    Str: 1,
    End: 2,
  }
}
