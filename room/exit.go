package room

import (
  "github.com/joelevering/gomud/interfaces"
)

type Exit struct {
  Desc   string `json:"description"`
  Key    string `json:"key"`
  RoomID int    `json:"room_id"`
  Room   interfaces.RoomI
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
