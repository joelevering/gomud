package room

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

	if room.GetID() != 1 {
		t.Errorf("Room id expected to be 1 but got %v", room.GetID())
	}

	if room.GetName() != "The Throne Room" {
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

	var npc = room.GetNpcs()[0]

	if npc.GetName() != "King Slime" {
		t.Errorf("Expected NPC to have name 'King Slime' but got %v", npc.GetName())
	}

	if npc.GetDesc() != "A massive pile of gelatinous goo adorned with two huge eyes" {
		t.Errorf("Expected NPC to have desc 'A massive pile of gelatinous goo adorned with two huge eyes' but got %v", npc.GetDesc())
	}

	if npc.GetHealth() != 999999 || npc.GetMaxHealth() != 999999 {
		t.Errorf("Expected NPC to have health and maxHealth of 999999 but got %d health and %d maxHealth", npc.GetHealth(), npc.GetMaxHealth())
	}

	if npc.GetStr() != 9999 || npc.GetEnd() != 9999 {
		t.Errorf("Expected NPC to have str and end of 9999 but got %d str and %d end", npc.GetStr(), npc.GetEnd())
	}

	if room.GetNpcs()[0].GetName() != npc.GetName() {
		t.Error("Expected %v to be in %v", npc.GetName(), room.GetName())
	}
}
