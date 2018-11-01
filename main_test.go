package main

import (
  "testing"

  "github.com/joelevering/gomud/pubsub"
  "github.com/joelevering/gomud/room"
)

func Test_InitializingNPs(t *testing.T) {
  room.LoadRooms("data/rooms.json")
  rooms := room.RoomStore.Rooms
  queue := pubsub.NewQueue()
  err := InitNPs(rooms, queue)
  if err != nil {
    t.Errorf("Error initializing NPs: %s", err)
  }

  room := rooms[0]
  npc := room.GetNPs()[0]

  if npc.GetName() != "King Slime" {
    t.Errorf("Expected NP to have name 'King Slime' but got %v", npc.GetName())
  }

  if npc.GetDesc() != "A massive pile of gelatinous goo adorned with two huge eyes" {
    t.Errorf("Expected NP to have desc 'A massive pile of gelatinous goo adorned with two huge eyes' but got %v", npc.GetDesc())
  }

  if npc.GetDet() != 999990 || npc.GetMaxDet() != 999990 {
    t.Errorf("Expected NP to have determination and max determination of 999990 but got %d det and %d max det", npc.GetDet(), npc.GetMaxDet())
  }

  if npc.GetStr() != 9990 {
    t.Errorf("Expected NP to have str of 9990 but got %d", npc.GetStr())
  }

  if npc.GetAtk() != 9990 {
    t.Errorf("Expected NP to have atk of 9990 but got %d", npc.GetAtk())
  }

  if npc.GetDef() != 9990 {
    t.Errorf("Expected NP to have def of 9990 but got %d", npc.GetDef())
  }

  if room.GetNPs()[0].GetName() != npc.GetName() {
    t.Errorf("Expected %v to be in %v", npc.GetName(), room.GetName())
  }

  if len(queue.Chans["pc-enters-1"]) != 1 {
    t.Error("Expected King Slime to sub to pc-enters-1, but it didn't")
  }

  if len(queue.Chans["pc-leaves-1"]) != 1 {
    t.Error("Expected King Slime to sub to pc-leaves-1, but it didn't")
  }
}
