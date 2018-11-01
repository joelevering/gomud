package room

import(
  "encoding/json"
  "io/ioutil"
)

var RoomStore *RoomFinder

type RoomFinder struct {
  RoomMap map[int]int // Room ID to index in []Room
  Rooms   []*Room
  Default *Room
}

func newRoomFinder(rooms []*Room) *RoomFinder {
  var roomMap = make(map[int]int, 0)

  for i, rm := range rooms {
    roomMap[rm.GetID()] = i
  }

  return &RoomFinder{
    Rooms:   rooms,
    RoomMap: roomMap,
  }
}

func (r *RoomFinder) Find(roomID int) *Room {
  index := r.RoomMap[roomID]
  return r.Rooms[index]
}

func LoadRooms(path string) (error) {
  var rooms []*Room

  f, err := ioutil.ReadFile(path)
  if err != nil {
    return err
  }

  err = json.Unmarshal(f, &rooms)
  if err != nil {
    return err
  }

  RoomStore = newRoomFinder(rooms)
  RoomStore.Default = RoomStore.Find(9)
  attachRoomsToExits(rooms)

  return nil
}

func attachRoomsToExits(rooms []*Room) {
  for _, room := range rooms {
    for _, exit := range room.GetExits() {
      exit.SetRoom(RoomStore.Find(exit.GetRoomID()))
    }
  }
}
