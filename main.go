// Log who's joining
// List who's in the room on join
// Command to see users in room
// Allow dynamic setup of IP/port for first person who joins
// Admin login and features (kick people)
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client chan<- string

const Port = "1919"

var entering = make(chan client)
var leaving = make(chan client)

func broadcaster(messages chan string) {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			stamp := time.Now().Format(time.Kitchen)
			for client := range clients {
				client <- stamp + " " + msg
			}
		case client := <-entering:
			clients[client] = true
		case client := <-leaving:
			delete(clients, client)
			close(client)
		}
	}
}

func handleConn(conn net.Conn, messages chan string) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	ch <- "Who are you?"
	namer := bufio.NewScanner(conn)
	namer.Scan()
	who := namer.Text()

	// who := conn.RemoteAddr().String()
	entering <- ch
	log.Print("User logged in: " + who)
	messages <- who + " has arrived!"

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	log.Print("User logged out: " + who)
	messages <- who + " has left!"
	leaving <- ch
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
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
	var messages = make(chan string)

	host := localIp() + ":" + Port
	log.Print("Hosting on: " + host)
	listener, err := net.Listen("tcp", host)

	if err != nil {
		log.Fatal(err)
	}

	go broadcaster(messages)

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn, messages)
	}
}
