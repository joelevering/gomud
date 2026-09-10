package classes

import (
  "strings"

  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/stats"
)

var StartingClasses = []*Class{
  Conscript,
  Athlete,
  Charmer,
  Augur,
  Sophist,
}

var PlayerClasses = []*Class{
  Conscript,
  Athlete,
  Charmer,
  Augur,
  Sophist,
  Minder,
}

var ByName = map[string]*Class{
  "Conscript":   Conscript,
  "Athlete":     Athlete,
  "Charmer":     Charmer,
  "Augur":       Augur,
  "Sophist":     Sophist,
  "Minder":      Minder,
  "Slime":       Slime,
  "Slime King":  SlimeKing,
  "Slime Chef":  SlimeChef,
  "Slime Guard": SlimeGuard,
  "Slime Baby":  SlimeBaby,
  "Dog God":     DogGod,
  "Crab":        Crab,

  // Fauna
  "Vermin": Vermin,
  "Vole": Vole,
  "Bird": Bird,
  "Goat": Goat,
  "Lizard": Lizard,
  "Amphibian": Amphibian,
  "Snake": Snake,
  "Insect": Insect,
  "Boar": Boar,
  "Plant": Plant,
  "Fungus": Fungus,

  // Constructs & Golems
  "Construct": Construct,
  "Golem": Golem,
  "Golem King": GolemKing,

  // Elementals & Spirits
  "Elemental": Elemental,
  "Elemental King": ElementalKing,
  "Air Elemental": AirElemental,
  "Dust Elemental": DustElemental,
  "Zephyr": Zephyr,
  "Spirit": Spirit,
  "Ghost": Ghost,
  "Undead": Undead,

  // Cultists & Shrine
  "Cultist": Cultist,
  "Cultist Acolyte": CultistAcolyte,
  "Cultist Zealot": CultistZealot,
  "Cultist Warband Scout": CultistWarbandScout,
  "Cultist High Priest": CultistHighPriest,
  "Rogue Cultist": RogueCultist,
  "Shrine Acolyte": ShrineAcolyte,
  "Shrine Devotee": ShrineDevotee,
  "Shrine Keeper": ShrineKeeper,

  // Seneca Downs townsfolk
  "Blacksmith": Blacksmith,
  "Guard": Guard,
  "Guard Captain": GuardCaptain,
  "Innkeeper": Innkeeper,
  "Merchant": Merchant,
  "Peddler": Peddler,
  "Shepherd": Shepherd,

  // Aggressive/elite variants
  "BirdSwarm": BirdSwarm,
  "Badger": Badger,
}

func Find(name string) *Class {
  return ByName[strings.Title(name)]
}

var Conscript = &Class{
  name: "Conscript",
  desc: "A strength-based class that's brutish attacks are often interrupted by bouts of laziness",
  fleeChance: 0.5,
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
  desc: "A flow-based fighter that uses speed and superior reaction time to chain attacks into damaging combos",
  fleeChance: 0.5,
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
  desc: "An unpredictable class that sabotages enemies with on-the-fly ingenuity",
  fleeChance: 0.5,
  growth: StatGrowth{
    Det: 25,
    Stm: 5,
    Foc: 5,
    Ing: 10,
  },
  atkStats: []stats.Stat{stats.Ing},
  defStats: []stats.Stat{stats.Ing},
  skills: []*ClassSkill{
    &ClassSkill{
      Skill: skills.WittyRetort,
      Level: 2,
    },
    &ClassSkill{
      Skill: skills.Ploy,
      Level: 4,
    },
    &ClassSkill{
      Skill: skills.Sideswipe,
      Level: 6,
    },
    &ClassSkill{
      Skill: skills.Plan,
      Level: 8,
    },
    &ClassSkill{
      Skill: skills.Sidestep,
      Level: 10,
    },
  },
}

var Augur = &Class{
  name: "Augur",
  desc: "This class uses research and knowledge to act with in the most effective manner possible",
  fleeChance: 0.5,
  growth: StatGrowth{
    Det: 25,
    Foc: 10,
    Kno: 10,
  },
  atkStats: []stats.Stat{stats.Kno},
  defStats: []stats.Stat{stats.Kno},
  skills: []*ClassSkill{
    &ClassSkill{
      Skill: skills.LowBlow,
      Level: 2,
    },
    &ClassSkill{
      Skill: skills.Spark,
      Level: 4,
    },
    &ClassSkill{
      Skill: skills.Concentrate,
      Level: 6,
    },
    &ClassSkill{
      Skill: skills.FirstAid,
      Level: 8,
    },
    &ClassSkill{
      Skill: skills.TargetedStrike,
      Level: 10,
    },
  },
}

