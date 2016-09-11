package main

import (
	"log"
	"net"
)

const port = "1919"

var entering = make(chan *Client)
var leaving = make(chan *Client)

var GameState = gameState{}

type gameState struct {
	Clients     map[string]*Client
	Rooms       []Room
	DefaultRoom *Room
}

func main() {
	initGameState()

	host := localIp() + ":" + port
	log.Print("Hosting on: " + host)
	listener, err := net.Listen("tcp", host)

	if err != nil {
		log.Fatal(err)
	}

  go Gatekeeper(entering, leaving)

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}

		go HandleConn(conn)
	}
}

func initGameState() {
	GameState.Clients = make(map[string]*Client)

	var rooms, err = LoadRooms()
	if err != nil {
		panic("Error loading rooms")
	}

	GameState.Rooms = rooms
	GameState.DefaultRoom = &GameState.Rooms[8]
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

func ClientEnters(cli *Client) {
	entering <- cli
}

func ClientLeft(cli *Client) {
	leaving <- cli
}
