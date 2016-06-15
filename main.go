package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type Client struct {
  channel chan<- string
  name string
}

const Port = "1919"

var entering = make(chan Client)
var leaving = make(chan Client)

func broadcaster(messages chan string) {
  clients := make(map[string]Client)
	for {
		select {
		case msg := <-messages:
			stamp := time.Now().Format(time.Kitchen)
			for _, cli := range clients {
				cli.channel <- stamp + " " + msg
			}
		case cli := <-entering:
      clients[cli.name] = cli
		case cli := <-leaving:
			delete(clients, cli.name)
      close(cli.channel)
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
  cli := Client{channel: ch, name: who}
	entering <- cli
	log.Print("User logged in: " + cli.name)
	messages <- cli.name + " has arrived!"

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- cli.name + ": " + input.Text()
	}

	log.Print("User logged out: " + cli.name)
	messages <- cli.name + " has left!"
	leaving <- cli
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
