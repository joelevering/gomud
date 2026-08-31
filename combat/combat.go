package combat

import (
  "fmt"
  "log"
  "time"

  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/structs"
)

const TickTime = 1500 * time.Millisecond

func Start(pc interfaces.Combatant, npc interfaces.Combatant, rm interfaces.RoomI) {
  // If a panic aborts the fight early (e.g. pc disconnected mid-combat),
  // neither combatant's WinCombat/LoseCombat -- and thus LeaveCombat -- ever
  // runs, so reset combat state manually to avoid leaving either one stuck
  // "in combat" for the rest of the process.
  defer func() {
    if r := recover(); r != nil {
      log.Printf("Recovered from panic in combat between %s and %s: %v", pc.GetName(), npc.GetName(), r)
      clearCombatState(pc)
      clearCombatState(npc)
    }
  }()

  rm.Message(fmt.Sprintf("%s and %s start fighting!", pc.GetName(), npc.GetName()))
  pc.EnterCombat(npc)
  npc.EnterCombat(pc)

  for true {
    combatOver := TickCombat(pc, npc)
    if combatOver {
      rm.Message(fmt.Sprintf("%s emerges victorious over %s!", pc.GetName(), npc.GetName()))
      break
    }
    time.Sleep(TickTime)

    combatOver = TickCombat(npc, pc)
    if combatOver {
      break
    }
    time.Sleep(TickTime)
  }
}

// clearCombatState resets a combatant's in-combat state via LeaveCombat, if
// it has one. Combatant doesn't expose LeaveCombat itself since it's only
// needed here, on the panic-recovery path, so a type assertion avoids
// widening the interface for every other caller.
func clearCombatState(c interfaces.Combatant) {
  if lc, ok := c.(interface{ LeaveCombat() }); ok {
    lc.LeaveCombat()
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
