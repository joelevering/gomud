package skills

import (
  "strings"

  "github.com/joelevering/gomud/statfx"
  "github.com/joelevering/gomud/stats"
)

var All = []*Skill{
  Shove,
  PowerNap,
  Charge,
  Conserve,
  DesperateBlow,
  FreneticPace,
  BackUp,
  Duck,
  Counter,
  Uppercut,
  WittyRetort,
  Ploy,
  Sideswipe,
  Plan,
  Sidestep,
  LowBlow,
  Concentrate,
  FirstAid,
  TargetedStrike,
  Spark,
  CastDoubt,
  Benumb,
  Radiate,
  Redirect,
  CastOff,

  Pince,
  Hide,
}

var ByName map[string]*Skill

func GetSkill(name string) *Skill {
  if len(ByName) == 0 {
    ByName = make(map[string]*Skill)
    for _, sk := range All {
      ByName[sk.Name] = sk
    }
  }

  return ByName[strings.Title(name)]
}

type Skill struct {
  Name        string
  Desc        string
  Effects     []Effect
  CostType    stats.Stat
  CostAmt     int
  Rstcn       Rstcn
  FollowUpReq *FollowUpReq
}

type Effect struct {
  Type   EffectType
  Value  interface{}
  Chance float64
}

type EffectType string

const(
  PctDmg = EffectType("pctDmg") // % Damage
  FlatDmg = EffectType("flatDmg") // Flat Damage
  PctHeal = EffectType("pctHeal") // % determination recovery
  FlatStm = EffectType("flatStm") // Flat stamina recovery
  OppFx = EffectType("oppFx") // Status Effect on Opponent
  SelfFx = EffectType("selfFx") // Status Effect on Self
  Dot = EffectType("doT") // Damage over Time
)

type Rstcn string // Restriction

const(
  OOCOnly = Rstcn("OOCOnly") // skill can only be used Out Of Combat
)

type FollowUpReq struct {
  Type   string
  Effect statfx.StatusEffect
}

const(
  SelfFollowUp = "selfFollowUp"
  OppFollowUp  = "oppFollowUp"
)

func (s *Skill) IsOOCOnly() bool {
  return s.Rstcn == OOCOnly
}

func (s *Skill) SelfFollowUpReq() (bool, statfx.StatusEffect) {
  if s.FollowUpReq != nil && s.FollowUpReq.Type == SelfFollowUp {
    return true, s.FollowUpReq.Effect
  }

  return false, ""
}

func (s *Skill) OppFollowUpReq() (bool, statfx.StatusEffect) {
  if s.FollowUpReq != nil && s.FollowUpReq.Type == OppFollowUp {
    return true, s.FollowUpReq.Effect
  }

  return false, ""
}

// Test Skills
var(
  T_PctDmg = &Skill{
    Name: "t_pctdmg",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: PctDmg,
        Value: 0.5,
      },
    },
  }
  T_FlatDmg = &Skill{
    Name: "t_flatdmg",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: FlatDmg,
        Value: 10,
      },
    },
  }
  T_Bleed = &Skill{
    Name: "t_bleed",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: Dot,
        Value: statfx.DotFx{
          Type: statfx.Bleed,
          Chance: 1,
          DmgPct: 0.5,
          DurationMin: 1,
          DurationMax: 1,
        },
      },
    },
  }
  T_NoChance = &Skill{
    Name: "t_nochance",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: FlatDmg,
        Value: 10,
        Chance: 0.00001,
      },
    },
  }
  T_Stun = &Skill{
    Name: "t_stun",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Stun,
          Chance: 1,
        },
      },
    },
  }
  T_PctHeal = &Skill{
    Name: "t_pctHeal",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: PctHeal,
        Value: 0.1,
      },
    },
  }
  T_FlatStm = &Skill{
    Name: "t_flatStm",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: FlatStm,
        Value: 10,
      },
    },
  }
)

