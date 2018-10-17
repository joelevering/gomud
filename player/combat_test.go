package player

import (
  "testing"

  "github.com/joelevering/gomud/mocks"
  "github.com/joelevering/gomud/skills"
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
  if !pc.ClearedCmbSkill || !npc.ClearedCmbSkill {
    t.Error("Expected both PC and NPC combat skills to be cleared on loop, but they weren't")
  }

  if !pc.LockedCmbSkill || !pc.UnlockedCmbSkill {
    t.Error("Expected PC combat skills to be locked and then unlocked in a CI loop")
  }

  if !npc.LockedCmbSkill || !npc.UnlockedCmbSkill {
    t.Error("Expected NPC combat skills to be locked and then unlocked in a CI loop")
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

func Test_StunNegatesDamage(t *testing.T) {
  pc := &mocks.MockPlayer{
    MockCharacter: &mocks.MockCharacter{
      Atk: 50,
      Def: 50,
      CmbSkill: skills.Stun,
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

  if pc.SetDetArg != pc.GetDet() {
    t.Errorf("Expected PC health to stay static when NPC is stunned, but health was set to %d (from %d)", pc.SetDetArg, pc.GetDet())
  }
}

func Test_HealRestoresDet(t *testing.T) {
  pc := mocks.NewMockPlayer()
  pc.CmbSkill = skills.PowerNap
  npc := mocks.NewMockNP()

  ci := CombatInstance{
    pc: pc,
    npc: npc,
  }

  ci.Loop(false)

  if pc.SetDetArg <= pc.GetDet() {
    t.Errorf("Expected PC health to increase when using a heal, but it was set to %d (from %d)", pc.SetDetArg, pc.GetDet())
  }
}
