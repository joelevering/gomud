package main

import (
	"fmt"
	"strings"
	"time"
)

type Client struct {
	Channel   chan<- string
	Name      string
	Room      *Room
	MaxHealth int
	Health    int
	Str       int
	End       int
}

func NewClient(ch chan<- string) *Client {
	return &Client{
		Channel:   ch,
		MaxHealth: 20,
		Health:    20,
		Str:       5,
		End:       5,
	}
}

func (cli Client) List() {
	names := []string{fmt.Sprintf("Yourself (%s)", cli.Name)}

	for _, otherCli := range cli.Room.Clients {
		if otherCli.Name != cli.Name {
			names = append(names, otherCli.Name)
		}
	}

	for _, npc := range cli.Room.Npcs {
		names = append(names, fmt.Sprintf("%s (NPC)", npc.Name))
	}

	cli.SendMsg(fmt.Sprintf("You look around and see: %s", strings.Join(names, ", ")))
}

func (cli Client) Look() {
	cli.SendMsg(fmt.Sprintf("~~%s~~", cli.Room.Name))
	cli.SendMsg(cli.Room.Desc)
	cli.SendMsg("")
	cli.SendMsg("Exits:")
	for _, exit := range cli.Room.Exits {
		cli.SendMsg(fmt.Sprintf("- %s", exit.Desc))
	}

	cli.SendMsg("")
	cli.List()
}

func (cli Client) LookNPC(npcName string) {
	look := func(cli *Client, npc *NPC) {
		cli.SendMsg(fmt.Sprintf("You look at %s and see:", npc.Name))
		cli.SendMsg(npc.Desc)
	}
	cli.findNpcAndExecute(npcName, "Who are you looking at??", look)
}

func (cli Client) AttackNPC(npcName string) {
	attack := func(cli *Client, npc *NPC) {
		cli.SendMsg(fmt.Sprintf("You attack %s!", npc.Name))
		ci := CombatInstance{cli: cli, npc: npc}
		go ci.Start()
	}

	cli.findNpcAndExecute(npcName, "Who are you attacking??", attack)
}

func (cli *Client) findNpcAndExecute(npcName, notFound string, function func(*Client, *NPC)) {
	found := false

	for _, npc := range cli.Room.Npcs {
		if strings.Contains(strings.ToUpper(npc.Name), strings.ToUpper(npcName)) {
			function(cli, &npc)
			found = true
		}
	}

	if found == false {
		cli.SendMsg(notFound)
	}
}

func (cli *Client) Move(exitKey string) {
	for _, exit := range cli.Room.Exits {
		if strings.ToUpper(exitKey) == strings.ToUpper(exit.Key) {
			RemoveClientFromRoom(cli, fmt.Sprintf("%s heads to %s!", cli.Name, exit.Room.Name))
			SetCurrentRoom(cli, exit.Room)
			cli.Look()
			return
		}
	}

	cli.SendMsg("Where are you trying to go??")
}

func (cli Client) Help() {
	cli.SendMsg(HelpMsg)
}

func (cli Client) Say(msg string) {
	if msg != "" {
		cli.Room.Message(fmt.Sprintf("%s says \"%s\"", cli.Name, msg))
	}
}

func (cli Client) Yell(msg string) {
	if msg != "" {
		fullMsg := fmt.Sprintf("%s yells \"%s\"", cli.Name, msg)
		cli.Room.Message(fullMsg)

		for _, exit := range cli.Room.Exits {
			exit.Room.Message(fullMsg)
		}
	}
}

func (cli Client) SendMsg(msg string) {
	stamp := time.Now().Format(time.Kitchen)
	cli.Channel <- fmt.Sprintf("%s %s", stamp, msg)
}
