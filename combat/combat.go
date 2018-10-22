package combat

import (
  "time"

  "github.com/joelevering/gomud/interfaces"
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
  resFx := def.ResistAtk(aggFx)
  agg.ReportAtk(def, resFx)
  def.ReportDef(agg, resFx)

  if def.IsDefeated() {
    agg.WinCombat(def)
    def.LoseCombat(agg)
    return true
  }

  return false
}
