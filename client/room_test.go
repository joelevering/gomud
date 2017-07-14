package client

import "testing"

func Test_LoadingRooms(t *testing.T) {
	var rooms, err = LoadRooms("rooms.json")
	if err != nil {
		t.Errorf("Error loading rooms: %s", err)
	}

	if len(rooms) == 0 {
		t.Errorf("No rooms loaded")
	}

	var room = rooms[0]

	if room.Id != 1 {
		t.Errorf("Room id expected to be 1 but got %v", room.Id)
	}

	if room.Name != "The Throne Room" {
		t.Errorf("Room name expected to be 'The Throne Room' but got %v", room.Name)
	}

	if room.Desc != "A large hall with a green throne in the center" {
		t.Errorf("Room desc expected to be 'The slime stands alone' but got %v", room.Desc)
	}

	var exit = room.Exits[0]

	if exit.Desc != "(O)ut to the Grand Stairs" {
		t.Errorf("Exit desc expected to be '(O)ut to the Grand Stairs' but got %v", exit.Desc)
	}

	if exit.Key != "o" {
		t.Errorf("Exit key expected to be 'o' but got %v", exit.Key)
	}

	if exit.RoomID != rooms[1].Id {
		t.Errorf("Exit room ID expected to be room %v but got %v", rooms[1].Id, exit.Room.Id)
	}

	if exit.Room.Id != rooms[1].Id {
		t.Errorf("Exit room ID expected to be room %v but got %v", rooms[1].Id, exit.Room.Id)
	}

	var npc = room.Npcs[0]

	if npc.Name != "King Slime" {
		t.Errorf("Expected NPC to have name 'King Slime' but got %v", npc.Name)
	}

	if npc.Desc != "A massive pile of gelatinous goo adorned with two huge eyes" {
		t.Errorf("Expected NPC to have desc 'A massive pile of gelatinous goo adorned with two huge eyes' but got %v", npc.Desc)
	}

	if npc.Health != 999999 || npc.MaxHealth != 999999 {
		t.Errorf("Expected NPC to have health and maxHealth of 999999 but got %d health and %d maxHealth", npc.Health, npc.MaxHealth)
	}

	if npc.Str != 9999 || npc.End != 9999 {
		t.Errorf("Expected NPC to have str and end of 9999 but got %d str and %d end", npc.Str, npc.End)
	}

	if room.Npcs[0].Name != npc.Name {
		t.Error("Expected %v to be in %v", npc.Name, room.Name)
	}
}
