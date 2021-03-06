package room

import "testing"

func Test_LoadingRooms(t *testing.T) {
  var err = LoadRooms("../data/rooms.json")
  if err != nil {
    t.Errorf("Error loading rooms: %s", err)
  }

  rooms := RoomStore.Rooms

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
