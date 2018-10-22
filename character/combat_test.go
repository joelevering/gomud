package character

import (
  "testing"

  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
)

func Test_CharAttacksByDefault(t *testing.T) {
  ch := NewCharacter()

  cFx := ch.AtkFx()

  if cFx.Dmg != ch.GetAtk() {
    t.Errorf("Expected default combat action to apply damage equal to attack, but dmg was %d (and atk is %d)", cFx.Dmg, ch.GetAtk())
  }
}

func Test_StunAppliesStun(t *testing.T) {
  ch := NewCharacter()
  ch.SetCmbSkill(skills.Stun)

  cFx := ch.AtkFx()

  if len(cFx.SFx) != 1 || cFx.SFx[0] != statfx.Stun {
    t.Errorf("Expected Stun to apply only stun effect, but it applied %v", cFx.SFx)
  }
}

func Test_ResistAtkLowersDmg(t *testing.T) {
  ch := NewCharacter()
  fx := CombatEffects{Dmg: 100}

  res := ch.ResistAtk(fx)

  if res.Dmg >= 100 {
    t.Errorf("Expected applying an attack to report lowered damage, but reported %d -> %d", fx.Dmg, res.Dmg)
  }
}

func Test_ResistAtkKeepsStatusEffects(t *testing.T) {
  ch := NewCharacter()
  fx := CombatEffects{
    SFx: []statfx.StatusEffect{
      statfx.Stun,
    },
  }

  res := ch.ResistAtk(fx)

  if len(res.SFx) != 1 || res.SFx[0] != statfx.Stun {
    t.Errorf("Expected status effects to remain the same on resist, but got %v", res.SFx)
  }
}
