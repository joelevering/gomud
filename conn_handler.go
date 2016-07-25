package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var HelpMsg = `
Available commands:
'say <message>' to communicate with people in your room
'move <exit key>' to move to a new room
'look' to see where you are
'look <npc name>' to see more details about an NPC
'list' to see who is currently in your room
'help' to repeat this message

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

func RemoveClientFromRoom(cli *Client, msg string) {
	oldRoom := cli.room
	oldRoomClients := oldRoom.clients
	for i, client := range oldRoomClients {
		if client == cli {
			oldRoomClients[i] = oldRoomClients[len(oldRoomClients)-1]
			oldRoomClients[len(oldRoomClients)-1] = nil
			cli.room.clients = oldRoomClients[:len(oldRoomClients)-1]
		}
	}

	if msg == "" {
		oldRoom.message(cli.name + " has left the room!")
	} else {
		oldRoom.message(msg)
	}
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
	case "":
	case "ls", "list":
		cli.List()
	case "l", "look":
		if len(words) == 1 {
			cli.Look()
		} else if len(words) > 1 {
			cli.LookNPC(words[1])
		}
	case "m", "move":
		cli.Move(words[1])
	case "h", "help":
		cli.Help()
	case "s", "say":
		cli.Say(strings.Join(words[1:], " "))
	default:
		cli.sendMsg("I'm not sure what you mean. Type '/help' for assistance.")
	}
}
