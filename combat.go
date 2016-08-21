package main

import (
	"fmt"
	"time"
)

type CombatInstance struct {
	cli  *Client
	npc  *NPC
	tick int // seconds
}

func (ci CombatInstance) Start() {
	ci.tick = 1

	for ci.noOneIsDead() {
		ci.npc.Health = ci.npc.Health - 1
		ci.cli.SendMsg(fmt.Sprintf("%s took 1 damage!", ci.npc.Name))
		ci.cli.SendMsg(fmt.Sprintf("You have %d/%d health left. %s has %d/%d", ci.cli.Health, ci.cli.MaxHealth, ci.npc.Name, ci.npc.Health, ci.npc.MaxHealth))
		time.Sleep(ci.tickDuration())
	}
}

func (ci CombatInstance) noOneIsDead() bool {
	if ci.cli.Health <= 0 || ci.npc.Health <= 0 {
		return false
	}

	return true
}

func (ci CombatInstance) tickDuration() time.Duration {
	return time.Duration(ci.tick) * time.Second
}
