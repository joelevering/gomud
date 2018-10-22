package player

import (
  "testing"

  "github.com/joelevering/gomud/mocks"
  "github.com/joelevering/gomud/skills"
)

func Test_LoopDealsDamage(t *testing.T) {
}

func Test_PCDefeat(t *testing.T) {
  pc := &mocks.MockPlayer{
    MockCharacter: &mocks.MockCharacter{ ShouldDie: true },
  }
  npc := mocks.NewMockNP()
  ci := CombatInstance{
    pc: pc,
    npc: npc,
  }

  ci.Loop(false)

  if pc.DefeatedBy != npc {
    t.Error("Expected PC to be defeated by NPC, but it wasn't")
  }

  if !pc.LeftCombat {
    t.Error("Expected PC to leave combat on defeat, but it did not")
  }

  if !pc.ClearedCmbSkill || !npc.ClearedCmbSkill {
    t.Error("Expected both PC and NPC combat skills to be cleared on PC defeat, but they weren't")
  }
}

func Test_NPCDefeat(t *testing.T) {
  pc := mocks.NewMockPlayer()
  npc := &mocks.MockNP{
    MockCharacter: &mocks.MockCharacter{ ShouldDie: true },
  }

  ci := CombatInstance{
    pc: pc,
    npc: npc,
  }

  ci.Loop(false)

  if npc.DefeatedBy != pc {
    t.Error("Expected NPC to be defeated by PC, but it wasn't")
  }

  if pc.Defeated != npc {
    t.Error("Expected PC to defeat NPC, but it didn't")
  }

  if !pc.LeftCombat {
    t.Error("Expected PC to leave combat on win, but it didn't")
  }

  if !pc.ClearedCmbSkill || !npc.ClearedCmbSkill {
    t.Error("Expected both PC and NPC combat skills to be cleared on NPC defeat, but they weren't")
  }
}