// Player Skills
var(
  Shove = &Skill{
    Name: "Shove",
    Desc: "Pushes an enemy back for partial damage, with a chance to surprise them.",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: PctDmg,
        Value: 0.75,
      },
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Surprise,
          Chance: 0.5,
          Duration: 1,
        },
      },
    },
  }
  PowerNap = &Skill{
    Name: "Power Nap",
    Desc: "Fall asleep on your feet for a turn, healing a small percentage of health.",
    CostType: stats.Stm,
    CostAmt: 0,
    Effects: []Effect{
      Effect{
        Type: PctHeal,
        Value: 0.20,
      },
    },
  }
  Charge = &Skill{
    Name: "Charge",
    Desc: "Close the gap to an enemy, dealing double damage.",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: PctDmg,
        Value: 2.0,
      },
    },
    Rstcn: OOCOnly,
  }
  Conserve = &Skill{
    Name: "Conserve",
    Desc: "Save strength, reducing stamina consumption on future turns",
    CostType: stats.Stm,
    CostAmt: 0,
    Effects: []Effect{
      Effect{
        Type: SelfFx,
        Value: statfx.SEInst{
          Effect: statfx.Conserve,
          Chance: 1,
          Duration: 2,
        },
      },
    },
  }
  DesperateBlow = &Skill{
    Name: "Desperate Blow",
    Desc: "An easily read attack that will do a LOT of damage if it hits",
    CostType: stats.Stm,
    CostAmt: 20,
    Effects: []Effect{
      Effect{
        Type: PctDmg,
        Value: 3.0,
        Chance: 0.34,
      },
    },
  }
  FreneticPace = &Skill{
    Name: "Frenetic Pace",
    Desc: "Focus on attacking the enemy this turn and the next one for a chance to surprise.",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Surprise,
          Chance: 0.6,
          Duration: 1,
        },
      },
      Effect{
        Type: SelfFx,
        Value: statfx.SEInst{
          Effect: statfx.Concentration,
          Chance: 1,
          Duration: 1,
        },
      },
      Effect{
        Type: PctDmg,
        Value: 1.0,
      },
    },
  }
  BackUp = &Skill{
    Name: "Back Up",
    Desc: "Take a possibly-surprising step back to weaken the effectiveness of your enemy's next attack.",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Surprise,
          Chance: 0.6,
          Duration: 1,
        },
      },
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Weak,
          Chance: 1,
        },
      },
    },
  }
  Duck = &Skill{
    Name: "Duck",
    Desc: "Attempt to dodge the enemy's next attack.",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: SelfFx,
        Value: statfx.SEInst{
          Effect: statfx.Dodging,
          Chance: 0.85,
          Duration: 1,
        },
      },
    },
  }
  Counter = &Skill{
    Name: "Counter",
    Desc: "A powerful counter strike that must follow a dodge.",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: PctDmg,
        Value: 2.2,
      },
    },
    FollowUpReq: &FollowUpReq{
      Type: SelfFollowUp,
      Effect: statfx.Dodging,
    },
  }
  Uppercut = &Skill{
    Name: "Uppercut",
    Desc: "A powerful hit that must follow a surprise attack.",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: PctDmg,
        Value: 2.0,
      },
    },
    FollowUpReq: &FollowUpReq{
      Type: OppFollowUp,
      Effect: statfx.Surprise,
    },
  }
  WittyRetort = &Skill{
    Name: "Witty Retort",
    Desc: "Mental warfare that could make your enemy vulnerable and/or weak for one turn.",
    CostType: stats.Foc,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Vulnerable,
          Chance: 0.66,
          Duration: 1,
        },
      },
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Weak,
          Chance: 0.66,
          Duration: 1,
        },
      },
    },
  }
  Ploy = &Skill{
    Name: "Ploy",
    Desc: "A clever-but-weak attack that has a chance to surprise the enemy",
    CostType: stats.Foc,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Surprise,
          Chance: 0.66,
          Duration: 1,
        },
      },
      Effect{
        Type: PctDmg,
        Value: 0.33,
      },
    },
  }
  Sideswipe = &Skill{
    Name: "Sideswipe",
    Desc: "A weaker attack with a decent chance to bleed the enemy",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: Dot,
        Value: statfx.DotFx{
          Type: statfx.Bleed,
          Chance: 0.75,
          DmgPct: 0.33,
          DurationMin: 2,
          DurationMax: 2,
        },
      },
      Effect{
        Type: PctDmg,
        Value: 0.5,
      },
    },
  }
  Plan = &Skill{
    Name: "Plan",
    Desc: "Spend the turn planning your next move, empowering future attacks and discovering enemy vulnerabilites",
    CostType: stats.Foc,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: SelfFx,
        Value: statfx.SEInst{
          Effect: statfx.Vulnerable,
          Chance: 1,
          Duration: 1,
        },
      },
      Effect{
        Type: SelfFx,
        Value: statfx.SEInst{
          Effect: statfx.Empowered,
          Chance: 1,
          Duration: 2,
        },
      },
    },
  }
  Sidestep = &Skill{
    Name: "Sidestep",
    Desc: "Attempt to dodge your enemy's next attack",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: SelfFx,
        Value: statfx.SEInst{
          Effect: statfx.Dodging,
          Chance: 0.85,
          Duration: 1,
        },
      },
    },
  }
  LowBlow = &Skill{
    Name: "Low Blow",
    Desc: "An uncouth strike with a chance to stun your enemy...or enrage them",
    CostType: stats.Foc,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Stun,
          Chance: 0.5,
          Duration: 1,
        },
      },
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Empowered,
          Chance: 0.5,
          Duration: 1,
        },
      },
    },
  }
  Concentrate = &Skill{
    Name: "Concentrate",
    Desc: "Focus on empowering future attacks",
    CostType: stats.Foc,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: SelfFx,
        Value: statfx.SEInst{
          Effect: statfx.Empowered,
          Chance: 1.0,
          Duration: 2,
        },
      },
    },
  }
  FirstAid = &Skill{
    Name: "First Aid",
    Desc: "Heal a moderate amount of health",
    CostType: stats.Foc,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: PctHeal,
        Value: 0.35,
      },
    },
  }
  TargetedStrike = &Skill{
    Name: "Targeted Strike",
    Desc: "A regular attack with a good chance of making the enemy vulnerable to your follow-up",
    CostType: stats.Foc,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Vulnerable,
          Chance: 0.75,
          Duration: 1,
        },
      },
      Effect{
        Type: PctDmg,
        Value: 1.0,
      },
    },
  }
  Spark = &Skill{
    Name: "Spark",
    Desc: "Strike an enemy with a firey burst",
    CostType: stats.Foc,
    CostAmt: 15,
    Effects: []Effect{
      Effect{
        Type: PctDmg,
        Value: 2.0,
      },
    },
  }
  CastDoubt = &Skill{
    Name: "Cast Doubt",
    Desc: "A biting word that can preoccupy an enemy, leaving them vulnerable for longer than you'd expect",
    CostType: stats.Foc,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Vulnerable,
          Chance: 0.5,
          Duration: 3,
        },
      },
    },
  }
  Benumb = &Skill{
    Name: "Benumb",
    Desc: "A chilly wind weakens the enemy's next attack.",
    CostType: stats.Foc,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Weak,
          Chance: 1.0,
          Duration: 1,
        },
      },
    },
  }
  Radiate = &Skill{
    Name: "Radiate",
    Desc: "A deceptive attack that's lack of immediate effect is proceeded by multiple rounds of fire damage",
    CostType: stats.Foc,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: Dot,
        Value: statfx.DotFx{
          Type: statfx.Fire,
          Chance: 1,
          DmgPct: 0.3,
          DurationMin: 2,
          DurationMax: 4,
        },
      },
    },
  }
  Redirect = &Skill{
    Name: "Redirect",
    Desc: "Reduce the impact of your enemy's next attack while causing the enemy to deal partial damage to themself",
    CostType: stats.Foc,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: SelfFx,
        Value: statfx.SEInst{
          Effect: statfx.Redirecting,
          Chance: 1.0,
          Duration: 1.0,
        },
      },
    },
  }
  CastOff = &Skill{
    Name: "Cast Off",
    Desc: "Shrug off physical fatigue, channeling focus to replenish stamina",
    CostType: stats.Foc,
    CostAmt: 20,
    Effects: []Effect{
      Effect{
        Type: FlatStm,
        Value: 20,
      },
    },
  }
)

// NPC Skills
var(
  Pince = &Skill{
    Name: "Pince",
    Desc: "Snip a foe for boosted damage",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: PctDmg,
        Value: 1.5,
      },
    },
  }
  Hide = &Skill{
    Name: "Hide",
    Desc: "Retreat inward, steeling oneself defensively",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: SelfFx,
        Value: statfx.SEInst{
          Effect: statfx.Steeled,
          Chance: 1,
          Duration: 1,
        },
      },
    },
  }
)
