package main

import (
  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/room"
)

func InitNPs(rooms []*room.Room, queue interfaces.QueueI) error {
  for _, room := range rooms {
    for _, np := range room.GetNPs() {
      np.Init(room, queue)
    }
  }

  return nil
}
