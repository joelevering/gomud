package player

import (
  "fmt"
  "math/rand"
  "time"

  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
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

func (ci *CombatInstance) Loop(report bool) (combatOver bool) {
  defer ci.pc.ClearCmbSkill()
  defer ci.npc.ClearCmbSkill()

  pcResults := ci.getPCResults()
  npcResults := ci.getNPCResults()

  npcDmg := pcResults.npcDmg + npcResults.npcDmg
  pcDmg := pcResults.pcDmg + npcResults.pcDmg

  ci.applyDamage(npcDmg, pcDmg)
  if report {
    ci.report(npcDmg, pcDmg)
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

func (ci *CombatInstance) getPCResults() *CombatResults {
  res := &CombatResults{}
  sk := ci.pc.GetCmbSkill()

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
    case skills.OppFx:
      v := e.Value.(statfx.SEInst)
      if (rand.Float64() <= v.Chance) {
        res.npcEffects = append(res.npcEffects, v.Effect)
      }
    }
  }

  return res
}

func (ci *CombatInstance) getNPCResults() *CombatResults {
  return &CombatResults{
    pcDmg: CalcAtkDmg(ci.npc.GetAtk(), ci.pc.GetDef()),
  }
}

// Block 1/1000th of the damage per point of Endurance
func CalcAtkDmg(atkStr int, defEnd int) int {
  endPercentDamageBlocked := float64(defEnd) * 0.001
  dmg := (1.0 - endPercentDamageBlocked) * float64(atkStr)

  return int(dmg)
}

func (ci *CombatInstance) applyDamage(npcDmg, pcDmg int) {
  npcDet := ci.npc.GetDet()
  pcDet := ci.pc.GetDet()

  if pcDet-pcDmg < 0 {
    ci.pc.SetDet(0)
  } else {
    ci.pc.SetDet(pcDet - pcDmg)
  }

  if npcDet-npcDmg < 0 {
    ci.npc.SetDet(0)
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

func (ci *CombatInstance) report(npcDmg, pcDmg int) {
  ci.pc.SendMsg(fmt.Sprintf("%s took %d damage!", ci.npc.GetName(), npcDmg))
  ci.pc.SendMsg(fmt.Sprintf("You took %d damage!", pcDmg))
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
