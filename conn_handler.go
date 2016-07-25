package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var HelpMsg = `
Available commands:
'move <exit key>' to move to a new room
'/look' or 'look' to see where you are
'/look <npc name>' or 'look <npc name>' to see more details about an NPC
'/list' or '/ls' to see who is currently in your room
'/help' or 'help' to repeat this message

Anything else will be broadcast as a message to the people in your room
`

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

func RemoveClientFromRoom(cli *Client) {
	oldRoom := cli.room
	oldRoomClients := oldRoom.clients
	for i, client := range oldRoomClients {
		if client == cli {
			oldRoomClients[i] = oldRoomClients[len(oldRoomClients)-1]
			oldRoomClients[len(oldRoomClients)-1] = nil
			cli.room.clients = oldRoomClients[:len(oldRoomClients)-1]
		}
	}

	oldRoom.message(cli.name + " has left the room!")
}

func SetCurrentRoom(cli *Client, room *Room) {
	room.message(cli.name + " has entered the room!")

	cli.room = room
	room.clients = append(room.clients, cli)
}

func handleCommand(cli *Client, cmd string) {
	words := strings.Split(cmd, " ")
	key := words[0]

	switch key {
	case "/list", "/ls":
		cli.List()
	case "/look", "look":
		if len(words) == 1 {
			cli.Look()
		} else if len(words) > 1 {
			cli.LookNPC(words[1])
		}
	case "move":
		cli.Move(words[1])
	case "/help", "help":
		cli.Help()
	default:
		cli.Say(cmd)
	}
}
