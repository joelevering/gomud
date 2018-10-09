package classes

import "github.com/joelevering/gomud/stats"

var Conscript = &Class{
  name: "Conscript",
  growth: StatGrowth{
    Det: 25,
    Stm: 10,
    Str: 10,
  },
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Str},
}
