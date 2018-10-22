package character

import (
  "testing"

  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
)

func Test_CharAttacksByDefault(t *testing.T) {
  ch := NewCharacter()

  cFx := ch.CmbTick()

  if cFx.Dmg != ch.GetAtk() {
    t.Errorf("Expected default combat action to apply damage equal to attack, but dmg was %d (and atk is %d)", cFx.Dmg, ch.GetAtk())
  }
}

func Test_StunAppliesStun(t *testing.T) {
  ch := NewCharacter()
  ch.SetCmbSkill(skills.Stun)

  cFx := ch.CmbTick()

  if len(cFx.SFx) != 1 || cFx.SFx[0] != statfx.Stun {
    t.Errorf("Expected Stun to apply only stun effect, but it applied %v", cFx.SFx)
  }
}
