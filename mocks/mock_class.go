package mocks

import (
  "github.com/joelevering/gomud/classes"
  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/stats"
)

type MockClass struct {
  LeveledUpChar interfaces.CharI
}

func (m *MockClass) GetName() string { return "Mock Class" }
func (m *MockClass) GetTier() classes.Tier { return 0 }

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

func (m *MockClass) SkillForLvl(_ int) *skills.Skill {
  return skills.Charge
}

func (m *MockClass) SkillsForLvl(_ int) []*skills.Skill {
  return []*skills.Skill{}
}

func (m *MockClass) GetSkill(_ string, _ int) *skills.Skill {
  return skills.Shove
}
