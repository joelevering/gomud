package structs

import "github.com/joelevering/gomud/statfx"

type CmbFx struct {
  Dmg int
  Heal int
  SFx []statfx.StatusEffect
  SelfSFx []statfx.StatusEffect
}

type CmbRep struct {
  CmbFx
  Stunned bool // TODO test and add report logic
}
