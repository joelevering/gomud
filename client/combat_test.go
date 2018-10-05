package client

import (
  "testing"
  "time"

  "github.com/joelevering/gomud/npc"
)

func Test_CalculateDamageWorks(t *testing.T) {
	cli := &Client{
    End: 100,
    Str: 100,
	}
  npc := &npc.NPC{
    End: 50,
    Str: 50,
  }
  ci := CombatInstance{
    cli: cli,
    npc: npc,
    tick: time.Duration(1) * time.Millisecond,
  }

  npcDmg, pcDmg := ci.calculateDamages()

  if npcDmg != 95 {
    t.Errorf("Expected npc damage to be calculated as 95 but got %d", npcDmg)
  }

  if pcDmg != 45 {
    t.Errorf("Expected pc damage to be calculated as 45 but got %d", pcDmg)
  }
}

func Test_ApplyDamageRemovesHealth(t *testing.T) {
	cli := &Client{
		MaxHealth:     200,
		Health:        200,
	}
  npc := &npc.NPC{
    MaxHealth: 100,
    Health:    100,
  }
  ci := CombatInstance{
    cli: cli,
    npc: npc,
    tick: time.Duration(1) * time.Millisecond,
  }

  ci.applyDamage(20, 10)

  if npc.GetHealth() != 80 {
    t.Errorf("Expected npc to lose 20 health, but they lost %d", npc.GetMaxHealth() - npc.GetHealth())
  }

  if cli.Health != 190 {
    t.Errorf("Expected cli to lose 10 health, but they lost %d", cli.MaxHealth - cli.Health)
  }
}

func Test_ApplyDamageZeroesHealthForLethalDamage(t *testing.T) {
	cli := &Client{
		MaxHealth:     200,
		Health:        200,
	}
  npc := &npc.NPC{
    MaxHealth: 100,
    Health:    100,
  }
  ci := CombatInstance{
    cli: cli,
    npc: npc,
    tick: time.Duration(1) * time.Millisecond,
  }

  ci.applyDamage(101, 201)

  if npc.GetHealth() != 0 {
    t.Errorf("Expected npc health to be 0 but it's %d", npc.GetHealth())
  }

  if cli.Health != 0 {
    t.Errorf("Expect cli health to be 0 but it's %d", cli.Health)
  }
}
