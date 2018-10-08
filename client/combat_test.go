package client

import (
  "testing"

  "github.com/joelevering/gomud/mocks"
  "github.com/joelevering/gomud/npc"
)

func Test_LoopDealsDamage(t *testing.T) {
	cli := &Client{
    Health: 100,
    MaxHealth:100,
    End: 100,
    Str: 100,
	}
  npc := &npc.NPC{
    Health: 100,
    MaxHealth:100,
    End: 50,
    Str: 50,
  }
  ci := CombatInstance{
    cli: cli,
    npc: npc,
  }

  ci.Loop(false)

  if cli.Health != 55 {
    t.Errorf("Expected damage to pc to be calculated as 45 but got %d", cli.MaxHealth - cli.Health)
  }

  if npc.GetHealth() != 5 {
    t.Errorf("Expected damage to npc to be calculated as 95 but got %d", npc.GetMaxHealth() - npc.GetHealth())
  }
}

func Test_LoopCallsDieOnPCDeath(t *testing.T) {
	cli := &mocks.MockClient{
		MaxHealth: 200,
		Health:    200,
    Str:       1,
    End:       0,
	}
  npc := &mocks.MockNPC{
    MaxHealth: 100,
    Health:    100,
    Str:       201,
    End:       1000,
  }
  ci := CombatInstance{
    cli: cli,
    npc: npc,
  }

  ci.Loop(false)

  if cli.KilledBy != npc {
    t.Error("Expected the client to be killed by npc, but it wasn't")
  }

  if cli.Defeated != nil {
    t.Error("Expected client to not defeat anyone, but they did")
  }

  if npc.KilledBy != nil {
    t.Error("Expected the npc to not be killed by the client, but it was")
  }
}

func Test_LoopCallsDefeatOnNPCDeath(t *testing.T) {
	cli := &mocks.MockClient{
		MaxHealth: 200,
		Health:    200,
    Str:       101,
    End:       1000,
	}
  npc := &mocks.MockNPC{
    MaxHealth: 100,
    Health:    100,
    Str:       1,
    End:       0,
  }
  ci := CombatInstance{
    cli: cli,
    npc: npc,
  }

  ci.Loop(false)

  if cli.Defeated != npc {
    t.Error("Expected the client to defeat the NPC, but it didn't")
  }

  if cli.KilledBy != nil {
    t.Error("Expected the client not to be killed, but it was")
  }

  if npc.KilledBy != cli {
    t.Error("Expected the NPC to be killed by the client, but it wasn't")
  }
}
