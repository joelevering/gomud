package client

import (
	"fmt"
	"time"

	"github.com/joelevering/gomud/interfaces"
)

type CombatInstance struct {
	cli  *Client
	npc  interfaces.NPCI
	npc  *NPC
	tick time.Duration
}

func (ci CombatInstance) Start() {
	ci.tick = time.Duration(1000) * time.Millisecond

	for ci.noOneIsDead() {
		npcDmg := ci.calculateDamage(ci.cli.Str, ci.npc.GetEnd())
		pcDmg := ci.calculateDamage(ci.npc.GetStr(), ci.cli.End)

		if ci.npc.GetHealth()-npcDmg < 0 {
			ci.npc.SetHealth(0)
		} else {
			ci.npc.SetHealth(ci.npc.GetHealth() - npcDmg)
		}

		if ci.cli.Health-pcDmg < 0 {
			ci.cli.Health = 0
		} else {
			ci.cli.Health = ci.cli.Health - pcDmg
		}

    ci.report(npcDmg, pcDmg)
		time.Sleep(ci.tick)
	}
}

func (ci CombatInstance) noOneIsDead() bool {
	if ci.cli.Health <= 0 || ci.npc.GetHealth() <= 0 {
		return false
	}

	return true
}

// Block 1/1000th of the damage per point of Endurance
func (ci CombatInstance) calculateDamage(atkStr, defEnd int) int {
	endPercentDamageBlocked := float64(defEnd) * 0.001
	dmg := (1.0 - endPercentDamageBlocked) * float64(atkStr)
	return int(dmg)
}

func (ci CombatInstance) report(npcDmg, pcDmg int) {
  ci.cli.SendMsg(fmt.Sprintf("%s took %d damage!", ci.npc.GetName(), npcDmg))
  ci.cli.SendMsg(fmt.Sprintf("You took %d damage!", pcDmg))

  if ci.cli.Health <= 0 {
    ci.cli.SendMsg("You've been defeated.")
    return
  }

  var npcMsg string

  if ci.npc.GetHealth() <= 0 {
    npcMsg = fmt.Sprintf("%s is deafeated!", ci.npc.GetName())
  } else {
    npcMsg = fmt.Sprintf("%s has %d/%d", ci.npc.GetName(), ci.npc.GetHealth(), ci.npc.GetMaxHealth())
  }

  ci.cli.SendMsg(fmt.Sprintf("You have %d/%d health left. %s", ci.cli.Health, ci.cli.MaxHealth, npcMsg))
}
