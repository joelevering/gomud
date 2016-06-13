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
)

type client chan<- string

var entering = make(chan client)
var leaving = make(chan client)

func broadcaster(messages chan string) {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for client := range clients {
				client <- msg
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
	messages <- who + " has arrived!"

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	messages <- who + " has left!"
	leaving <- ch
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	var messages = make(chan string)

	listener, err := net.Listen("tcp", "1.1.1.1:1")

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
