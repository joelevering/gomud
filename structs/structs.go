package structs

import (
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
)

type CmbFx struct {
  Dmg     int
  SelfDmg int
  DotDmgs []statfx.DotInst
  Heal    int
  StmRec  int
  SFx     []statfx.SEInst
  SelfSFx []statfx.SEInst
  Dots    []statfx.DotInst
  Skill   skills.Skill
  Req     statfx.StatusEffect
}

type CmbRep struct {
  CmbFx
  Missed bool
  Stunned bool
  Empowered bool
  Steeled bool
  Weak bool
  Vulnerable bool
  Concentrating bool
  Dodged bool
  Redirected bool
  Surprised SurpriseRep

  FollowUpReq statfx.StatusEffect
}

type SurpriseRep struct {
  Stunned bool
  Weak bool
  Vulnerable bool
}
