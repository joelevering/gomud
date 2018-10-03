package room

import (
	"encoding/json"
	"io/ioutil"

	"github.com/joelevering/gomud/interfaces"
)

type Room struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Desc    string  `json:"description"`
	Exits   []*Exit `json:"exits"`
	ExitIs  []interfaces.ExitI
	Npcs    []*NPC `json:"npcs"`
	NpcIs   []interfaces.NPCI
	Clients []interfaces.CliI
}

type Exit struct {
	Desc   string `json:"description"`
	Key    string `json:"key"`
	RoomID int    `json:"room_id"`
	Room   interfaces.RoomI
}

type RoomFinder struct {
	roomMap map[int]int // Room ID to index in []Room
	Rooms   []*Room
}

func newRoomFinder(rooms []*Room) *RoomFinder {
	var roomMap = make(map[int]int, 0)

	for i, rm := range rooms {
		roomMap[rm.GetID()] = i
	}

	return &RoomFinder{
		Rooms:   rooms,
		roomMap: roomMap,
	}
}

func (r *RoomFinder) Find(roomID int) interfaces.RoomI {
	index := r.roomMap[roomID]
	return r.Rooms[index]
}

func LoadRooms(path string) ([]interfaces.RoomI, error) {
	var rooms []*Room

	f, _ := ioutil.ReadFile(path)
	json.Unmarshal(f, &rooms)

  setNPCSpawns(rooms)

	roomFinder := newRoomFinder(rooms)
	attachRoomsToExits(rooms, roomFinder)

	roomIs := []interfaces.RoomI{}
	for _, room := range rooms {
		roomIs = append(roomIs, interfaces.RoomI(room))
	}

	return roomIs, nil
}

func attachRoomsToExits(rooms []*Room, roomFinder *RoomFinder) {
	for _, room := range rooms {
		for _, exit := range room.GetExits() {
			exit.SetRoom(roomFinder.Find(exit.GetRoomID()))
		}
	}
}

func setNPCSpawns(rooms []*Room) {
	for _, room := range rooms {
		for _, npc := range room.GetNpcs() {
			npc.SetSpawn(room)
		}
	}
}

func (room Room) Message(msg string) {
	for _, client := range room.Clients {
		client.SendMsg(msg)
	}
}

func (room *Room) RemoveCli(cli interfaces.CliI, msg string) {
	for i, client := range room.Clients {
		if client.GetName() == cli.GetName() {
			room.Clients[i] = room.Clients[len(room.Clients)-1]
			room.Clients[len(room.Clients)-1] = nil
			room.Clients = room.Clients[:len(room.Clients)-1]
		}
	}

	room.Message(msg)
}

func (room *Room) AddCli(cli interfaces.CliI) {
	room.Message(cli.GetName() + " has entered the room!")
	room.Clients = append(room.Clients, cli)
}

func (room *Room) GetExits() []interfaces.ExitI {
	if room.ExitIs == nil {
		for _, exit := range room.Exits {
			room.ExitIs = append(room.ExitIs, interfaces.ExitI(exit))
		}
	}
	return room.ExitIs
}

func (room *Room) GetNpcs() []interfaces.NPCI {
	if room.NpcIs == nil {
		for _, npc := range room.Npcs {
			room.NpcIs = append(room.NpcIs, interfaces.NPCI(npc))
		}
	}
	return room.NpcIs
}

func (room *Room) GetClients() []interfaces.CliI {
	return room.Clients
}

func (room *Room) GetName() string {
	return room.Name
}

func (room *Room) GetDesc() string {
	return room.Desc
}

func (room *Room) GetID() int {
	return room.Id
}

func (e *Exit) GetRoom() interfaces.RoomI {
	return e.Room
}

func (e *Exit) SetRoom(newRoom interfaces.RoomI) {
	e.Room = newRoom
}

func (e *Exit) GetRoomID() int {
	return e.RoomID
}

func (e *Exit) GetKey() string {
	return e.Key
}

func (e *Exit) GetDesc() string {
	return e.Desc
}
