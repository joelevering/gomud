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
	id      int
	name    string
	desc    string
	exits   []Exit
	clients []*Client
	npcs    []NPC
}

func (room Room) message(msg string) {
	for _, client := range room.clients {
		client.sendMsg(msg)
	}
}

type Exit struct {
	desc string `json:"desc"`
	key  string `json:"key"`
	room *Room  `json:"room,omitempty"`
}

func loadRooms() ([]Room, error) {
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

		if hr.id == 0 {
			log.Fatal(roomLine)
		}

		rooms = append(rooms, hr)
	}

	// Room ID to index in rooms slice
	var roomMap = make(map[int]int, 0)

	for i, rm := range rooms {
		roomMap[rm.id] = i
	}

	loadExits(rooms, roomMap)
	loadNPCs(rooms, roomMap)

	return rooms, nil
}

func hydrateRoom(roomLine []string) Room {
	id, err := strconv.Atoi(roomLine[0])
	if err != nil {
		return Room{id: 0}
	}

	room := Room{
		id:   id,
		name: roomLine[1],
		desc: roomLine[2],
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

		exit := Exit{desc: exitLine[1], key: exitLine[2], room: &rooms[eri]}

		room := &rooms[roomMap[rid]]

		room.exits = append(room.exits, exit)
	}
}
