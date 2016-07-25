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
	id        int
	name      string
	desc      string
	maxHealth int
	health    int
	str       int
	end       int
}

func loadNPCs(rooms []Room, roomMap map[int]int) {
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
			id:        id,
			name:      npcLine[2],
			desc:      npcLine[3],
			maxHealth: maxHealth,
			health:    health,
			str:       str,
			end:       end,
		}

		if npc.id == 0 {
			log.Fatal(npcLine)
		}

		rid, err := strconv.Atoi(npcLine[1])

		room := &rooms[roomMap[rid]]

		room.npcs = append(room.npcs, npc)
	}
}
