package player

import (
  "fmt"
  "time"

  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
  "github.com/joelevering/gomud/util"
)

const tick = 1500 * time.Millisecond

type CombatResults struct {
  pcDmg      int
  npcDmg     int
  pcHeal     int
  npcHeal    int
  pcEffects  []statfx.StatusEffect
  npcEffects []statfx.StatusEffect
}

type CombatInstance struct {
  pc  interfaces.PlI
  npc interfaces.NPI
}

func (ci *CombatInstance) Start() {
  ci.pc.EnterCombat()

  for true {
    combatOver := ci.Loop(true)
    if combatOver { break }

    time.Sleep(tick)
  }
}

func (ci *CombatInstance) getSkills() (pcSk, npcSk *skills.Skill) {
  ci.pc.LockCmbSkill()
  ci.npc.LockCmbSkill()

  return ci.pc.GetCmbSkill(), ci.npc.GetCmbSkill()
}

func (ci *CombatInstance) clearSkills() {
  ci.pc.UnlockCmbSkill()
  ci.pc.ClearCmbSkill() // Low probability race condition

  ci.npc.UnlockCmbSkill()
  ci.npc.ClearCmbSkill()
}

func (ci *CombatInstance) Loop(report bool) (combatOver bool) {
  pcSk, npcSk := ci.getSkills()
  defer ci.clearSkills()

  pcRes := ci.getPCResults(pcSk)
  npcRes := ci.getNPCResults(npcSk)

  ci.applySFX(pcRes, npcRes)
  ci.applyDamage(pcRes, npcRes)

  if report {
    ci.report(pcRes, npcRes)
  }

  if ci.pcIsDead() {
    ci.pc.LeaveCombat()
    ci.pc.LoseCombat(ci.npc)
    return true
  }

  if ci.npcIsDead() {
    ci.pc.LeaveCombat()
    ci.npc.LoseCombat(ci.pc)
    ci.pc.WinCombat(ci.npc)
    return true
  }

  return false
}

func (ci *CombatInstance) getPCResults(sk *skills.Skill) *CombatResults {
  res := &CombatResults{}

  if sk == nil {
    res.npcDmg = CalcAtkDmg(ci.pc.GetAtk(), ci.npc.GetDef())
    return res
  }

  for _, e := range sk.Effects {
    switch e.Type {
    case skills.PctDmg:
      baseDmg := CalcAtkDmg(ci.pc.GetAtk(), ci.npc.GetDef())
      res.npcDmg = int(float64(baseDmg) * e.Value.(float64))
    case skills.FlatDmg:
      baseDmg := CalcAtkDmg(ci.pc.GetAtk(), ci.npc.GetDef())
      res.npcDmg = baseDmg + e.Value.(int)
    case skills.PctHeal:
      healAmt := float64(ci.pc.GetMaxDet()) * e.Value.(float64)
      res.pcHeal = int(healAmt)
    case skills.OppFx:
      v := e.Value.(statfx.SEInst)
      if (util.RandF() <= v.Chance) {
        res.npcEffects = append(res.npcEffects, v.Effect)
      }
    }
  }

  return res
}

func (ci *CombatInstance) getNPCResults(_ *skills.Skill) *CombatResults {
  return &CombatResults{
    pcDmg: CalcAtkDmg(ci.npc.GetAtk(), ci.pc.GetDef()),
  }
}

func (ci *CombatInstance) applySFX(pcRes, npcRes *CombatResults) {
  for _, pcFx := range [][]statfx.StatusEffect{pcRes.pcEffects, npcRes.pcEffects} {
    for _, e := range pcFx {
      if e == statfx.Stun {
        pcRes.pcDmg = 0
        pcRes.npcDmg = 0
      }
    }
  }

  for _, npcFx := range [][]statfx.StatusEffect{pcRes.npcEffects, npcRes.npcEffects} {
    for _, e := range npcFx {
      if e == statfx.Stun {
        npcRes.pcDmg = 0
        npcRes.npcDmg = 0
      }
    }
  }
}

// Block 1/1000th of the damage per point of Endurance
func CalcAtkDmg(atkStr int, defEnd int) int {
  endPercentDamageBlocked := float64(defEnd) * 0.001
  dmg := (1.0 - endPercentDamageBlocked) * float64(atkStr)

  return int(dmg)
}

func (ci *CombatInstance) applyDamage(pcRes, npcRes *CombatResults) {
  pcHeal := pcRes.pcHeal + npcRes.pcHeal
  pcDmg := pcRes.pcDmg + npcRes.pcDmg - pcHeal
  npcHeal := pcRes.npcHeal + npcRes.npcHeal
  npcDmg := pcRes.npcDmg + npcRes.npcDmg - npcHeal

  npcDet := ci.npc.GetDet()
  pcDet := ci.pc.GetDet()

  if pcDet-pcDmg < 0 {
    ci.pc.SetDet(0)
  } else if pcDet-pcDmg > ci.pc.GetMaxDet() {
    ci.pc.Heal()
  } else {
    ci.pc.SetDet(pcDet - pcDmg)
  }

  if npcDet-npcDmg < 0 {
    ci.npc.SetDet(0)
  } else if npcDet-npcDmg > ci.npc.GetMaxDet() {
    ci.npc.Heal()
  } else {
    ci.npc.SetDet(npcDet - npcDmg)
  }
}

func (ci *CombatInstance) pcIsDead() bool {
  return ci.pc.GetDet() <= 0
}

func (ci *CombatInstance) npcIsDead() bool {
  return ci.npc.GetDet() <= 0
}

func (ci *CombatInstance) report(pcRes, npcRes *CombatResults) {
  pcHealTotal := pcRes.pcHeal + npcRes.pcHeal
  if pcHealTotal > 0 {
    ci.pc.SendMsg(fmt.Sprintf("You recovered %d determination!", pcHealTotal))
  }

  npcHealTotal := pcRes.npcHeal + npcRes.npcHeal
  if npcHealTotal > 0 {
    ci.pc.SendMsg(fmt.Sprintf("%s recovered %d determination!", ci.npc.GetName(), npcHealTotal))
  }

  npcDmgTotal := pcRes.npcDmg + npcRes.npcDmg
  if npcDmgTotal > 0 {
    ci.pc.SendMsg(fmt.Sprintf("%s took %d damage!", ci.npc.GetName(), npcDmgTotal))
  }

  pcDmgTotal := pcRes.pcDmg + npcRes.pcDmg
  if pcDmgTotal > 0 {
    ci.pc.SendMsg(fmt.Sprintf("You took %d damage!", pcDmgTotal))
  }

  ci.pc.SendMsg("")

  if ci.pcIsDead() {
    return // handled in Start()
  }

  var npcMsg string

  if ci.npcIsDead() {
    npcMsg = fmt.Sprintf("%s is defeated!", ci.npc.GetName())
  } else {
    npcMsg = fmt.Sprintf("%s has %d/%d", ci.npc.GetName(), ci.npc.GetDet(), ci.npc.GetMaxDet())
  }

  ci.pc.SendMsg(fmt.Sprintf("You have %d/%d health left. %s", ci.pc.GetDet(), ci.pc.GetMaxDet(), npcMsg))
}
