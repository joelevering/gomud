package combat

import (
  "testing"

  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/mocks"
  "github.com/joelevering/gomud/structs"
)

// panickyPlayer simulates a player who disconnected mid-combat: sending them
// a message (e.g. via ReportAtk) panics, the way sending on a closed
// Player.Channel does in the real game.
type panickyPlayer struct {
  *mocks.MockPlayer
}

func (p *panickyPlayer) ReportAtk(_ interfaces.Combatant, _ structs.CmbRep) {
  panic("send on closed channel")
}

func Test_Start_RecoversFromPanic(t *testing.T) {
  pc := &panickyPlayer{MockPlayer: mocks.NewMockPlayer()}
  npc := mocks.NewMockNP()
  rm := &mocks.MockRoom{}

  defer func() {
    if r := recover(); r != nil {
      t.Fatalf("Expected combat.Start to recover from panic, but it propagated: %v", r)
    }
  }()

  Start(pc, npc, rm)
}

func Test_Start_RecoveryClearsCombatState(t *testing.T) {
  pc := &panickyPlayer{MockPlayer: mocks.NewMockPlayer()}
  npc := mocks.NewMockNP()
  rm := &mocks.MockRoom{}

  Start(pc, npc, rm)

  if !pc.LeftCombat {
    t.Error("Expected pc's combat state to be cleared after the panic was recovered, but LeftCombat is false")
  }

  if !npc.LeftCombat {
    t.Error("Expected npc's combat state to be cleared after the panic was recovered, but LeftCombat is false")
  }
}
