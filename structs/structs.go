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
  Stunned bool
  LoweredAtk bool
  LoweredDef bool
  Surprised SurpriseRep
}

type SurpriseRep struct {
  Stunned bool
  LowerAtk bool
  LowerDef bool
}
