package client

import (
  "testing"

  "github.com/joelevering/gomud/mocks"
)

func Test_LoopDealsDamage(t *testing.T) {
  pc := &mocks.MockCharacter{
    Atk: 50,
    Def: 50,
  }
  npcC := &mocks.MockCharacter{
    Atk: 25,
    Def: 25,
  }
  ci := CombatInstance{
    cli: &Client{},
    npc: &mocks.MockNPC{},
    pc: pc,
    npcC: npcC,
  }

  ci.Loop(false)

  if pc.SetDetArg != 127 {
    t.Errorf("Expected pc to go down from 150 to 127 damage but it went down to %d", pc.SetDetArg)
  }
  if npcC.SetDetArg != 102 {
    t.Errorf("Expected npc to go down from 150 to 102 damage but it went down to %d", npcC.SetDetArg)
  }
}

func Test_PCDefeat(t *testing.T) {
  cli := &mocks.MockClient{}
  pc := &mocks.MockCharacter{ ShouldDie: true }
  npcC := &mocks.MockCharacter{}
  ci := CombatInstance{
    cli: cli,
    npc: &mocks.MockNPC{},
    pc: pc,
    npcC: npcC,
  }

  ci.Loop(false)

  if cli.DefeatedBy != npcC {
    t.Error("Expected PC to be defeated by NPC, but it wasn't")
  }

  if !pc.LeftCombat {
    t.Error("Expected PC to leave combat on defeat, but it did not")
  }
}

func Test_NPCDefeat(t *testing.T) {
  cli := &mocks.MockClient{}
  npc := &mocks.MockNPC{}
  pc := &mocks.MockCharacter{}
  npcC := &mocks.MockCharacter{ ShouldDie: true }

  ci := CombatInstance{
    cli: cli,
    npc: npc,
    pc: pc,
    npcC: npcC,
  }

  ci.Loop(false)

  if npc.DefeatedBy != pc {
    t.Error("Expected NPC to be defeated by PC, but it wasn't")
  }

  if cli.Defeated != npcC {
    t.Error("Expected PC to defeat NPC, but it didn't")
  }

  if !pc.LeftCombat {
    t.Error("Expected PC to leave combat on win, but it didn't")
  }
}

// func Test_NPCDefeat(t *testing.T) {
//   pc := NewCharacter()
//   pc.InCombat = true
// 	cli := &mocks.MockClient{}
//   npc := &mocks.MockNPC{
//     Det:    0,
//   }
//   ci := CombatInstance{
//     cli: cli,
//     pc: pc,
//     npc: npc,
//   }
//
//   ci.Loop(false)
//
//   if pc.IsInCombat() {
//     t.Error("Expected NPC defeat to remove PC from combat, but it's still in combat")
//   }
//
//   if cli.Defeated != npc {
//     t.Error("Expected the client to defeat the NPC, but it didn't")
//   }
//
//   if cli.DefeatedBy != nil {
//     t.Error("Expected the client not to be defeated, but it was")
//   }
//
//   if npc.DefeatedBy != pc {
//     t.Error("Expected the NPC to be defeated by the PC, but it wasn't")
//   }
// }
