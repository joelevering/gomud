package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type NPC struct {
	Id        int
	Name      string
	Desc      string
	MaxHealth int
	Health    int
	Str       int
	End       int
}

func LoadNPCs(rooms []Room, roomMap map[int]int) {
	f, _ := os.Open("npcs.csv")
	r := csv.NewReader(bufio.NewReader(f))

	for {
		npcLine, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		id, err := strconv.Atoi(npcLine[0])
		maxHealth, err := strconv.Atoi(npcLine[4])
		health, err := strconv.Atoi(npcLine[5])
		str, err := strconv.Atoi(npcLine[6])
		end, err := strconv.Atoi(npcLine[7])

		npc := NPC{
			Id:        id,
			Name:      npcLine[2],
			Desc:      npcLine[3],
			MaxHealth: maxHealth,
			Health:    health,
			Str:       str,
			End:       end,
		}

		if npc.Id == 0 {
			log.Fatal(npcLine)
		}

		rid, err := strconv.Atoi(npcLine[1])

		room := &rooms[roomMap[rid]]

		room.Npcs = append(room.Npcs, npc)
	}
}
