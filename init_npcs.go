package main

import "github.com/joelevering/gomud/interfaces"

func InitNPCs(rooms []interfaces.RoomI, queue interfaces.QueueI) error {
	for _, room := range rooms {
		for _, npc := range room.GetNpcs() {
			npc.SetSpawn(room)
      npc.SetBehavior(queue)
      npc.Spawn()
		}
	}

  return nil
}
