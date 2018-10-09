package classes

import "github.com/joelevering/gomud/stats"

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
