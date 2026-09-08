package room

import(
  "encoding/json"
  "fmt"
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
  if roomID < 0 {
    return nil
  }

  index := r.RoomMap[roomID]
  return r.Rooms[index]
}

func (r *RoomFinder) SetDefault(roomID int) error {
  if _, ok := r.RoomMap[roomID]; !ok {
    return fmt.Errorf("default room %d does not exist", roomID)
  }

  r.Default = r.Find(roomID)
  return nil
}

func LoadRooms(path string, defaultRoomID int) (error) {
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
  if err := RoomStore.SetDefault(defaultRoomID); err != nil {
    return err
  }
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
