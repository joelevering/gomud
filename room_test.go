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

	if room.desc != "A large hall with a green throne in the center" {
		t.Errorf("Room desc expected to be 'The slime stands alone' but got %v", room.desc)
	}

	var exit = room.exits[0]

	if exit.desc != "(O)ut to the Grand Stairs" {
		t.Errorf("Exit desc expected to be '(O)ut to the Grand Stairs' but got %v", exit.desc)
	}

	if exit.key != "o" {
		t.Errorf("Exit key expected to be 'o' but got %v", exit.key)
	}

	if exit.room.id != rooms[1].id {
		t.Errorf("Exit room expected to be room with ID %v but got ID %v", rooms[1].id, exit.room.id)
	}

	var npc = room.npcs[0]

	if npc.name != "King Slime" {
		t.Errorf("Expected NPC to have name 'King Slime' but got %v", npc.name)
	}

	if npc.desc != "A massive pile of gelatinous goo adorned with two huge eyes" {
		t.Errorf("Expected NPC to have desc 'A massive pile of gelatinous goo adorned with two huge eyes' but got %v", npc.desc)
	}

	if npc.health != 999999 || npc.maxHealth != 999999 {
		t.Errorf("Expected NPC to have health and maxHealth of 999999 but got %d health and %d maxHealth", npc.health, npc.maxHealth)
	}

	if npc.str != 9999 || npc.end != 9999 {
		t.Errorf("Expected NPC to have str and end of 9999 but got %d str and %d end", npc.str, npc.end)
	}

	if room.npcs[0].name != npc.name {
		t.Error("Expected %v to be in %v", npc.name, room.name)
	}
}
