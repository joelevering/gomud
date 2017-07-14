package main

import (
	"encoding/json"
	"io/ioutil"
)

// use pointers for slices
type Room struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Desc    string  `json:"description"`
	Exits   []*Exit `json:"exits"`
	Npcs    []NPC   `json:"npcs"`
	Clients []*Client
}

type Exit struct {
	Desc   string `json:"description"`
	Key    string `json:"key"`
	RoomID int    `json:"room_id"`
	Room   *Room
}

type NPC struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"description"`
	MaxHealth int    `json:"max_health"`
	Health    int    `json:"health"`
	Str       int    `json:"strength"`
	End       int    `json:"endurance"`
}

type RoomFinder struct {
	roomMap map[int]int // Room ID to index in []Room
	Rooms   []*Room
}

func newRoomFinder(rooms []*Room) *RoomFinder {
	var roomMap = make(map[int]int, 0)

	for i, rm := range rooms {
		roomMap[rm.Id] = i
	}

	return &RoomFinder{
		Rooms:   rooms,
		roomMap: roomMap,
	}
}

func (r *RoomFinder) Find(roomID int) *Room {
	index := r.roomMap[roomID]
	return r.Rooms[index]
}

func LoadRooms() ([]*Room, error) {
	var rooms []*Room

	f, _ := ioutil.ReadFile("rooms.json")
	json.Unmarshal(f, &rooms)

	attachRoomsToExits(rooms)

	return rooms, nil
}

func attachRoomsToExits(rooms []*Room) {
	roomFinder := newRoomFinder(rooms)

	for _, room := range rooms {
		for _, exit := range room.Exits {
			exit.Room = roomFinder.Find(exit.RoomID)
		}
	}
}

func (room Room) Message(msg string) {
	for _, client := range room.Clients {
		client.SendMsg(msg)
	}
}

func (room *Room) RemoveCli(cli *Client, msg string) {
	for i, client := range room.Clients {
		if client == cli {
			room.Clients[i] = room.Clients[len(room.Clients)-1]
			room.Clients[len(room.Clients)-1] = nil
			room.Clients = room.Clients[:len(room.Clients)-1]
		}
	}

	room.Message(msg)
}

func (room *Room) AddCli(cli *Client) {
	room.Message(cli.Name + " has entered the room!")
	room.Clients = append(room.Clients, cli)
}
