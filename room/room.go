package room

import (
	"github.com/joelevering/gomud/interfaces"
	"github.com/joelevering/gomud/npc"
)

type Room struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Desc    string  `json:"description"`
	Exits   []*Exit `json:"exits"`
	ExitIs  []interfaces.ExitI
	Npcs    []*npc.NPC `json:"npcs"`
	NpcIs   []interfaces.NPCI
	Players []interfaces.PlI
}

type Exit struct {
	Desc   string `json:"description"`
	Key    string `json:"key"`
	RoomID int    `json:"room_id"`
	Room   interfaces.RoomI
}

func (room Room) Message(msg string) {
	for _, player := range room.Players {
		player.SendMsg(msg)
	}
}

func (room *Room) RemovePlayer(player interfaces.PlI, msg string) {
	for i, player := range room.Players {
		if player.GetName() == player.GetName() {
			room.Players[i] = room.Players[len(room.Players)-1]
			room.Players[len(room.Players)-1] = nil
			room.Players = room.Players[:len(room.Players)-1]
      break
		}
	}

	room.Message(msg)
}

func (room *Room) AddPlayer(player interfaces.PlI) {
	room.Message(player.GetName() + " has entered the room!")
	room.Players = append(room.Players, player)
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

func (room *Room) GetPlayers() []interfaces.PlI {
	return room.Players
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
