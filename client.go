package main

import (
	"fmt"
	"strings"
	"time"
)

type Client struct {
	channel   chan<- string
	name      string
	room      *Room
	maxHealth int
	health    int
	str       int
	end       int
}

func (cli Client) List() {
	names := []string{fmt.Sprintf("Yourself (%s)", cli.name)}

	for _, otherCli := range cli.room.clients {
		if otherCli.name != cli.name {
			names = append(names, otherCli.name)
		}
	}

	for _, npc := range cli.room.npcs {
		names = append(names, npc.name+" (NPC)")
	}

	cli.sendMsg(fmt.Sprintf("You look around and see: %s", strings.Join(names, ", ")))
}

func (cli Client) Look() {
	cli.sendMsg(fmt.Sprintf("~~%s~~", cli.room.name))
	cli.sendMsg(cli.room.desc)
	cli.sendMsg("")
	cli.sendMsg("Exits:")
	for _, exit := range cli.room.exits {
		cli.sendMsg(fmt.Sprintf("- %s", exit.desc))
	}

	cli.sendMsg("")
	cli.List()
}

func (cli Client) LookNPC(npcName string) {
	found := false

	for _, npc := range cli.room.npcs {
		if strings.Contains(strings.ToUpper(npc.name), strings.ToUpper(npcName)) {
			cli.sendMsg(fmt.Sprintf("You look at %s and see:", npc.name))
			cli.sendMsg(npc.desc)
			found = true
		}
	}

	if found == false {
		cli.sendMsg("Who are you looking at??")
	}
}

func (cli *Client) Move(exitKey string) {
	for _, exit := range cli.room.exits {
		if strings.ToUpper(exitKey) == strings.ToUpper(exit.key) {
			RemoveClientFromRoom(cli, fmt.Sprintf("%s heads to %s!", cli.name, exit.room.name))
			SetCurrentRoom(cli, exit.room)
			cli.Look()
			return
		}
	}

	cli.sendMsg("Where are you trying to go??")
}

func (cli Client) Help() {
	cli.sendMsg(HelpMsg)
}

func (cli Client) Say(msg string) {
	if msg != "" {
		cli.room.message(fmt.Sprintf("%s says \"%s\"", cli.name, msg))
	}
}

func (cli Client) Yell(msg string) {
	if msg != "" {
		fullMsg := fmt.Sprintf("%s yells \"%s\"", cli.name, msg)
		cli.room.message(fullMsg)

		for _, exit := range cli.room.exits {
			exit.room.message(fullMsg)
		}
	}
}

func (cli Client) sendMsg(msg string) {
	stamp := time.Now().Format(time.Kitchen)
	cli.channel <- fmt.Sprintf("%s %s", stamp, msg)
}
