package room

import (
  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/nonplayer"
)

type Room struct {
  Id      int                    `json:"id"`
  Name    string                 `json:"name"`
  Desc    string                 `json:"description"`
  Exits   []*Exit                `json:"exits"`
  ExitIs  []interfaces.ExitI
  NPs     []*nonplayer.NonPlayer `json:"npcs"`
  NPIs    []interfaces.NPI
  Players []interfaces.PlI
}

func (room Room) Message(msg string) {
  for _, player := range room.Players {
    player.SendMsg(msg)
  }
}

func (room *Room) RemovePlayer(pl interfaces.PlI, msg string) {
  for i, player := range room.Players {
    if player.GetName() == pl.GetName() {
      room.Players[i] = room.Players[len(room.Players)-1]
      room.Players[len(room.Players)-1] = nil
      room.Players = room.Players[:len(room.Players)-1]

      break
    }
  }

  room.Message(msg)
}

func (room *Room) AddPlayer(pl interfaces.PlI) {
  room.Message(pl.GetName() + " has entered the room!")
  room.Players = append(room.Players, pl)
}

func (room *Room) GetExits() []interfaces.ExitI {
  if room.ExitIs == nil {
    for _, exit := range room.Exits {
      room.ExitIs = append(room.ExitIs, interfaces.ExitI(exit))
    }
  }
  return room.ExitIs
}

func (room *Room) GetNPs() []interfaces.NPI {
  if room.NPIs == nil {
    for _, np := range room.NPs {
      room.NPIs = append(room.NPIs, interfaces.NPI(np))
    }
  }
  return room.NPIs
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
