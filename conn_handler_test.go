package main

import "testing"

func Test_SetCurrentRoom(t *testing.T) {
	cli := Client{}
	oldRoom := Room{id: 99, clients: []*Client{&cli}}
	cli.room = &oldRoom
	room := Room{id: 101}

	SetCurrentRoom(&cli, &room)

	if cli.room.id != 101 {
		t.Errorf("Expected client room to be set as %d but it was set as %d", room.id, cli.room.id)
	}

	if room.clients[0] != &cli {
		t.Errorf("Expected client to be the first of the room's clients")
	}

	if len(room.clients) != 1 {
		t.Errorf("Expected room to only have one client, but it has %d", len(room.clients))
	}
}

func Test_RemoveClientFromRoom(t *testing.T) {
	cli := Client{}
	oldRoom := Room{clients: []*Client{&cli}}
	cli.room = &oldRoom

	RemoveClientFromRoom(&cli)

	if len(oldRoom.clients) != 0 {
		t.Errorf("Expected oldRoom to have no clients, but it has %d", len(oldRoom.clients))
	}
}
