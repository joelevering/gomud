package character

import (
  "math/rand"

  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
  "github.com/joelevering/gomud/structs"
  "github.com/joelevering/gomud/util"
)

func (ch *Character) EnterCombat(opp *Character) {
  ch.InCombat = true
}

func (ch *Character) LeaveCombat() {
  for k := range ch.Fx {
    delete(ch.Fx, k)
  }

  ch.InCombat = false
}

func (ch *Character) IsInCombat() bool {
  return ch.InCombat
}

// Based on current combat skill, locks/retrieves/clears skill and figures out what the effects are (taking status effects into account)
// Returns a structs.CmbFx obj with summary of intended effects on target
func (ch *Character) AtkFx(rep *structs.CmbRep) structs.CmbFx {
  sk := ch.getAndClearCmbSkill()

  if ch.isConcentrating() {
    if sk != nil {
      rep.Skill = *sk
      sk = nil
    }
    rep.Concentrating = true
  } else if sk != nil {
    if ch.payForSkill(*sk) {
      rep.Skill = *sk
    } else { // couldn't pay for skill
      sk = nil
    }
  }

  cFx := ch.calcCmbFx(sk, rep)
  if ch.isWeak() {
    cFx.Dmg /= 2
    rep.LowerAtk = true
  }
  return cFx
}

// Called when a character is being attacked. Applies damage reduction/status effect resistances etc.
// Calculates damage and effects after factoring in resistances
// Returns another structs.CmbFx obj with updated summary of damage
func (ch *Character) ResistAtk(fx structs.CmbFx, rep *structs.CmbRep) structs.CmbFx {
  var dmg int
  if ch.isDodging() {
    rep.Dodged = true
  } else {
    dmg = ch.calcDmg(fx.Dmg)
  }

  sfx := ch.calcSFx(fx.SFx)

  if ch.isVulnerable() {
    dmg *= 2
    rep.LowerDef = true
  }

  return structs.CmbFx{
    Dmg: dmg,
    Heal: fx.Heal,
    SFx: sfx,
    SelfSFx: fx.SelfSFx,
  }
}

// Apply CmbFx for which you are the attacker
func (ch *Character) ApplyAtk(fx structs.CmbFx, rep *structs.CmbRep) {
  if fx.Heal > 0 {
    oldDet := ch.GetDet()
    ch.Heal(fx.Heal)
    healed := ch.GetDet() - oldDet

    rep.Heal = healed
  }

  ch.applySFx(fx.SelfSFx, rep)
  rep.SelfSFx = fx.SelfSFx
}

// Apply CmbFx for which you are the defender
func (ch *Character) ApplyDef(fx structs.CmbFx, rep *structs.CmbRep) {
  if fx.Dmg > 0 {
    ch.SetDet(ch.GetDet() - fx.Dmg)
    rep.Dmg = fx.Dmg
  }

  ch.applySFx(fx.SFx, rep)
  rep.SFx = fx.SFx
}

func (ch *Character) calcCmbFx(sk *skills.Skill, rep *structs.CmbRep) structs.CmbFx {
  fx := structs.CmbFx{}
  if ch.isStunned() {
    rep.Stunned = true

    return fx
  }

  if sk == nil {
    fx.Dmg = ch.GetAtk()
    return fx
  }

  fx.Skill = *sk

  for _, e := range sk.Effects {
    if e.Chance == 0 || (util.RandF() <= e.Chance) {
      switch e.Type {
      case skills.PctDmg:
        fx.Dmg = int(float64(ch.GetAtk()) * e.Value.(float64))
      case skills.FlatDmg:
        fx.Dmg = ch.GetAtk() + e.Value.(int)
      case skills.PctHeal:
        healAmt := float64(ch.GetMaxDet()) * e.Value.(float64)
        fx.Heal = int(healAmt)
      case skills.OppFx:
        v := e.Value.(statfx.SEInst)
        if (util.RandF() <= v.Chance) {
          fx.SFx = append(fx.SFx, v)
        }
      case skills.SelfFx:
        v := e.Value.(statfx.SEInst)
        if (util.RandF() <= v.Chance) {
          fx.SelfSFx = append(fx.SelfSFx, v)
        }
      }
    } else {
      rep.Missed = true
    }
  }

  return fx
}

func (ch *Character) calcDmg(dmg int) int {
  pctDmgBlocked := float64(ch.GetDef()) * 0.001
  adjDmg := (1.0 - pctDmgBlocked) * float64(dmg)

  return int(adjDmg)
}

func (ch *Character) calcSFx(sfx []statfx.SEInst) []statfx.SEInst {
  return sfx
}

func (ch *Character) applySFx(sFx []statfx.SEInst, rep *structs.CmbRep) {
  if len(sFx) > 0 {
    for _, e := range sFx {
      switch e.Effect {
      case statfx.Surprise:
        n := rand.Intn(3)
        if n == 0 {
          srpFx := statfx.SEInst{
            Effect: statfx.Stun,
          }
          ch.addFx(srpFx)

          rep.Surprised = structs.SurpriseRep{Stunned: true}
        } else if n == 1 {
          srpFx := statfx.SEInst{
            Effect: statfx.Weak,
            Duration: (rand.Intn(2)),
          }
          ch.addFx(srpFx)

          rep.Surprised = structs.SurpriseRep{LowerAtk: true}
        } else if n == 2 {
          srpFx := statfx.SEInst{
            Effect: statfx.Vulnerable,
            Duration: (rand.Intn(2)),
          }
          ch.addFx(srpFx)

          rep.Surprised = structs.SurpriseRep{LowerDef: true}
        }
      default:
        ch.addFx(e)
      }
    }
  }
}
