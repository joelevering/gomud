package main

import (
  "log"
  "net"

  "github.com/joelevering/gomud/player"
  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/pubsub"
  "github.com/joelevering/gomud/storage"
)

const port = "1919"

type GameState struct {
  Players     map[string]*player.Player
  Rooms       []interfaces.RoomI
  DefaultRoom interfaces.RoomI
  Queue       interfaces.QueueI
  Store       *storage.Storage
}

func main() {
  gameState := initGameState()

  host := localIp() + ":" + port
  log.Print("Hosting on: " + host)
  listener, err := net.Listen("tcp", host)

  if err != nil {
    log.Fatal(err)
  }

  var entering = make(chan *player.Player)
  var leaving = make(chan *player.Player)

  connHandler := ConnHandler{
    entering: entering,
    leaving:  leaving,
    state:    gameState,
  }

  gateKeeper := Gatekeeper{
    entering: entering,
    leaving:  leaving,
    state:    gameState,
  }

  go gateKeeper.KeepTheGate()

  for {
    conn, err := listener.Accept()

    if err != nil {
			log.Print(err)
			continue
    }

    go connHandler.Handle(conn)
  }
}

func initGameState() *GameState {
  var state = GameState{
    Queue: pubsub.NewQueue(),
    Store: storage.LoadStore("data/store.json"),
  }

  state.Players = make(map[string]*player.Player)

  rooms, err := LoadRooms("data/rooms.json")
  if err != nil {
    panic("Error loading rooms")
  }
  err = InitNPs(rooms, state.Queue)
  if err != nil {
    panic("Error loading npcs")
  }

  state.Rooms = rooms
  state.DefaultRoom = state.Rooms[8]

  return &state
}

func localIp() string {
  addrs, _ := net.InterfaceAddrs()
  for _, addr := range addrs {
    ipnet := addr.(*net.IPNet)
    if !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
      return ipnet.IP.String()
    }
  }

  return ""
}
