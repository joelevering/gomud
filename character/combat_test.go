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

func Test_AtkFxWithConcentration(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  conInst := statfx.SEInst{
    Effect: statfx.Concentration,
    Chance: 1,
  }
  ch.addFx(conInst)
  ch.SetCmbSkill(skills.T_Stun)

  cFx := ch.AtkFx(rep)

  if len(cFx.SFx) != 0 {
    t.Errorf("Concentrating should not result in statfx combat fx, but it did (%v)", cFx.SFx)
  }

  if cFx.Dmg != ch.GetAtk() {
    t.Error("CFx while concentrating should have regular damage, but they don't")
  }

  if rep.Skill.Name != skills.T_Stun.Name {
    t.Error("Skill should still be reported when concetrating, but it wasn't")
  }

  if !rep.Concentrating {
    t.Error("Rep should report concentration, but it's not")
  }
}

func Test_AtkFxPctDmg(t *testing.T){
  ch := NewCharacter()
  pctDmgCh := NewCharacter()
  pctDmgCh.SetCmbSkill(skills.T_PctDmg)
  rep := &structs.CmbRep{}

  cFx := ch.AtkFx(rep)
  pctDmgCFx := pctDmgCh.AtkFx(rep)

  if pctDmgCFx.Dmg != (cFx.Dmg/2) {
    t.Errorf("Expected .5 pct dmg skill to deal half damage, but it dealt %d compared to %d", pctDmgCFx.Dmg, cFx.Dmg)
  }
}

func Test_AtkFxFlatDmg(t *testing.T){
  ch := NewCharacter()
  flatDmgCh := NewCharacter()
  flatDmgCh.SetCmbSkill(skills.T_FlatDmg)
  rep := &structs.CmbRep{}

  cFx := ch.AtkFx(rep)
  flatDmgCFx := flatDmgCh.AtkFx(rep)

  if flatDmgCFx.Dmg != (cFx.Dmg + 10) {
    t.Errorf("Expected 10 flat damage skill to deal dmg + 10, but it dealt %d compared to %d", flatDmgCFx.Dmg, cFx.Dmg)
  }
}

func Test_AtkFxDmgChance(t *testing.T) {
  ch := NewCharacter()
  ch.SetCmbSkill(skills.T_NoChance)
  rep := &structs.CmbRep{}

  cFx := ch.AtkFx(rep)

  if cFx.Dmg != 0 {
    t.Errorf("Expected no chance dmg skill to generate 0 dmg, but it generated %d", cFx.Dmg)
  }
}

func Test_AtkFxIsImpactedByWeak(t *testing.T) {
  ch := NewCharacter()
  chWeak := NewCharacter()
  weakInst := statfx.SEInst{
    Effect: statfx.Weak,
    Chance: 1,
    Duration: 1,
  }
  chWeak.addFx(weakInst)
  rep := &structs.CmbRep{}

  fx := ch.AtkFx(rep)
  weakFx := chWeak.AtkFx(rep)

  if weakFx.Dmg >= fx.Dmg {
    t.Errorf("Expected LowerAtk to result in a less damaging atk than usual, but was %d compared to %d", weakFx.Dmg, fx.Dmg)
  }
}

func Test_StunnedCharsDoNotAttack(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  ch.addFx(skills.T_Stun.Effects[0].Value.(statfx.SEInst))

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
  ch.CmbSkill = skills.T_Stun
  ch.addFx(skills.T_Stun.Effects[0].Value.(statfx.SEInst))

  cFx := ch.AtkFx(rep)

  if len(cFx.SFx) != 0 {
    t.Errorf("Expected stunned character to return cmb fx with no sfx, but got %v", cFx.SFx)
  }
}

func Test_AtkFxReducesResourcesWhenUsingSkill(t *testing.T) {
  ch := NewCharacter()
  ch.CmbSkill = skills.T_Stun
  rep := &structs.CmbRep{}

  ch.AtkFx(rep)

  if ch.GetStm() != (ch.GetMaxStm() - skills.T_Stun.CostAmt) {
    t.Errorf("Expected generating fx for an attack that uses Stun skill to reduce char stamina by 10, but it was %d/%d", ch.GetStm(), ch.GetMaxStm())
  }

  if rep.Skill.Name != skills.T_Stun.Name {
    t.Errorf("Expected generating fx an attack using Stun skill to report Stun as skill used, but report has %s for skill name", rep.Skill.Name)
  }
}

