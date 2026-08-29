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
    // Victory is only announced to the room on the pc's turn -- when the npc
    // wins, LoseCombat already narrates the pc's defeat to the room itself.
    if takeTurn(pc, npc, rm, true) {
      break
    }
    time.Sleep(TickTime)

    if takeTurn(npc, pc, rm, false) {
      break
    }
    time.Sleep(TickTime)
  }
}

// takeTurn resolves one combatant's turn against their opponent, checking a
// flee attempt before falling back to a normal combat tick. It returns
// whether combat is now over. Since it operates purely on the Combatant
// interface, it works unmodified if a non-pc combatant is ever given the
// ability to flee.
func takeTurn(actor, opponent interfaces.Combatant, rm interfaces.RoomI, announceVictory bool) (combatOver bool) {
  if actor.WantsToFlee() {
    if actor.AttemptFlee() {
      actor.Flee(opponent)
      opponent.LeaveCombat()
      rm.Message(fmt.Sprintf("%s flees from the fight with %s!", actor.GetName(), opponent.GetName()))
      return true
    }
    rm.Message(fmt.Sprintf("%s tries to flee, but fails!", actor.GetName()))
    return false
  }

  if TickCombat(actor, opponent) {
    if announceVictory {
      rm.Message(fmt.Sprintf("%s emerges victorious over %s!", actor.GetName(), opponent.GetName()))
    }
    return true
  }

  return false
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
