package room

import(
  "encoding/json"
  "io/ioutil"
  "os"
  "strconv"
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

func (r *RoomFinder) SetDefault() {
  var defaultID int
  envDefault := os.Getenv("DEFAULT_ROOM_ID")

  if envDefault == "" {
    defaultID = 15
  } else {
    var err error
    defaultID, err = strconv.Atoi(envDefault)
    if err != nil {
      panic("Couldn't load ENV-based default room!")
    }
  }

  r.Default = r.Find(defaultID)
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
  RoomStore.SetDefault()
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
