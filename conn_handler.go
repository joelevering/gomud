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

Most commands have their first letter as a shortcut
`

func HandleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	cli := Client{channel: ch}

	var confirmed string
	var who string

	for strings.ToUpper(confirmed) != "Y" {
		cli.sendMsg("Who are you?")
		input := bufio.NewScanner(conn)
		input.Scan()
		who = input.Text()

		cli.sendMsg(fmt.Sprintf("Are you sure you want to be called \"%s\"? (Y to confirm)", who))
		input.Scan()
		confirmed = input.Text()
	}

	cli.name = who

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

	switch strings.ToLower(key) {
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
		if len(words) > 1 {
			cli.Move(words[1])
		} else {
			cli.sendMsg("Where are you trying to go??")
		}
	case "h", "help":
		cli.Help()
	case "s", "say":
		cli.Say(strings.Join(words[1:], " "))
	case "y", "yell":
		cli.Yell(strings.Join(words[1:], " "))
	default:
		cli.sendMsg("I'm not sure what you mean. Type 'help' for assistance.")
	}
}
