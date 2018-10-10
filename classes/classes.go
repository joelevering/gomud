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

var Athlete = &Class{
  name: "Athlete",
  growth: StatGrowth{
    Det: 25,
    Stm: 10,
    Flo: 10,
  },
  atkStats: []stats.Stat{stats.Flo},
  defStats: []stats.Stat{stats.Flo},
}

var Charmer = &Class{
  name: "Charmer",
  growth: StatGrowth{
    Det: 25,
    Stm: 5,
    Foc: 5,
    Ing: 10,
  },
  atkStats: []stats.Stat{stats.Ing},
  defStats: []stats.Stat{stats.Ing},
}

var Augur = &Class{
  name: "Augur",
  growth: StatGrowth{
    Det: 25,
    Foc: 10,
    Kno: 10,
  },
  atkStats: []stats.Stat{stats.Kno},
  defStats: []stats.Stat{stats.Kno},
}

var Sophist = &Class{
  name: "Sophist",
  growth: StatGrowth{
    Det: 25,
    Foc: 10,
    Sag: 10,
  },
  atkStats: []stats.Stat{stats.Sag},
  defStats: []stats.Stat{stats.Sag},
}
