package classes

import (
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/stats"
)

var PlayerClasses = []*Class{
  Conscript,
  Athlete,
  Charmer,
  Augur,
  Sophist,
}

var ByName = map[string]*Class{
  "Conscript":   Conscript,
  "Athlete":     Athlete,
  "Charmer":     Charmer,
  "Augur":       Augur,
  "Sophist":     Sophist,
  "Slime":       Slime,
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
  skills: []*ClassSkill{
    &ClassSkill{
      Skill: skills.Shove,
      Level: 2,
    },
    &ClassSkill{
      Skill: skills.Charge,
      Level: 4,
    },
    &ClassSkill{
      Skill: skills.DesperateBlow,
      Level: 6,
    },
    &ClassSkill{
      Skill: skills.PowerNap,
      Level: 8,
    },
    &ClassSkill{
      Skill: skills.Conserve,
      Level: 10,
    },
  },
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
  skills: []*ClassSkill{
    &ClassSkill{
      Skill: skills.FreneticPace,
      Level: 2,
    },
    &ClassSkill{
      Skill: skills.BackUp,
      Level: 4,
    },
    &ClassSkill{
      Skill: skills.Duck,
      Level: 6,
    },
    &ClassSkill{
      Skill: skills.Counter,
      Level: 8,
    },
    &ClassSkill{
      Skill: skills.Uppercut,
      Level: 10,
    },
  },
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

var Crab = &Class{
  name: "Crab",
  growth: StatGrowth{
    Det: 10,
    Str: 1,
    Flo: 2,
  },
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Flo},
}

// Slime Forest

var Slime = &Class{
  name: "Slime",
  growth: StatGrowth{
    Det: 50,
    Str: 10,
    Ing: 5,
  },
  atkStats: []stats.Stat{stats.Ing},
  defStats: []stats.Stat{stats.Str},
}

// Slime Castle

var SlimeKing = &Class{
  name: "Slime King",
  growth: StatGrowth{
    Det: 99999,
    Stm: 99999,
    Foc: 99999,
    Str: 999,
    Flo: 999,
    Ing: 999,
    Kno: 999,
    Sag: 999,
  },
  atkStats: []stats.Stat{stats.Kno},
  defStats: []stats.Stat{stats.Sag},
}

var SlimeChef = &Class{
  name: "Slime Chef",
  growth: StatGrowth{
    Det: 88888,
    Stm: 88888,
    Foc: 88888,
    Str: 888,
    Flo: 888,
    Ing: 888,
    Kno: 888,
    Sag: 888,
  },
  atkStats: []stats.Stat{stats.Ing},
  defStats: []stats.Stat{stats.Ing},
}

var SlimeGuard = &Class{
  name: "Slime Guard",
  growth: StatGrowth{
    Det: 77777,
    Stm: 77777,
    Foc: 77777,
    Str: 777,
    Flo: 777,
    Ing: 777,
    Kno: 777,
    Sag: 777,
  },
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Str},
}

var SlimeBaby = &Class{
  name: "Slime Baby",
  growth: StatGrowth{
    Det: 66666,
    Stm: 66666,
    Foc: 66666,
    Str: 666,
    Flo: 666,
    Ing: 666,
    Kno: 666,
    Sag: 666,
  },
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Flo},
}

var DogGod = &Class{
  name: "Dog God",
  growth: StatGrowth{
    Det: 100000,
    Stm: 100000,
    Foc: 100000,
    Str: 1000,
    Flo: 1000,
    Ing: 1000,
    Kno: 1000,
    Sag: 1000,
  },
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Flo},
}
