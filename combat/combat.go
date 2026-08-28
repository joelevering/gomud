package combat

import (
  "fmt"
  "time"

  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/structs"
)

const TickTime = 1500 * time.Millisecond

func Start(pc interfaces.Combatant, npc interfaces.Combatant, rm interfaces.RoomI) {
  rm.Message(fmt.Sprintf("%s and %s start fighting!", pc.GetName(), npc.GetName()))
  pc.EnterCombat(npc)
  npc.EnterCombat(pc)

  for true {
    if pc.WantsToFlee() {
      if pc.AttemptFlee() {
        pc.Flee(npc)
        npc.LeaveCombat()
        rm.Message(fmt.Sprintf("%s flees from the fight with %s!", pc.GetName(), npc.GetName()))
        break
      }
      rm.Message(fmt.Sprintf("%s tries to flee, but fails!", pc.GetName()))
    } else if TickCombat(pc, npc) {
      rm.Message(fmt.Sprintf("%s emerges victorious over %s!", pc.GetName(), npc.GetName()))
      break
    }

    time.Sleep(TickTime)

    if TickCombat(npc, pc) {
      break
    }
    time.Sleep(TickTime)
  }
}

func TickCombat(agg, def interfaces.Combatant) (combatOver bool) {
  report := &structs.CmbRep{}

  aggFx := agg.AtkFx(report)
  resFx := def.ResistAtk(aggFx, report)

  agg.ApplyAtk(resFx, report)
  def.ApplyDef(resFx, report)

  agg.ReportAtk(def, *report)
  def.ReportDef(agg, *report)

  if def.IsDefeated() {
    agg.WinCombat(def)
    def.LoseCombat(agg)
    return true
  }

  agg.TickFx()

  return false
}
