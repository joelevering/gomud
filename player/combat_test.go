package player

import (
  "testing"

  "github.com/joelevering/gomud/mocks"
)

func Test_LoopDealsDamage(t *testing.T) {
  pc := &mocks.MockPlayer{
    MockCharacter: &mocks.MockCharacter{
      Atk: 50,
      Def: 50,
    },
  }
  npc := &mocks.MockNP{
    MockCharacter: &mocks.MockCharacter{
      Atk: 25,
      Def: 25,
    },
  }
  ci := CombatInstance{
    pc: pc,
    npc: npc,
  }

  ci.Loop(false)

  if pc.SetDetArg != 127 {
    t.Errorf("Expected pc to go down from 150 to 127 damage but it went down to %d", pc.SetDetArg)
  }
  if npc.SetDetArg != 102 {
    t.Errorf("Expected npc to go down from 150 to 102 damage but it went down to %d", npc.SetDetArg)
  }
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
}
