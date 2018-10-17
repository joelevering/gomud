package classes

import (
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/stats"
)

type StatGrowth struct {
  Det int
  Stm int
  Foc int

  Str int
  Flo int
  Ing int
  Kno int
  Sag int
}

type Class struct {
  name     string
  growth   StatGrowth
  atkStats []stats.Stat
  defStats []stats.Stat
  skills   []*skills.Skill
  skillMap map[string]*skills.Skill
}

func (c *Class) GetName() string {
  return c.name
}

func (c *Class) GetStatGrowth() StatGrowth {
  return c.growth
}

func (c *Class) GetAtkStats() []stats.Stat {
  return c.atkStats
}

func (c *Class) GetDefStats() []stats.Stat {
  return c.defStats
}

func (c *Class) GetSkill(s string) *skills.Skill {
  if c.skillMap == nil {
    c.GenerateSkillMap()
  }

  return c.skillMap[s]
}

func (c *Class) GenerateSkillMap() {
  c.skillMap = make(map[string]*skills.Skill)

  for _, s := range c.skills {
    c.skillMap[s.Name] = s
  }
}
