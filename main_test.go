package main

import (
  "testing"

  "github.com/joelevering/gomud/pubsub"
)

func Test_LoadingRooms(t *testing.T) {
  var rooms, err = LoadRooms("data/rooms.json")
  if err != nil {
    t.Errorf("Error loading rooms: %s", err)
  }

  if len(rooms) == 0 {
    t.Errorf("No rooms loaded")
    return
  }

  var room = rooms[0]

  if room.GetID() != 1 {
    t.Errorf("Room id expected to be 1 but got %v", room.GetID())
  }

  if room.GetName() != "Slime Castle - Throne Room" {
    t.Errorf("Room name expected to be 'The Throne Room' but got %v", room.GetName())
  }

  if room.GetDesc() != "A large hall with a green throne in the center" {
    t.Errorf("Room desc expected to be 'The slime stands alone' but got %v", room.GetDesc())
  }

  var exit = room.GetExits()[0]

  if exit.GetDesc() != "(O)ut to the Grand Stairs" {
    t.Errorf("Exit desc expected to be '(O)ut to the Grand Stairs' but got %v", exit.GetDesc())
  }

  if exit.GetKey() != "o" {
    t.Errorf("Exit key expected to be 'o' but got %v", exit.GetKey())
  }

  if exit.GetRoomID() != rooms[1].GetID() {
    t.Errorf("Exit room ID expected to be room %v but got %v", rooms[1].GetID(), exit.GetRoom().GetID())
  }

  if exit.GetRoom().GetID() != rooms[1].GetID() {
    t.Errorf("Exit room ID expected to be room %v but got %v", rooms[1].GetID(), exit.GetRoom().GetID())
  }
}

func Test_InitializingNPs(t *testing.T) {
  rooms, err := LoadRooms("data/rooms.json")
  queue := pubsub.NewQueue()
  err = InitNPs(rooms, queue)
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
