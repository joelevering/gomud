package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	ch <- "Who are you?"
	namer := bufio.NewScanner(conn)
	namer.Scan()
	who := namer.Text()

	cli := Client{channel: ch, name: who}
	ClientEnters(&cli)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		handleCommand(&cli, input.Text())
	}

	ClientLeft(&cli)
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
		if otherCli.name != cli.name {
			clientNames = append(clientNames, otherCli.name)
		}
	}

	cli.sendMsg("You look around and see: " + strings.Join(clientNames, ", "))
}

func DescribeCurrentRoom(cli Client) {
	cli.sendMsg("~~" + cli.room.name + "~~")
	cli.sendMsg(cli.room.desc)
	cli.sendMsg("")
	cli.sendMsg("Exits:")
	for _, exit := range cli.room.exits {
		cli.sendMsg("- " + exit.desc)
	}

	cli.sendMsg("")
	ListClients(cli)
}

func RemoveClientFromRoom(cli *Client) {
	oldRoomClients := cli.room.clients
	for i, client := range oldRoomClients {
		if client == cli {
			oldRoomClients[i] = oldRoomClients[len(oldRoomClients)-1]
			oldRoomClients[len(oldRoomClients)-1] = nil
			cli.room.clients = oldRoomClients[:len(oldRoomClients)-1]
		}
	}
}

func SetCurrentRoom(cli *Client, room *Room) {
	cli.room = room
	room.clients = append(room.clients, cli)
}

func handleCommand(cli *Client, cmd string) {
	words := strings.Split(cmd, " ")
	key := words[0]

	switch key {
	case "/list", "/ls":
		ListClients(*cli)
	case "/look", "look":
		DescribeCurrentRoom(*cli)
	case "move":
		for _, exit := range cli.room.exits {
			if strings.ToUpper(words[1]) == strings.ToUpper(exit.key) {
				RemoveClientFromRoom(cli)
				SetCurrentRoom(cli, exit.room)
				DescribeCurrentRoom(*cli)
				break
			}
		}
	default:
		sendMsg(cli.name + ": " + cmd)
	}
}
