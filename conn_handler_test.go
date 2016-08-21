package main

import "testing"

func Test_SetCurrentRoom(t *testing.T) {
	cli := Client{}
	oldRoom := Room{Id: 99, Clients: []*Client{&cli}}
	cli.Room = &oldRoom
	room := Room{Id: 101}

	SetCurrentRoom(&cli, &room)

	if cli.Room.Id != 101 {
		t.Errorf("Expected client room to be set as %d but it was set as %d", room.Id, cli.Room.Id)
	}

	if room.Clients[0] != &cli {
		t.Errorf("Expected client to be the first of the room's clients")
	}

	if len(room.Clients) != 1 {
		t.Errorf("Expected room to only have one client, but it has %d", len(room.Clients))
	}
}

func Test_RemoveClientFromRoom(t *testing.T) {
	cli := Client{}
	oldRoom := Room{Clients: []*Client{&cli}}
	cli.Room = &oldRoom

	RemoveClientFromRoom(&cli, "")

	if len(oldRoom.Clients) != 0 {
		t.Errorf("Expected oldRoom to have no clients, but it has %d", len(oldRoom.Clients))
	}
}
