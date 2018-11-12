package character

import (
  "math/rand"

  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
  "github.com/joelevering/gomud/structs"
  "github.com/joelevering/gomud/util"
)

func (ch *Character) LeaveCombat() {
  for e := range ch.Fx {
    delete(ch.Fx, e)
  }

  for d := range ch.Dots {
    delete(ch.Dots, d)
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
      rep.Skill = *sk // report the skill
      sk = nil // but don't actually use it
    }
    rep.Concentrating = true
  }

  if sk != nil {
    hasSelfReq, effect := sk.SelfFollowUpReq()
    if hasSelfReq && !ch.hasEffect(effect) {
      rep.FollowUpReq = effect
      rep.Skill = *sk
      return structs.CmbFx{}
    }

    if ch.payForSkill(*sk) {
      rep.Skill = *sk
    } else { // couldn't pay for skill
      sk = nil
    }
  }

  cFx := ch.calcCmbFx(sk, rep)
  if ch.isWeak() && cFx.Dmg > 1 {
    cFx.Dmg /= 2
    rep.Weak = true
  }
  if ch.isEmpowered() {
    cFx.Dmg *= 2
    if cFx.Dmg > 0 {
      rep.Empowered = true
    }
  }

  return cFx
}

func (ch *Character) ResistAtk(fx structs.CmbFx, rep *structs.CmbRep) structs.CmbFx {
  newFx := structs.CmbFx{
    DotDmgs: fx.DotDmgs,
  }

  if fx.Req != "" && !ch.hasEffect(fx.Req) {
    rep.FollowUpReq = fx.Req
    return newFx
  }

  newFx.Heal = fx.Heal
  newFx.StmRec = fx.StmRec
  newFx.SelfSFx = fx.SelfSFx
  newFx.Dots = fx.Dots

  var dmg, selfDmg int
  if ch.isDodging() {
    rep.Dodged = true
  } else if ch.isRedirecting() {
    dmg = fx.Dmg/2
    selfDmg = dmg + fx.SelfDmg
    rep.Redirected = true
  } else {
    dmg = ch.calcDmg(fx.Dmg)
  }

  sfx := ch.calcSFx(fx.SFx)

  if ch.isVulnerable() {
    dmg *= 2
    if dmg != 0 {
      rep.Vulnerable = true
    }
  }

  if ch.isSteeled() && dmg >1 {
    dmg /= 2
    rep.Steeled = true
  }

  newFx.Dmg = dmg
  newFx.SelfDmg = selfDmg
  newFx.SFx = sfx

  return newFx
}

// Apply CmbFx for which you are the attacker
func (ch *Character) ApplyAtk(fx structs.CmbFx, rep *structs.CmbRep) {
  ch.SetDet(ch.GetDet() - fx.SelfDmg)
  rep.SelfDmg = fx.SelfDmg

  for _, dotDmg := range fx.DotDmgs {
    ch.SetDet(ch.GetDet() - dotDmg.Dmg)
  }
  rep.DotDmgs = fx.DotDmgs

  rep.Heal = ch.recoverDet(fx.Heal)
  rep.StmRec = ch.recoverStm(fx.StmRec)

  ch.applySFx(fx.SelfSFx, rep)
  rep.SelfSFx = fx.SelfSFx
}

// Apply CmbFx for which you are the defender
func (ch *Character) ApplyDef(fx structs.CmbFx, rep *structs.CmbRep) {
  ch.SetDet(ch.GetDet() - fx.Dmg)
  rep.Dmg = fx.Dmg

  ch.applySFx(fx.SFx, rep)
  rep.SFx = fx.SFx

  ch.applyDots(fx.Dots, rep)
  rep.Dots = fx.Dots
}

func (ch *Character) selfCmbFx() structs.CmbFx {
  dots := []statfx.DotInst{}
  for _, d := range ch.Dots {
    dots = append(dots, *d)
  }

  return structs.CmbFx{DotDmgs: dots}
}

func (ch *Character) calcCmbFx(sk *skills.Skill, rep *structs.CmbRep) structs.CmbFx {
  fx := ch.selfCmbFx()

  if ch.isStunned() {
    rep.Stunned = true

    return fx
  }

  if sk == nil {
    fx.Dmg = ch.GetAtk()
    return fx
  }

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
      case skills.FlatStm:
        fx.StmRec = e.Value.(int)
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
      case skills.Dot:
        v := e.Value.(statfx.DotFx)
        if (util.RandF() <= v.Chance) {
          duration := util.RandI(v.DurationMin, v.DurationMax)
          i := statfx.DotInst{
            Type: v.Type,
            Dmg: ch.GetAtk(),
            Duration: duration,
          }
          fx.Dots = append(fx.Dots, i)
        }
      }
    } else {
      rep.Missed = true
    }
  }

  hasOppReq, effect := sk.OppFollowUpReq()
  if hasOppReq {
    fx.Req = effect
  }

  fx.Skill = *sk

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
  for _, e := range sFx {
    if e.Effect == statfx.Surprise {
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

        rep.Surprised = structs.SurpriseRep{Weak: true}
      } else if n == 2 {
        srpFx := statfx.SEInst{
          Effect: statfx.Vulnerable,
          Duration: (rand.Intn(2)),
        }
        ch.addFx(srpFx)

        rep.Surprised = structs.SurpriseRep{Vulnerable: true}
      }
    }

    ch.addFx(e)
  }
}

func (ch *Character) applyDots(dots []statfx.DotInst, rep *structs.CmbRep) {
  for _, d := range dots {
    ch.addDot(d)
  }
}

func (ch *Character) recoverDet(amt int) int {
  if amt > 0 {
    oldDet := ch.GetDet()
    ch.Heal(amt)
    return ch.GetDet() - oldDet
  }

  return 0
}

func (ch *Character) recoverStm(amt int) int {
  if amt > 0 {
    oldStm := ch.GetStm()
    ch.Recuperate(amt)
    return ch.GetStm() - oldStm
  }

  return 0
}
