package combat

import (
  "log"
  "time"

  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/structs"
)

const tickTime = 1500 * time.Millisecond

func Start(pc interfaces.Combatant, npc interfaces.Combatant) {
  pc.EnterCombat(npc)
  npc.EnterCombat(pc)

  for true {
    combatOver := TickCombat(pc, npc)
    if combatOver { break }
    time.Sleep(tickTime)

    combatOver = TickCombat(npc, pc)
    if combatOver { break }
    time.Sleep(tickTime)
  }
}

func TickCombat(agg, def interfaces.Combatant) (combatOver bool) {
  aggFx := agg.AtkFx()
  log.Printf("%s fx: %v", agg.GetName(), aggFx)
  resFx := def.ResistAtk(aggFx)

  report := &structs.CmbRep{}
  agg.ApplyAtk(resFx, report)
  def.ApplyDef(resFx, report)

  agg.ReportAtk(def, *report)
  def.ReportDef(agg, *report)

  if def.IsDefeated() {
    agg.WinCombat(def)
    def.LoseCombat(agg)
    return true
  }

  return false
}
