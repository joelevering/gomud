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
	exits   []string
	clients []*Client
}

func loadRooms() ([]Room, error) {
	var rooms = []Room{}

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

		var room = hydrateRoom(roomLine)

		if room.id == 0 {
			log.Fatal(roomLine)
		}

		rooms = append(rooms, hydrateRoom(roomLine))
	}

	return rooms, nil
}

func hydrateRoom(roomLine []string) Room {
	id, err := strconv.Atoi(roomLine[0])
	if err != nil {
		return Room{id: 0}
	}

	return Room{id: id, name: roomLine[1], desc: roomLine[2]}
}