var Sophist = &Class{
  name: "Sophist",
  desc: "This adaptable classes uses simple sagacious movements to disarm and destroy foes",
  fleeChance: 0.5,
  growth: StatGrowth{
    Det: 25,
    Foc: 10,
    Sag: 10,
  },
  atkStats: []stats.Stat{stats.Sag},
  defStats: []stats.Stat{stats.Sag},
  skills: []*ClassSkill{
    &ClassSkill{
      Skill: skills.CastDoubt,
      Level: 2,
    },
    &ClassSkill{
      Skill: skills.Benumb,
      Level: 4,
    },
    &ClassSkill{
      Skill: skills.Radiate,
      Level: 6,
    },
    &ClassSkill{
      Skill: skills.Redirect,
      Level: 8,
    },
    &ClassSkill{
      Skill: skills.CastOff,
      Level: 10,
    },
  },
}

var Minder = &Class{
  name: "Minder",
  desc: "This defensive class focuses on protection and healing",
  tier: Tier2,
  fleeChance: 0.5,
  growth: StatGrowth{
    Det: 25,
    Foc: 10,
    Kno: 10,
  },
  atkStats: []stats.Stat{stats.Kno},
  defStats: []stats.Stat{stats.Kno},
  skills: []*ClassSkill{
    &ClassSkill{
      Skill: skills.Shield,
      Level: 2,
    },
  },
  reqs: []*Class{
    Conscript,
    Augur,
  },
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
  skills: []*ClassSkill{
    &ClassSkill{
      Skill: skills.Pince,
      Level: 1,
    },
    &ClassSkill{
      Skill: skills.Hide,
      Level: 1,
    },
  },
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

// Fauna

var Vermin = &Class{
  name: "Vermin",
  growth: StatGrowth{
    Det: 18,
    Stm: 5,
    Foc: 5,
    Flo: 6,
  },
  atkStats: []stats.Stat{stats.Flo},
  defStats: []stats.Stat{stats.Flo},
}

var Vole = &Class{
  name: "Vole",
  growth: StatGrowth{
    Det: 12,
    Stm: 3,
    Foc: 3,
    Flo: 4,
  },
  atkStats: []stats.Stat{stats.Flo},
  defStats: []stats.Stat{stats.Flo},
}

var Bird = &Class{
  name: "Bird",
  growth: StatGrowth{
    Det: 14,
    Stm: 4,
    Foc: 4,
    Flo: 5,
  },
  atkStats: []stats.Stat{stats.Flo},
  defStats: []stats.Stat{stats.Flo},
}

var Goat = &Class{
  name: "Goat",
  growth: StatGrowth{
    Det: 14,
    Stm: 4,
    Foc: 3,
    Str: 4,
  },
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Str},
}

var Lizard = &Class{
  name: "Lizard",
  growth: StatGrowth{
    Det: 14,
    Stm: 4,
    Foc: 3,
    Flo: 4,
  },
  atkStats: []stats.Stat{stats.Flo},
  defStats: []stats.Stat{stats.Flo},
}

var Amphibian = &Class{
  name: "Amphibian",
  growth: StatGrowth{
    Det: 24,
    Stm: 7,
    Foc: 5,
    Str: 7,
  },
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Str},
}

var Snake = &Class{
  name: "Snake",
  growth: StatGrowth{
    Det: 20,
    Stm: 5,
    Foc: 5,
    Flo: 7,
  },
  atkStats: []stats.Stat{stats.Flo},
  defStats: []stats.Stat{stats.Flo},
}

var Insect = &Class{
  name: "Insect",
  growth: StatGrowth{
    Det: 24,
    Stm: 6,
    Foc: 6,
    Flo: 8,
  },
  atkStats: []stats.Stat{stats.Flo},
  defStats: []stats.Stat{stats.Flo},
}

var Boar = &Class{
  name: "Boar",
  growth: StatGrowth{
    Det: 20,
    Stm: 8,
    Foc: 4,
    Str: 8,
  },
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Str},
}

var Plant = &Class{
  name: "Plant",
  growth: StatGrowth{
    Det: 26,
    Stm: 8,
    Foc: 5,
    Str: 9,
  },
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Str},
}

