package main

import (
  "log"
  "net"
  "time"

  "github.com/joelevering/gomud/player"
  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/pubsub"
  "github.com/joelevering/gomud/room"
  "github.com/joelevering/gomud/storage"
)

const port = "1919"

type GameState struct {
  Players     map[string]*player.Player
  Queue       interfaces.QueueI
  Store       *storage.Storage
}

func main() {
  config := LoadConfiguration(Config)
  gameState := initGameState(config)

  host := localIp() + ":" + port
  log.Print("Hosting on: " + host)
  listener, err := net.Listen("tcp", host)

  if err != nil {
    log.Fatal(err)
  }

  var entering = make(chan *player.Player)
  var leaving = make(chan *player.Player)
  var claimName = make(chan nameClaimRequest)
  var releaseName = make(chan string)

  connHandler := ConnHandler{
    entering:          entering,
    leaving:           leaving,
    claimName:         claimName,
    releaseName:       releaseName,
    state:             gameState,
    idleWarnAfter:     time.Duration(config.Idle.WarnAfterMinutes) * time.Minute,
    idleWarnInterval:  time.Duration(config.Idle.WarnIntervalMinutes) * time.Minute,
    idleKickAfter:     time.Duration(config.Idle.KickAfterMinutes) * time.Minute,
    idleCheckInterval: time.Duration(config.Idle.CheckIntervalSeconds) * time.Second,
  }

  gateKeeper := Gatekeeper{
    entering:     entering,
    leaving:      leaving,
    claimName:    claimName,
    releaseName:  releaseName,
    state:        gameState,
    pendingNames: make(map[string]bool),
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

func initGameState(config *Configuration) *GameState {
  var state = GameState{
    Queue: pubsub.NewQueue(),
    Store: storage.LoadStore("data/store.json"),
    Players: make(map[string]*player.Player),
  }

  err := room.LoadRooms("data/rooms.json", config.DefaultRoomID)
  if err != nil {
    panic("Error loading rooms")
  }

  err = InitNPs(room.RoomStore.Rooms, state.Queue)
  if err != nil {
    panic("Error loading npcs")
  }

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
