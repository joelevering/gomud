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
  pctDmg = EffectType("pctDmg") // % Damage
  flatDmg = EffectType("flatDmg") // Flat Damage
  oppFx = EffectType("oppFx") // Status Effect on Opponent
)

var(
  Bash = &Skill{
    Name: "bash",
    CostType: stats.Stm,
    CostAmt: 10,
    Effects: []Effect{
      Effect{
        Type: pctDmg,
        Value: 1.25,
      },
      Effect{
        Type: oppFx,
        Value: statfx.SEInst{
          Effect: statfx.Stun,
          Chance: 0.25,
        },
      },
    },
  }
)
