package main

import(
	"encoding/json"
	"io/ioutil"

	"github.com/joelevering/gomud/interfaces"
	"github.com/joelevering/gomud/room"
)

type RoomFinder struct {
	roomMap map[int]int // Room ID to index in []Room
	Rooms   []*room.Room
}

func newRoomFinder(rooms []*room.Room) *RoomFinder {
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
	var rooms []*room.Room

	f, _ := ioutil.ReadFile(path)
	json.Unmarshal(f, &rooms)

	roomFinder := newRoomFinder(rooms)
	attachRoomsToExits(rooms, roomFinder)

	roomIs := []interfaces.RoomI{}
	for _, room := range rooms {
		roomIs = append(roomIs, interfaces.RoomI(room))
	}

	return roomIs, nil
}

func attachRoomsToExits(rooms []*room.Room, roomFinder *RoomFinder) {
	for _, room := range rooms {
		for _, exit := range room.GetExits() {
			exit.SetRoom(roomFinder.Find(exit.GetRoomID()))
		}
	}
}