var Fungus = &Class{
  name: "Fungus",
  growth: StatGrowth{
    Det: 18,
    Stm: 6,
    Foc: 4,
    Str: 5,
  },
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Str},
}

// Constructs & Golems

var Construct = &Class{
  name: "Construct",
  growth: StatGrowth{
    Det: 24,
    Stm: 6,
    Foc: 7,
    Kno: 8,
  },
  atkStats: []stats.Stat{stats.Kno},
  defStats: []stats.Stat{stats.Kno},
}

var Golem = &Class{
  name: "Golem",
  growth: StatGrowth{
    Det: 30,
    Stm: 10,
    Foc: 8,
    Kno: 10,
  },
  atkStats: []stats.Stat{stats.Kno},
  defStats: []stats.Stat{stats.Kno},
}

var GolemKing = &Class{
  name: "Golem King",
  growth: StatGrowth{
    Det: 60,
    Stm: 20,
    Foc: 18,
    Kno: 20,
  },
  atkStats: []stats.Stat{stats.Kno},
  defStats: []stats.Stat{stats.Kno},
}

// Elementals & Spirits

var Elemental = &Class{
  name: "Elemental",
  growth: StatGrowth{
    Det: 28,
    Stm: 8,
    Foc: 8,
    Sag: 9,
  },
  atkStats: []stats.Stat{stats.Sag},
  defStats: []stats.Stat{stats.Sag},
}

var ElementalKing = &Class{
  name: "Elemental King",
  growth: StatGrowth{
    Det: 65,
    Stm: 20,
    Foc: 20,
    Sag: 22,
  },
  atkStats: []stats.Stat{stats.Sag},
  defStats: []stats.Stat{stats.Sag},
}

var AirElemental = &Class{
  name: "Air Elemental",
  growth: StatGrowth{
    Det: 26,
    Stm: 7,
    Foc: 7,
    Flo: 9,
  },
  atkStats: []stats.Stat{stats.Flo},
  defStats: []stats.Stat{stats.Flo},
}

var DustElemental = &Class{
  name: "Dust Elemental",
  growth: StatGrowth{
    Det: 22,
    Stm: 6,
    Foc: 6,
    Flo: 8,
  },
  atkStats: []stats.Stat{stats.Flo},
  defStats: []stats.Stat{stats.Flo},
}

var Zephyr = &Class{
  name: "Zephyr",
  growth: StatGrowth{
    Det: 20,
    Stm: 6,
    Foc: 6,
    Flo: 8,
  },
  atkStats: []stats.Stat{stats.Flo},
  defStats: []stats.Stat{stats.Flo},
}

var Spirit = &Class{
  name: "Spirit",
  growth: StatGrowth{
    Det: 24,
    Stm: 6,
    Foc: 8,
    Sag: 9,
  },
  atkStats: []stats.Stat{stats.Sag},
  defStats: []stats.Stat{stats.Sag},
}

var Ghost = &Class{
  name: "Ghost",
  growth: StatGrowth{
    Det: 22,
    Stm: 5,
    Foc: 7,
    Sag: 8,
  },
  atkStats: []stats.Stat{stats.Sag},
  defStats: []stats.Stat{stats.Sag},
}

var Undead = &Class{
  name: "Undead",
  growth: StatGrowth{
    Det: 28,
    Stm: 8,
    Foc: 7,
    Sag: 10,
  },
  atkStats: []stats.Stat{stats.Sag},
  defStats: []stats.Stat{stats.Sag},
}

// Cultists & Shrine

var Cultist = &Class{
  name: "Cultist",
  growth: StatGrowth{
    Det: 22,
    Stm: 6,
    Foc: 6,
    Ing: 8,
  },
  atkStats: []stats.Stat{stats.Ing},
  defStats: []stats.Stat{stats.Ing},
}

var CultistAcolyte = &Class{
  name: "Cultist Acolyte",
  growth: StatGrowth{
    Det: 22,
    Stm: 6,
    Foc: 6,
    Ing: 8,
  },
  atkStats: []stats.Stat{stats.Ing},
  defStats: []stats.Stat{stats.Ing},
}

var CultistZealot = &Class{
  name: "Cultist Zealot",
  growth: StatGrowth{
    Det: 26,
    Stm: 7,
    Foc: 6,
    Ing: 9,
  },
  atkStats: []stats.Stat{stats.Ing},
  defStats: []stats.Stat{stats.Ing},
}

