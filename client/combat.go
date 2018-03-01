package client

import (
	"fmt"
	"time"

	"github.com/joelevering/gomud/interfaces"
)

type CombatInstance struct {
	cli  *Client
	npc  interfaces.NPCI
	tick int // seconds
}

func (ci CombatInstance) Start() {
	ci.tick = 1

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

		ci.cli.SendMsg(fmt.Sprintf("%s took %d damage!", ci.npc.GetName(), npcDmg))
		ci.cli.SendMsg(fmt.Sprintf("You took %d damage!", pcDmg))
		ci.cli.SendMsg(fmt.Sprintf("You have %d/%d health left. %s has %d/%d", ci.cli.Health, ci.cli.MaxHealth, ci.npc.GetName(), ci.npc.GetHealth(), ci.npc.GetMaxHealth()))
		time.Sleep(ci.tickDuration())
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

func (ci CombatInstance) tickDuration() time.Duration {
	return time.Duration(ci.tick) * time.Second
}
