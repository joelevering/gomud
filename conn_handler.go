package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var entering = make(chan Client)
var leaving = make(chan Client)

func handleConn(conn net.Conn, clients map[string]Client) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	ch <- "Who are you?"
	namer := bufio.NewScanner(conn)
	namer.Scan()
	who := namer.Text()

	// who := conn.RemoteAddr().String()
	cli := Client{channel: ch, name: who}
	clientEnters(cli)
	log.Print("User logged in: " + cli.name)
	listClients(cli)
	sendMsg(cli.name + " has arrived!")

	input := bufio.NewScanner(conn)
	for input.Scan() {
		handleCommand(cli, input.Text())
	}

	log.Print("User logged out: " + cli.name)
	sendMsg(cli.name + " has left!")
	clientLeft(cli)
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func listClients(cli Client) {
	var clientNames []string

	for _, otherCli := range clients {
		clientNames = append(clientNames, otherCli.name)
	}

	cli.sendMsg("Logged in users: " + strings.Join(clientNames, ", "))
}

func handleCommand(cli Client, cmd string) {
	words := strings.Split(cmd, " ")
	key := words[0]
	// args := words[1:]
	switch key {
	case "/list", "/ls":
		listClients(cli)
	default:
		sendMsg(cli.name + ": " + cmd)
	}
}
