package main

import (
	"log"
	"net"
)

const port = "1919"

type GameState struct {
	Clients     map[string]*Client
	Rooms       []Room
	DefaultRoom *Room
}

func main() {
	gameState := initGameState()

	host := localIp() + ":" + port
	log.Print("Hosting on: " + host)
	listener, err := net.Listen("tcp", host)

	if err != nil {
		log.Fatal(err)
	}

	var entering = make(chan *Client)
	var leaving = make(chan *Client)

	connHandler := ConnHandler{
		entering: entering,
		leaving:  leaving,
	}

  gateKeeper := Gatekeeper{
    entering: entering,
    leaving: leaving,
    state: gameState,
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
  var state = GameState{}

	state.Clients = make(map[string]*Client)

	var rooms, err = LoadRooms()
	if err != nil {
		panic("Error loading rooms")
	}

	state.Rooms = rooms
	state.DefaultRoom = &state.Rooms[8]

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
