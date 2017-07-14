package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/joelevering/gomud/client"
)

const helpMsg = `
Available commands:
'say <message>' to communicate with people in your room
'move <exit key>' to move to a new room
'look' to see where you are
'look <npc name>' to see more details about an NPC
'list' to see who is currently in your room
'help' to repeat this message

Most commands have their first letter as a shortcut
`

type ConnHandler struct {
	entering chan *client.Client
	leaving  chan *client.Client
}

func (handler *ConnHandler) Handle(conn net.Conn) {
	defer conn.Close()

	ch := make(chan string)
	cli := client.NewClient(ch)
	go cli.StartWriter(conn)

	cli.Name = confirmName(cli, conn)

	handler.entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		handleCommand(cli, input.Text())
	}

	handler.leaving <- cli
}

func confirmName(cli *client.Client, conn net.Conn) string {
	var confirmed, who string

	for strings.ToUpper(confirmed) != "Y" {
		cli.SendMsg("Who are you?")
		input := bufio.NewScanner(conn)
		input.Scan()
		who = input.Text()

		cli.SendMsg(fmt.Sprintf("Are you sure you want to be called \"%s\"? (Y to confirm)", who))
		input.Scan()
		confirmed = input.Text()
	}

	return who
}

func handleCommand(cli *client.Client, cmd string) {
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
			cli.SendMsg("Where are you trying to go??")
		}
	case "h", "help":
		cli.SendMsg(helpMsg)
	case "s", "say":
		cli.Say(strings.Join(words[1:], " "))
	case "y", "yell":
		cli.Yell(strings.Join(words[1:], " "))
	case "a", "attack":
		cli.AttackNPC(words[1])
	default:
		cli.SendMsg("I'm not sure what you mean. Type 'help' for assistance.")
	}
}
