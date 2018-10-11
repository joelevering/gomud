package main

import "github.com/joelevering/gomud/interfaces"

func InitNPs(rooms []interfaces.RoomI, queue interfaces.QueueI) error {
	for _, room := range rooms {
		for _, np := range room.GetNPs() {
      np.Init(room, queue)
		}
	}

  return nil
}
