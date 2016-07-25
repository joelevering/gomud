package main

import (
	"log"
	"net"
)

const port = "1919"

var Entering = make(chan *Client)
var Leaving = make(chan *Client)

var GameState = gameState{}

type gameState struct {
	clients     map[string]*Client
	rooms       []Room
	defaultRoom *Room
}

func initializeGameState() {
	GameState.clients = make(map[string]*Client)

	var rooms, err = loadRooms()
	if err != nil {
		panic("Error loading rooms")
	}

	GameState.rooms = rooms
	GameState.defaultRoom = &GameState.rooms[8]
}

func ClientEnters(cli *Client) {
	Entering <- cli
}

func ClientLeft(cli *Client) {
	Leaving <- cli
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

func main() {
	initializeGameState()

	host := localIp() + ":" + port
	log.Print("Hosting on: " + host)
	listener, err := net.Listen("tcp", host)

	if err != nil {
		log.Fatal(err)
	}

	go Broadcaster()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}

		go HandleConn(conn)
	}
}
