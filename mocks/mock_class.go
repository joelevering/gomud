package mocks

import (
  "github.com/joelevering/gomud/classes"
  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/stats"
)

type MockClass struct {
  LeveledUpChar interfaces.CharI
}

func (m *MockClass) GetName() string { return "Mock Class" }

func (m *MockClass) GetStatGrowth() classes.StatGrowth {
  return classes.StatGrowth{
    Det: 10,
    Str: 1,
    Flo: 2,
  }
}

func (m *MockClass) GetAtkStats() []stats.Stat {
  return []stats.Stat{stats.Str}
}

func (m *MockClass) GetDefStats() []stats.Stat {
  return []stats.Stat{stats.Flo}
}
