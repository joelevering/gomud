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
var messages = make(chan string)

func broadcaster() {
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

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
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
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}
