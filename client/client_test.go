package client

import (
	"strings"
	"testing"

	"github.com/joelevering/gomud/interfaces"
	"github.com/joelevering/gomud/room"
)

func Test_EnterRoom(t *testing.T) {
	cli := Client{}
	oldRoom := room.Room{Id: 99, Clients: []interfaces.CliI{&cli}}
	cli.Room = &oldRoom
	room := room.Room{Id: 101}

	cli.EnterRoom(&room)

	if cli.GetRoom().GetID() != 101 {
		t.Errorf("Expected client room to be set as %d but it was set as %d", room.GetID(), cli.GetRoom().GetID())
	}

	if room.Clients[0] != &cli {
		t.Errorf("Expected client to be the first of the room's clients")
	}

	if len(room.Clients) != 1 {
		t.Errorf("Expected room to only have one client, but it has %d", len(room.Clients))
	}
}

func Test_LeaveRoom(t *testing.T) {
	cli := Client{}
	oldRoom := room.Room{Clients: []interfaces.CliI{&cli}}
	cli.Room = &oldRoom

	cli.LeaveRoom("")

	if len(oldRoom.Clients) != 0 {
		t.Errorf("Expected oldRoom to have no clients, but it has %d", len(oldRoom.Clients))
	}
}

func Test_SendMsg(t *testing.T) {
	ch := make(chan string)
	cli := NewClient(ch)

	go cli.SendMsg("testing SendMsg")

	res := <-ch

	if !strings.Contains(res, "testing SendMsg") {
		t.Error("Expected SendMsg to send 'testing SendMsg' to channel, but it didn't")
	}
}

//
// // This test is broken because the Client does not have a Room
// // I want to mock Room but this would require that Client take
// // An interface instead of a room, which would in turn necessitate
// // Getters for all exposed Room attributes and (if making a separate
// // mocks package) somehow getting the Mocks package to recognize
// // structs defined in main (e.g. Exits and other return values required
// // for the mock Room to adhere to the new interface)
// func TestSay(t *testing.T) {
// 	ch := make(chan string)
// 	cli := NewClient(ch)
//
// 	go cli.Say("testing Say")
//
// 	res := <-ch
//
// 	if !strings.Contains(res, "testing Say") {
// 		t.Error("Expected Say to send 'testing Say' to the room, but it didn't")
// 	}
// }
