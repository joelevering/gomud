package structs

import (
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
)

type CmbFx struct {
  Dmg int
  Heal int
  SFx []statfx.StatusEffect
  SelfSFx []statfx.StatusEffect
  Skill skills.Skill
}

type CmbRep struct {
  CmbFx
  Stunned bool
  LowerAtk bool
  LowerDef bool
  Surprised SurpriseRep
}

type SurpriseRep struct {
  Stunned bool
  LowerAtk bool
  LowerDef bool
}
