package main

import "strings"

func (cli Client) List() {
	names := []string{"Yourself (" + cli.name + ")"}

	for _, otherCli := range cli.room.clients {
		if otherCli.name != cli.name {
			names = append(names, otherCli.name)
		}
	}

	for _, npc := range cli.room.npcs {
		names = append(names, npc.name+" (NPC)")
	}

	cli.sendMsg("You look around and see: " + strings.Join(names, ", "))
}

func (cli Client) Look() {
	cli.sendMsg("~~" + cli.room.name + "~~")
	cli.sendMsg(cli.room.desc)
	cli.sendMsg("")
	cli.sendMsg("Exits:")
	for _, exit := range cli.room.exits {
		cli.sendMsg("- " + exit.desc)
	}

	cli.sendMsg("")
	cli.List()
}

func (cli Client) LookNPC(npcName string) {
	found := false

	for _, npc := range cli.room.npcs {
		if strings.Contains(strings.ToUpper(npc.name), strings.ToUpper(npcName)) {
			cli.sendMsg("You look at " + npc.name + " and see:")
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
			RemoveClientFromRoom(cli)
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
