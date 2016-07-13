package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var entering = make(chan Client)
var leaving = make(chan Client)

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	ch <- "Who are you?"
	namer := bufio.NewScanner(conn)
	namer.Scan()
	who := namer.Text()

	cli := Client{channel: ch, name: who}
	clientEnters(cli)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		handleCommand(cli, input.Text())
	}

	clientLeft(cli)
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func ListClients(cli Client) {
	clientNames := []string{"Yourself (" + cli.name + ")"}

	for _, otherCli := range GameState.clients {
		if otherCli != cli {
			clientNames = append(clientNames, otherCli.name)
		}
	}

	cli.sendMsg("You look around and see: " + strings.Join(clientNames, ", "))
}

func handleCommand(cli Client, cmd string) {
	words := strings.Split(cmd, " ")
	key := words[0]
	// args := words[1:]
	switch key {
	case "/list", "/ls":
		ListClients(cli)
	default:
		sendMsg(cli.name + ": " + cmd)
	}
}
