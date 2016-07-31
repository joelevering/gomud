package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Room struct {
	Id      int
	Name    string
	Desc    string
	Exits   []Exit
	Clients []*Client
	Npcs    []NPC
}

func (room Room) Message(msg string) {
	for _, client := range room.Clients {
		client.SendMsg(msg)
	}
}

type Exit struct {
	Desc string `json:"desc"`
	Key  string `json:"key"`
	Room *Room  `json:"room,omitempty"`
}

func LoadRooms() ([]Room, error) {
	var rooms []Room

	f, _ := os.Open("rooms.csv")
	r := csv.NewReader(bufio.NewReader(f))

	for {
		roomLine, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var hr = hydrateRoom(roomLine)

		if hr.Id == 0 {
			log.Fatal(roomLine)
		}

		rooms = append(rooms, hr)
	}

	// Room ID to index in rooms slice
	var roomMap = make(map[int]int, 0)

	for i, rm := range rooms {
		roomMap[rm.Id] = i
	}

	loadExits(rooms, roomMap)
	LoadNPCs(rooms, roomMap)

	return rooms, nil
}

func hydrateRoom(roomLine []string) Room {
	id, err := strconv.Atoi(roomLine[0])
	if err != nil {
		return Room{Id: 0}
	}

	room := Room{
		Id:   id,
		Name: roomLine[1],
		Desc: roomLine[2],
	}

	return room
}

func loadExits(rooms []Room, roomMap map[int]int) {
	f, _ := os.Open("exits.csv")
	r := csv.NewReader(bufio.NewReader(f))

	for {
		exitLine, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Get the id of the room the exit connects to
		erid, err := strconv.Atoi(exitLine[3])
		if err != nil {
			log.Fatal(err)
		}

		// Get the ID of the room the exit belongs to
		rid, err := strconv.Atoi(exitLine[0])
		if err != nil {
			log.Fatal(err)
		}

		eri := roomMap[erid] // Connecting room ID

		exit := Exit{Desc: exitLine[1], Key: exitLine[2], Room: &rooms[eri]}

		room := &rooms[roomMap[rid]]

		room.Exits = append(room.Exits, exit)
	}
}