func Test_AtkFxIsImpactedByConserveUse(t *testing.T) {
  ch := NewCharacter()
  ch.addFx(skills.Conserve.Effects[0].Value.(statfx.SEInst))
  ch.CmbSkill = skills.T_Stun
  rep := &structs.CmbRep{}

  ch.AtkFx(rep)
  stmUsed := ch.GetMaxStm() - ch.GetStm()

  if stmUsed == skills.T_Stun.CostAmt {
    t.Error("Expected Conserve to reduce stamina cost of skill, but it didn't")
  }
}

func Test_AtkFxDoesNotUseSkillWhenLackingResources(t *testing.T) {
  ch := NewCharacter()
  ch.SetStm(0)
  ch.CmbSkill = skills.T_Stun
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
  vulnCh := NewCharacter()
  vulnInst := statfx.SEInst{
    Effect: statfx.Vulnerable,
    Chance: 1,
    Duration: 1,
  }
  vulnCh.addFx(vulnInst)
  rep := &structs.CmbRep{}
  fx := structs.CmbFx{Dmg: 100}

  regRes := ch.ResistAtk(fx, rep)
  vulnRes := vulnCh.ResistAtk(fx, rep)

  if vulnRes.Dmg <= regRes.Dmg {
    t.Errorf("Expected Vulnerable to result in a more damaging atk than usual, but was %d compared to %d", vulnRes.Dmg, regRes.Dmg)
  }
}

func Test_ResistAtkKeepsStatusEffects(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  fx := structs.CmbFx{
    SFx: []statfx.SEInst{
      statfx.SEInst{
        Effect: statfx.Stun,
      },
    },
  }

  res := ch.ResistAtk(fx, rep)

  if len(res.SFx) != 1 || res.SFx[0].Effect != statfx.Stun {
    t.Errorf("Expected status effects to remain the same on resist, but got %v", res.SFx)
  }
}

func Test_ResistAtkKeepsSelfStatusEffects(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  fx := structs.CmbFx{
    SelfSFx: []statfx.SEInst{
      statfx.SEInst{
        Effect: statfx.Conserve,
      },
    },
  }

  res := ch.ResistAtk(fx, rep)

  if len(res.SelfSFx) != 1 || res.SelfSFx[0].Effect != statfx.Conserve {
    t.Errorf("Expected self status effects to remain the same on resist, but got %v", res.SelfSFx)
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
  ch.SetCmbSkill(skills.T_Stun)

  cFx := ch.AtkFx(rep)
  ch.ApplyDef(cFx, rep)

  if !ch.isStunned() {
    t.Error("Expected character to be stunned when using ApplyDef with stun SE, but char was not")
  }

  if len(rep.SFx) != 1 || rep.SFx[0].Effect != statfx.Stun {
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
  ch.CmbSkill = skills.PowerNap
  ch.SetDet(1)

  cFx := ch.AtkFx(rep)
  ch.ApplyAtk(cFx, rep)

  if ch.GetDet() != 21 {
    t.Errorf("Expected ApplyAtk with healing to increase health from 1 to 21, but it's %d", ch.GetDet())
  }

  if rep.Heal != 20 {
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

func Test_ApplyAtkAppliesSelfFx(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  ch.SetCmbSkill(skills.Conserve)

  cFx := ch.AtkFx(rep)
  ch.ApplyAtk(cFx, rep)

  if !ch.isConserving() {
    t.Error("Expected atkfx + applyAtk with Conserve skill to apply Conserve to attacker, but it didn't")
  }
}

func Test_ApplyAtkDoesNotApplyDmgOrStatfx(t *testing.T) {
  ch := NewCharacter()
  rep := &structs.CmbRep{}
  ch.SetCmbSkill(skills.T_Stun)

  cFx := ch.AtkFx(rep)
  cFx.Dmg = 10
  ch.ApplyAtk(cFx, rep)

  if ch.isStunned() {
    t.Error("Expected character to not be stunned with ApplyAtk with stun, but they were")
  }

  if ch.GetDet() < ch.GetMaxDet() {
    t.Errorf("Expected character to not lose health with ApplyAtk, but they're at %d/%d", ch.GetDet(), ch.GetMaxDet())
  }
}
