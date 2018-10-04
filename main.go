package main

import (
	"log"
	"net"

	"github.com/joelevering/gomud/client"
	"github.com/joelevering/gomud/interfaces"
	"github.com/joelevering/gomud/pubsub"
)

const port = "1919"

type GameState struct {
	Clients     map[string]*client.Client
	Rooms       []interfaces.RoomI
	DefaultRoom interfaces.RoomI
  Queue       interfaces.QueueI
}

func main() {
	gameState := initGameState()

	host := localIp() + ":" + port
	log.Print("Hosting on: " + host)
	listener, err := net.Listen("tcp", host)

	if err != nil {
		log.Fatal(err)
	}

	var entering = make(chan *client.Client)
	var leaving = make(chan *client.Client)

	connHandler := ConnHandler{
		entering: entering,
		leaving:  leaving,
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
  }

	state.Clients = make(map[string]*client.Client)

  rooms, err := LoadRooms("data/rooms.json")
	if err != nil {
		panic("Error loading rooms")
	}
  err = InitNPCs(rooms, state.Queue)
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
