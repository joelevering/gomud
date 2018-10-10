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
  npc := &mocks.MockNPC{
    Det: 100,
    Atk: 25,
    Def: 25,
  }
  ci := CombatInstance{
    cli: &Client{Character: pc},
    pc: pc,
    npc: npc,
  }

  ci.Loop(false)

  if pc.SetDetArg != 176 {
    t.Errorf("Expected pc to go down from 199 to 176 damage but it went down to %d", pc.SetDetArg)
  }
  if npc.SetDetArg != 52 {
    t.Errorf("Expected npc to go down from 99 to 52 damage but it went down to %d", npc.SetDetArg)
  }
}

func Test_PCDefeat(t *testing.T) {
  pc := NewCharacter()
  pc.Det = 0
	cli := &mocks.MockClient{}
  npc := &mocks.MockNPC{}
  ci := CombatInstance{
    cli: cli,
    pc: pc,
    npc: npc,
  }

  ci.Loop(false)

  if pc.IsInCombat() {
    t.Error("Expected PC defeat to remove it from combat, but it's still in combat")
  }

  if cli.DefeatedBy != npc {
    t.Error("Expected the client to be defeated by npc, but it wasn't")
  }

  if cli.Defeated != nil {
    t.Error("Expected client to not defeat anyone, but they did")
  }

  if npc.DefeatedBy != nil {
    t.Error("Expected the npc to not be defeated by the client, but it was")
  }
}

func Test_NPCDefeat(t *testing.T) {
  pc := NewCharacter()
  pc.InCombat = true
	cli := &mocks.MockClient{}
  npc := &mocks.MockNPC{
    Det:    0,
  }
  ci := CombatInstance{
    cli: cli,
    pc: pc,
    npc: npc,
  }

  ci.Loop(false)

  if pc.IsInCombat() {
    t.Error("Expected NPC defeat to remove PC from combat, but it's still in combat")
  }

  if cli.Defeated != npc {
    t.Error("Expected the client to defeat the NPC, but it didn't")
  }

  if cli.DefeatedBy != nil {
    t.Error("Expected the client not to be defeated, but it was")
  }

  if npc.DefeatedBy != pc {
    t.Error("Expected the NPC to be defeated by the PC, but it wasn't")
  }
}
