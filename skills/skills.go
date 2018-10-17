package skills

import (
  "github.com/joelevering/gomud/statfx"
  "github.com/joelevering/gomud/stats"
)

type Skill struct {
  Name     string
  Effects  []Effect
  CostType stats.Stat
  CostAmt  int
}

type Effect struct {
  Type  EffectType
  Value interface{}
}

type EffectType string

const(
  PctDmg = EffectType("pctDmg") // % Damage
  FlatDmg = EffectType("flatDmg") // Flat Damage
  PctHeal = EffectType("pctHeal") // % Healing
  OppFx = EffectType("oppFx") // Status Effect on Opponent
)

var(
  Stun = &Skill{
    Name: "stun",
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
  Bash = &Skill{
    Name: "bash",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: PctDmg,
        Value: 1.25,
      },
      Effect{
        Type: OppFx,
        Value: statfx.SEInst{
          Effect: statfx.Stun,
          Chance: 0.25,
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
)
