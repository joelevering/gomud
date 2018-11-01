package structs

import (
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
)

type CmbFx struct {
  Dmg     int
  Heal    int
  SFx     []statfx.SEInst
  SelfSFx []statfx.SEInst
  Skill   skills.Skill
  Req     statfx.StatusEffect
}

type CmbRep struct {
  CmbFx
  Missed bool
  Stunned bool
  LowerAtk bool
  LowerDef bool
  Concentrating bool
  Dodged bool
  Surprised SurpriseRep

  FollowUpReq statfx.StatusEffect
}

type SurpriseRep struct {
  Stunned bool
  LowerAtk bool
  LowerDef bool
}
