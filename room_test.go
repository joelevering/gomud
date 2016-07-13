package main

import "testing"

func Test_LoadingRooms(t *testing.T) {
	var rooms, err = loadRooms()
	if err != nil {
		t.Errorf("Error loading rooms: %s", err)
	}

	if len(rooms) == 0 {
		t.Errorf("No rooms loaded")
	}

	var room = rooms[0]

	if room.id != 1 {
		t.Errorf("Room id expected to be 1 but got %v", room.id)
	}

	if room.name != "The Throne Room" {
		t.Errorf("Room name expected to be 'The Throne Room' but got %v", room.name)
	}

	if room.desc != "The slime stands alone" {
		t.Errorf("Room desc expected to be 'The slime stands alone' but got %v", room.desc)
	}
}
