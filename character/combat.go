package character

import (
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
  "github.com/joelevering/gomud/structs"
  "github.com/joelevering/gomud/util"
)

// Character needs to hold all Status Effects related to it

func (ch *Character) AtkFx() structs.CmbFx {
  // Based on current combat skill, locks/retrieves/clears skill and figures out what the effects are (taking status effects into account)
  // Returns a structs.CmbFx obj with summary of intended effects on target

  sk := ch.getAndClearCmbSkill()
  cFx := ch.calcCmbFx(sk)
  return cFx
}

func (ch *Character) ResistAtk(fx structs.CmbFx) structs.CmbFx {
  // Called when a character is being attacked. Applies damage reduction/status effect resistances etc.
  // Calculates damage and effects after factoring in resistances
  // Returns another structs.CmbFx obj with updated summary of damage

  dmg := ch.calcDmg(fx.Dmg)
  sfx := ch.calcSFx(fx.SFx)

  return structs.CmbFx{
    Dmg: dmg,
    SFx: sfx,
  }
}

func (ch *Character) calcCmbFx(sk *skills.Skill) structs.CmbFx {
  res := structs.CmbFx{}
  if ch.Stunned {
    ch.Stunned = false
    // Figure out a way to report the influence of a SE on an attack so we can report "you were stunned!"
    return res
  }

  if sk == nil {
    res.Dmg = ch.GetAtk()
    return res
  }

  for _, e := range sk.Effects {
    switch e.Type {
    case skills.PctDmg:
      res.Dmg = int(float64(ch.GetAtk()) * e.Value.(float64))
    case skills.FlatDmg:
      res.Dmg = ch.GetAtk() + e.Value.(int)
    case skills.OppFx:
      v := e.Value.(statfx.SEInst)
      if (util.RandF() <= v.Chance) {
        res.SFx = append(res.SFx, v.Effect)
      }
    }
  }

  return res
}

func (ch *Character) calcDmg(dmg int) int {
  pctDmgBlocked := float64(ch.GetDef()) * 0.001
  adjDmg := (1.0 - pctDmgBlocked) * float64(dmg)

  return int(adjDmg)
}

func (ch *Character) calcSFx(sfx []statfx.StatusEffect) []statfx.StatusEffect {
  return sfx
}

func (ch *Character) EnterCombat(opp *Character) {
  ch.InCombat = true
}

func (ch *Character) LeaveCombat() {
  ch.InCombat = false
}

func (ch *Character) IsInCombat() bool {
  return ch.InCombat
}
