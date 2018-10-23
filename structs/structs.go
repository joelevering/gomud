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
  SkName string // skill name
  Stunned bool // TODO test and add report logic
}
