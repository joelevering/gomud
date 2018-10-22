package player

// import (
//   "fmt"
//   "log"
//   "time"
//
//   "github.com/joelevering/gomud/interfaces"
//   "github.com/joelevering/gomud/skills"
//   "github.com/joelevering/gomud/statfx"
//   "github.com/joelevering/gomud/util"
// )

const tick = 1500 * time.Millisecond

type CombatInstance struct {
  pc  interfaces.PlI
  npc interfaces.NPI
}

func (ci *CombatInstance) Start() {
  ci.pc.EnterCombat()

  for true {
    combatOver := ci.Tick(pc, npc)
    if combatOver { break }
    time.Sleep(tick)

    combatOver := ci.Tick(npc, pc)
    if combatOver { break }
    time.Sleep(tick)
  }
}

func (ci *CombatInstance) Tick(agg *Character, def *Character) (combatOver bool) {
  aggFx := agg.AtkFx()
  resFx := def.ResistAtk(aggFx)
  agg.ReportAtk(resFx)
  def.ReportDef(resFx)

  if def.IsDefeated() {
    agg.WinCombat(def)
    def.LoseCombat(agg)
    return true
  }

  return false
}

// func (ci *CombatInstance) applySFX(pcRes, npcRes *CombatResults) {
//   for _, pcFx := range [][]statfx.StatusEffect{pcRes.pcEffects, npcRes.pcEffects} {
//     for _, e := range pcFx {
//       if e == statfx.Stun {
//         pcRes.pcDmg = 0
//         pcRes.npcDmg = 0
//       }
//     }
//   }
//
//   for _, npcFx := range [][]statfx.StatusEffect{pcRes.npcEffects, npcRes.npcEffects} {
//     for _, e := range npcFx {
//       if e == statfx.Stun {
//         npcRes.pcDmg = 0
//         npcRes.npcDmg = 0
//       }
//     }
//   }
// }

// // Block 1/1000th of the damage per point of Endurance
// func CalcAtkDmg(atkStr int, defEnd int) int {
//   endPercentDamageBlocked := float64(defEnd) * 0.001
//   dmg := (1.0 - endPercentDamageBlocked) * float64(atkStr)
//
//   return int(dmg)
// }
//
// func (ci *CombatInstance) applyDamage(npcDmg, pcDmg int) {
//   npcDet := ci.npc.GetDet()
//   pcDet := ci.pc.GetDet()
//
//   if pcDet-pcDmg < 0 {
//     ci.pc.SetDet(0)
//   } else {
//     ci.pc.SetDet(pcDet - pcDmg)
//   }
//
//   if npcDet-npcDmg < 0 {
//     ci.npc.SetDet(0)
//   } else {
//     ci.npc.SetDet(npcDet - npcDmg)
//   }
// }

// func (ci *CombatInstance) report(npcDmg, pcDmg int) {
//   ci.pc.SendMsg(fmt.Sprintf("%s took %d damage!", ci.npc.GetName(), npcDmg))
//   ci.pc.SendMsg(fmt.Sprintf("You took %d damage!", pcDmg))
//   ci.pc.SendMsg("")
//
//   if ci.pcIsDead() {
//     return // handled in Start()
//   }
//
//   var npcMsg string
//
//   if ci.npcIsDead() {
//     npcMsg = fmt.Sprintf("%s is defeated!", ci.npc.GetName())
//   } else {
//     npcMsg = fmt.Sprintf("%s has %d/%d", ci.npc.GetName(), ci.npc.GetDet(), ci.npc.GetMaxDet())
//   }
//
//   ci.pc.SendMsg(fmt.Sprintf("You have %d/%d health left. %s", ci.pc.GetDet(), ci.pc.GetMaxDet(), npcMsg))
// }
