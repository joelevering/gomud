package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const Port = "1919"

var entering = make(chan Client)
var leaving = make(chan Client)
var messages = make(chan string)
var clients = make(map[string]Client)

type Client struct {
	channel chan<- string
	name    string
}

func (cli Client) sendMsg(msg string) {
	stamp := time.Now().Format(time.Kitchen)
	cli.channel <- stamp + " " + msg
}

func broadcaster() {
	for {
		select {
		case msg := <-messages:
			for _, cli := range clients {
				cli.sendMsg(msg)
			}
		case cli := <-entering:
			clients[cli.name] = cli
		case cli := <-leaving:
			delete(clients, cli.name)
			close(cli.channel)
		}
	}
}

func handleConn(conn net.Conn) {
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
	listClients(cli)
	messages <- cli.name + " has arrived!"

	input := bufio.NewScanner(conn)
	for input.Scan() {
		handleCommand(cli, input.Text())
	}

	log.Print("User logged out: " + cli.name)
	messages <- cli.name + " has left!"
	leaving <- cli
	conn.Close()
}

func handleCommand(cli Client, cmd string) {
	words := strings.Split(cmd, " ")
	key := words[0]
	// args := words[1:]
	switch key {
	case "/list", "/ls":
		listClients(cli)
	default:
		messages <- cli.name + ": " + cmd
	}
}

func listClients(cli Client) {
	var clientNames []string

	for _, otherCli := range clients {
		clientNames = append(clientNames, otherCli.name)
	}

	cli.sendMsg("Logged in users: " + strings.Join(clientNames, ", "))
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
	host := localIp() + ":" + Port
	log.Print("Hosting on: " + host)
	listener, err := net.Listen("tcp", host)

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
