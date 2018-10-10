package classes

import "github.com/joelevering/gomud/stats"

var ByName = map[string]*Class{
  "Conscript":   Conscript,
  "Athlete":     Athlete,
  "Charmer":     Charmer,
  "Augur":       Augur,
  "Sophist":     Sophist,
  "Slime King":  SlimeKing,
  "Slime Chef":  SlimeChef,
  "Slime Guard": SlimeGuard,
  "Slime Baby":  SlimeBaby,
  "Dog God":     DogGod,
  "Crab":        Crab,
}

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

var SlimeKing = &Class{
  name: "Slime King",
  growth: StatGrowth{},
  atkStats: []stats.Stat{stats.Kno},
  defStats: []stats.Stat{stats.Sag},
}

var SlimeChef = &Class{
  name: "Slime Chef",
  growth: StatGrowth{},
  atkStats: []stats.Stat{stats.Ing},
  defStats: []stats.Stat{stats.Ing},
}

var SlimeGuard = &Class{
  name: "Slime Guard",
  growth: StatGrowth{},
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Str},
}

var SlimeBaby = &Class{
  name: "Slime Baby",
  growth: StatGrowth{},
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Flo},
}

var DogGod = &Class{
  name: "Dog God",
  growth: StatGrowth{},
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Flo},
}

var Crab = &Class{
  name: "Crab",
  growth: StatGrowth{},
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Flo},
}
