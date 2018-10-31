package skills

import (
  "github.com/joelevering/gomud/statfx"
  "github.com/joelevering/gomud/stats"
)

type Skill struct {
  Name        string
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
  PctHeal = EffectType("pctHeal") // % Healing
  OppFx = EffectType("oppFx") // Status Effect on Opponent
  SelfFx = EffectType("selfFx") // Status Effect on Self
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

func (s *Skill) IsFollowUp() bool {
  return s.FollowUpReq != nil
}

func (s *Skill) SelfFollowUpReq() (bool, statfx.StatusEffect) {
  if s.FollowUpReq.Type == SelfFollowUp {
    return true, s.FollowUpReq.Effect
  }

  return false, ""
}

func (s *Skill) OppFollowUpReq() (bool, statfx.StatusEffect) {
  if s.FollowUpReq.Type == OppFollowUp {
    return true, s.FollowUpReq.Effect
  }

  return false, ""
}

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
  Shove = &Skill{
    Name: "shove",
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
        },
      },
    },
  }
  PowerNap = &Skill{
    Name: "power nap",
    CostType: stats.Stm,
    CostAmt: 0,
    Effects: []Effect{
      Effect{
        Type: PctHeal,
        Value: 0.1,
      },
    },
  }
  Charge = &Skill{
    Name: "charge",
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
    Name: "conserve",
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
    Name: "desperate blow",
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
    Name: "frenetic pace",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Surprise,
          Chance: 0.6,
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
    Name: "back up",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Surprise,
          Chance: 0.6,
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
    Name: "duck",
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
    Name: "counter",
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
)