var CultistWarbandScout = &Class{
  name: "Cultist Warband Scout",
  growth: StatGrowth{
    Det: 26,
    Stm: 7,
    Foc: 7,
    Ing: 9,
  },
  atkStats: []stats.Stat{stats.Ing},
  defStats: []stats.Stat{stats.Ing},
}

var CultistHighPriest = &Class{
  name: "Cultist High Priest",
  growth: StatGrowth{
    Det: 34,
    Stm: 10,
    Foc: 12,
    Ing: 13,
  },
  atkStats: []stats.Stat{stats.Ing},
  defStats: []stats.Stat{stats.Ing},
}

var RogueCultist = &Class{
  name: "Rogue Cultist",
  growth: StatGrowth{
    Det: 26,
    Stm: 7,
    Foc: 6,
    Ing: 9,
  },
  atkStats: []stats.Stat{stats.Ing},
  defStats: []stats.Stat{stats.Ing},
}

var ShrineAcolyte = &Class{
  name: "Shrine Acolyte",
  growth: StatGrowth{
    Det: 22,
    Stm: 6,
    Foc: 7,
    Sag: 8,
  },
  atkStats: []stats.Stat{stats.Sag},
  defStats: []stats.Stat{stats.Sag},
}

var ShrineDevotee = &Class{
  name: "Shrine Devotee",
  growth: StatGrowth{
    Det: 22,
    Stm: 6,
    Foc: 7,
    Sag: 8,
  },
  atkStats: []stats.Stat{stats.Sag},
  defStats: []stats.Stat{stats.Sag},
}

var ShrineKeeper = &Class{
  name: "Shrine Keeper",
  growth: StatGrowth{
    Det: 26,
    Stm: 7,
    Foc: 8,
    Sag: 9,
  },
  atkStats: []stats.Stat{stats.Sag},
  defStats: []stats.Stat{stats.Sag},
}

// Seneca Downs townsfolk

var Blacksmith = &Class{
  name: "Blacksmith",
  growth: StatGrowth{
    Det: 20,
    Stm: 7,
    Foc: 4,
    Str: 7,
  },
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Str},
}

var Guard = &Class{
  name: "Guard",
  growth: StatGrowth{
    Det: 18,
    Stm: 6,
    Foc: 4,
    Str: 6,
  },
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Str},
}

var GuardCaptain = &Class{
  name: "Guard Captain",
  growth: StatGrowth{
    Det: 24,
    Stm: 8,
    Foc: 5,
    Str: 9,
  },
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Str},
}

var Innkeeper = &Class{
  name: "Innkeeper",
  growth: StatGrowth{
    Det: 14,
    Stm: 4,
    Foc: 4,
    Sag: 4,
  },
  atkStats: []stats.Stat{stats.Sag},
  defStats: []stats.Stat{stats.Sag},
}

var Merchant = &Class{
  name: "Merchant",
  growth: StatGrowth{
    Det: 14,
    Stm: 3,
    Foc: 5,
    Ing: 4,
  },
  atkStats: []stats.Stat{stats.Ing},
  defStats: []stats.Stat{stats.Ing},
}

var Peddler = &Class{
  name: "Peddler",
  growth: StatGrowth{
    Det: 16,
    Stm: 4,
    Foc: 5,
    Ing: 5,
  },
  atkStats: []stats.Stat{stats.Ing},
  defStats: []stats.Stat{stats.Ing},
}

var Shepherd = &Class{
  name: "Shepherd",
  growth: StatGrowth{
    Det: 18,
    Stm: 5,
    Foc: 5,
    Sag: 6,
  },
  atkStats: []stats.Stat{stats.Sag},
  defStats: []stats.Stat{stats.Sag},
}

// Aggressive/elite variants -- distinct from their base animal classes so
// upgrading one encounter doesn't inadvertently buff every trivial NPC
// sharing that base class.

var BirdSwarm = &Class{
  name: "BirdSwarm",
  growth: StatGrowth{
    Det: 24,
    Stm: 6,
    Foc: 6,
    Flo: 9,
  },
  atkStats: []stats.Stat{stats.Flo},
  defStats: []stats.Stat{stats.Flo},
}

var Badger = &Class{
  name: "Badger",
  growth: StatGrowth{
    Det: 26,
    Stm: 9,
    Foc: 4,
    Str: 9,
  },
  atkStats: []stats.Stat{stats.Str},
  defStats: []stats.Stat{stats.Str},
}
