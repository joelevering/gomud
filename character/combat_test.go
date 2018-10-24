package character

import (
  "testing"

  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
  "github.com/joelevering/gomud/structs"
)

func Test_CharAttacksByDefault(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}

  cFx := ch.AtkFx(rep)

  if cFx.Dmg != ch.GetAtk() {
    t.Errorf("Expected default combat action to apply damage equal to attack, but dmg was %d (and atk is %d)", cFx.Dmg, ch.GetAtk())
  }
}

func Test_AtkFxIsImpactedByLowerAtk(t *testing.T) {
  ch := NewCharacter()
  chLowAtk := NewCharacter()
  chLowAtk.LowerAtk = true
  rep := &structs.CmbRep{}

  fx := ch.AtkFx(rep)
  lowAtkFx := chLowAtk.AtkFx(rep)

  if lowAtkFx.Dmg >= fx.Dmg {
    t.Errorf("Expected LowerAtk to result in a less damaging atk than usual, but was %d compared to %d", lowAtkFx.Dmg, fx.Dmg)
  }
}

func Test_StunnedCharsDoNotAttack(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  ch.Stunned = true

  cFx := ch.AtkFx(rep)

  if cFx.Dmg != 0 {
    t.Errorf("Expected stunned character to return cmb fx with no dmg, but got %d", cFx.Dmg)
  }

  if !rep.Stunned {
    t.Error("Expected report from AtkFx of stunned character to be stunned: true, but it wasn't")
  }
}

func Test_StunnedCharsDoNotUseSkills(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  ch.CmbSkill = skills.Stun
  ch.Stunned = true

  cFx := ch.AtkFx(rep)

  if len(cFx.SFx) != 0 {
    t.Errorf("Expected stunned character to return cmb fx with no sfx, but got %v", cFx.SFx)
  }
}

func Test_AtkFxReducesResourcesWhenUsingSkill(t *testing.T) {
  ch := NewCharacter()
  ch.CmbSkill = skills.Stun
  rep := &structs.CmbRep{}

  ch.AtkFx(rep)

  if ch.GetStm() != (ch.GetMaxStm() - skills.Stun.CostAmt) {
    t.Errorf("Expected generating fx for an attack that uses Stun skill to reduce char stamina by 10, but it was %d/%d", ch.GetStm(), ch.GetMaxStm())
  }

  if rep.Skill.Name != skills.Stun.Name {
    t.Errorf("Expected generating fx an attack using Stun skill to report Stun as skill used, but report has %s for skill name", rep.Skill.Name)
  }
}

func Test_AtkFxDoesNotUseSkillWhenLackingResources(t *testing.T) {
  ch := NewCharacter()
  ch.SetStm(0)
  ch.CmbSkill = skills.Stun
  rep := &structs.CmbRep{}

  fx := ch.AtkFx(rep)

  if len(fx.SFx) != 0 {
    t.Errorf("Expected AtkFx to not return statfx for using skill with no resources, but SFx was %v", fx.SFx)
  }

  if rep.Skill.Name != "" {
    t.Errorf("Expected AtkFx to not report a Skill used when no resource to pay for it, but reported %s", rep.Skill.Name)
  }
}

func Test_ResistAtkLowersDmg(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  fx := structs.CmbFx{Dmg: 100}

  res := ch.ResistAtk(fx, rep)

  if res.Dmg >= 100 {
    t.Errorf("Expected applying an attack to report lowered damage, but reported %d -> %d", fx.Dmg, res.Dmg)
  }
}

func Test_ResistAtkIsImpactedByLowerDef(t *testing.T) {
  ch := NewCharacter()
  chLowDef := NewCharacter()
  chLowDef.LowerDef = true
  rep := &structs.CmbRep{}
  fx := structs.CmbFx{Dmg: 100}

  regRes := ch.ResistAtk(fx, rep)
  lowDefRes := chLowDef.ResistAtk(fx, rep)

  if lowDefRes.Dmg <= regRes.Dmg {
    t.Errorf("Expected LowerDef to result in a more damaging atk than usual, but was %d compared to %d", lowDefRes.Dmg, regRes.Dmg)
  }
}

func Test_ResistAtkKeepsStatusEffects(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  fx := structs.CmbFx{
    SFx: []statfx.StatusEffect{
      statfx.Stun,
    },
  }

  res := ch.ResistAtk(fx, rep)

  if len(res.SFx) != 1 || res.SFx[0] != statfx.Stun {
    t.Errorf("Expected status effects to remain the same on resist, but got %v", res.SFx)
  }
}

func Test_ApplyDefDealsDamage(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  cFx := structs.CmbFx{Dmg: 10}

  ch.ApplyDef(cFx, rep)

  if ch.GetDet() != ch.GetMaxDet() - 10 {
    t.Errorf("Expected ApplyDef with 10 dmg to reduce health by 10, but got %d/%d", ch.GetDet(), ch.GetMaxDet())
  }

  if rep.Dmg != 10 {
    t.Errorf("Expected ApplyDef to apply dmg to report, but rep dmg is %d", rep.Dmg)
  }
}

func Test_ApplyDefAppliesStatfx(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  ch.SetCmbSkill(skills.Stun)

  cFx := ch.AtkFx(rep)
  ch.ApplyDef(cFx, rep)

  if !ch.Stunned {
    t.Error("Expected character to be stunned when using ApplyDef with stun SE, but char was not")
  }

  if len(rep.SFx) != 1 || rep.SFx[0] != statfx.Stun {
    t.Errorf("Expected ApplyDef to apply sfx to report, but report sfx are %v", rep.SFx)
  }
}

func Test_ApplyDefDoesNotHeal(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  ch.SetDet(1)
  cFx := structs.CmbFx{Heal: 10}

  ch.ApplyDef(cFx, rep)

  if ch.GetDet() != 1 {
    t.Error("Expected no healing with ApplyDef, but the char healed")
  }
}

func Test_ApplyAtkHeals(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  ch.SetDet(1)
  cFx := structs.CmbFx{Heal: 10}

  ch.ApplyAtk(cFx, rep)

  if ch.GetDet() != 11 {
    t.Errorf("Expected ApplyAtk with 10 healing to increase health from 1 to 11, but it's %d", ch.GetDet())
  }

  if rep.Heal != 10 {
    t.Errorf("Expected ApplyAtk to apply healing to report, but rep healing is %d", rep.Heal)
  }
}

func Test_ApplyAtkHealingReportsAccurately(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  ch.SetDet(ch.GetMaxDet() - 1)
  cFx := structs.CmbFx{Heal: 10}

  ch.ApplyAtk(cFx, rep)

  if ch.GetDet() != ch.GetMaxDet() {
    t.Errorf("Expected ApplyAtk with healing to max out health but got %d/%d", ch.GetDet(), ch.GetMaxDet())
  }

  if rep.Heal != 1 {
    t.Errorf("Expected ApplyAtk to apply accurate healing when maxing health, but rep healing is %d", rep.Heal)
  }
}

func Test_ApplyAtkDoesNotApplyDmgOrStatfx(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  ch.SetCmbSkill(skills.Stun)

  cFx := ch.AtkFx(rep)
  cFx.Dmg = 10
  ch.ApplyAtk(cFx, rep)

  if ch.Stunned {
    t.Error("Expected character to not be stunned with ApplyAtk with stun, but they were")
  }

  if ch.GetDet() < ch.GetMaxDet() {
    t.Errorf("Expected character to not lose health with ApplyAtk, but they're at %d/%d", ch.GetDet(), ch.GetMaxDet())
  }
}
