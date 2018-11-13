package classes

import (
  "strings"

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

type ClassSkill struct {
  Skill *skills.Skill
  Level int
}

type Tier int

const(
  Tier1 = Tier(0)
  Tier2 = Tier(1)
  Tier3 = Tier(2)
  Tier4 = Tier(3)
  Tier5 = Tier(4)
)

type Class struct {
  name     string
  desc     string
  tier     Tier
  growth   StatGrowth
  atkStats []stats.Stat
  defStats []stats.Stat
  skills   []*ClassSkill
  skillMap map[string]*ClassSkill
  reqs     []*Class
}

func (c *Class) GetName() string {
  return c.name
}

func (c *Class) GetDesc() string {
  return c.desc
}

func (c *Class) GetTier() Tier {
  return c.tier
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

func (c *Class) GetSkills() []*ClassSkill {
  return c.skills
}

func (c *Class) GetReqs() []*Class {
  return c.reqs
}

func (c *Class) SkillForLvl(lvl int) *skills.Skill {
  for _, cs := range c.skills {
    if lvl == cs.Level {
      return cs.Skill
    }
  }

  return nil
}

func (c *Class) SkillsForLvl(lvl int) []*skills.Skill {
  skills := []*skills.Skill{}
  for _, cs := range c.skills {
    if lvl >= cs.Level {
      skills = append(skills, cs.Skill)
    }
  }
  return skills
}

func (c *Class) GetSkill(s string, lvl int) *skills.Skill {
  if c.skillMap == nil {
    c.generateSkillMap()
  }

  cs := c.skillMap[strings.Title(s)]
  if cs != nil && lvl >= cs.Level {
    return cs.Skill
  } else {
    return nil
  }
}

func (c *Class) generateSkillMap() {
  c.skillMap = make(map[string]*ClassSkill)

  for _, cs := range c.skills {
    c.skillMap[cs.Skill.Name] = cs
  }
}
