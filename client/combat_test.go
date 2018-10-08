package client

import (
  "testing"

  "github.com/joelevering/gomud/mocks"
  "github.com/joelevering/gomud/npc"
)

func Test_LoopDealsDamage(t *testing.T) {
  pc := &Character{
    Health: 100,
    MaxHealth:100,
    End: 100,
    Str: 100,
  }
	cli := &Client{}
  npc := &npc.NPC{
    Health: 100,
    MaxHealth:100,
    End: 50,
    Str: 50,
  }
  ci := CombatInstance{
    cli: cli,
    pc: pc,
    npc: npc,
  }

  ci.Loop(false)

  if pc.Health != 55 {
    t.Errorf("Expected damage to pc to be calculated as 45 but got %d", pc.MaxHealth - pc.Health)
  }

  if npc.GetHealth() != 5 {
    t.Errorf("Expected damage to npc to be calculated as 95 but got %d", npc.GetMaxHealth() - npc.GetHealth())
  }
}

func Test_PCDefeat(t *testing.T) {
  pc := &Character{
		MaxHealth: 200,
		Health:    200,
    Str:       1,
    End:       0,
    InCombat:  true,
  }

	cli := &mocks.MockClient{}

  npc := &mocks.MockNPC{
    MaxHealth: 100,
    Health:    100,
    Str:       201,
    End:       1000,
  }
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
  pc := &Character{
		MaxHealth: 200,
		Health:    200,
    Str:       101,
    End:       1000,
    InCombat:  true,
  }
	cli := &mocks.MockClient{}
  npc := &mocks.MockNPC{
    MaxHealth: 100,
    Health:    100,
    Str:       1,
    End:       0,
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
