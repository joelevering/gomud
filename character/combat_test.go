package character

import (
  "testing"

  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
  "github.com/joelevering/gomud/structs"
)

func Test_CharAttacksByDefault(t *testing.T) {
  ch := NewCharacter()

  cFx := ch.AtkFx()

  if cFx.Dmg != ch.GetAtk() {
    t.Errorf("Expected default combat action to apply damage equal to attack, but dmg was %d (and atk is %d)", cFx.Dmg, ch.GetAtk())
  }
}

func Test_ResistAtkLowersDmg(t *testing.T) {
  ch := NewCharacter()
  fx := structs.CmbFx{Dmg: 100}

  res := ch.ResistAtk(fx)

  if res.Dmg >= 100 {
    t.Errorf("Expected applying an attack to report lowered damage, but reported %d -> %d", fx.Dmg, res.Dmg)
  }
}

func Test_ResistAtkKeepsStatusEffects(t *testing.T) {
  ch := NewCharacter()
  fx := structs.CmbFx{
    SFx: []statfx.StatusEffect{
      statfx.Stun,
    },
  }

  res := ch.ResistAtk(fx)

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

  cFx := ch.AtkFx()
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

  cFx := ch.AtkFx()
  cFx.Dmg = 10
  ch.ApplyAtk(cFx, rep)

  if ch.Stunned {
    t.Error("Expected character to not be stunned with ApplyAtk with stun, but they were")
  }

  if ch.GetDet() < ch.GetMaxDet() {
    t.Errorf("Expected character to not lose health with ApplyAtk, but they're at %d/%d", ch.GetDet(), ch.GetMaxDet())
  }
}
