package client

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/joelevering/gomud/interfaces"
)

type Client struct {
	Channel   chan string
	Name      string
	Room      interfaces.RoomI
	MaxHealth int
	Health    int
	Str       int
	End       int
  Spawn     interfaces.RoomI
}

func NewClient(ch chan string) *Client {
	return &Client{
		Channel:   ch,
		MaxHealth: 200,
		Health:    200,
		Str:       20,
		End:       50,
	}
}

func (cli Client) StartWriter(conn net.Conn) {
	for msg := range cli.Channel {
		fmt.Fprintln(conn, msg)
	}
}

func (cli Client) List() {
	names := []string{fmt.Sprintf("Yourself (%s)", cli.Name)}

	for _, otherCli := range cli.Room.GetClients() {
		if otherCli.GetName() != cli.Name {
			names = append(names, otherCli.GetName())
		}
	}

	for _, npc := range cli.Room.GetNpcs() {
		names = append(names, fmt.Sprintf("%s (NPC)", npc.GetName()))
	}

	cli.SendMsg(fmt.Sprintf("You look around and see: %s", strings.Join(names, ", ")))
}

func (cli Client) Look() {
	cli.SendMsg(
		fmt.Sprintf("~~%s~~", cli.Room.GetName()),
		cli.Room.GetDesc(),
		"",
		"Exits:",
	)

	for _, exit := range cli.Room.GetExits() {
		cli.SendMsg(fmt.Sprintf("- %s", exit.GetDesc()))
	}

	cli.SendMsg("")
	cli.List()
}

func (cli Client) LookNPC(npcName string) {
	look := func(cli *Client, npc interfaces.NPCI) {
		cli.SendMsg(
			fmt.Sprintf("You look at %s and see:", npc.GetName()),
			npc.GetDesc(),
		)
	}
	cli.findNpcAndExecute(npcName, "Who are you looking at??", look)
}

func (cli Client) AttackNPC(npcName string) {
	attack := func(cli *Client, npc interfaces.NPCI) {
		cli.SendMsg(fmt.Sprintf("You attack %s!", npc.GetName()))
		ci := CombatInstance{cli: cli, npc: npc}
		go ci.Start()
	}

	cli.findNpcAndExecute(npcName, "Who are you attacking??", attack)
}

func (cli *Client) findNpcAndExecute(npcName, notFound string, function func(*Client, interfaces.NPCI)) {
	for _, npc := range cli.Room.GetNpcs() {
		if strings.Contains(strings.ToUpper(npc.GetName()), strings.ToUpper(npcName)) {
			function(cli, npc)
			return
		}
	}

	cli.SendMsg(notFound)
}

func (cli *Client) Move(exitKey string) {
	for _, exit := range cli.Room.GetExits() {
		if strings.ToUpper(exitKey) == strings.ToUpper(exit.GetKey()) {
			cli.LeaveRoom(fmt.Sprintf("%s heads to %s!", cli.Name, exit.GetRoom().GetName()))
			cli.EnterRoom(exit.GetRoom())
			cli.Look()
			return
		}
	}

	cli.SendMsg("Where are you trying to go??")
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

		for _, exit := range cli.Room.GetExits() {
			exit.GetRoom().Message(fullMsg)
		}
	}
}

func (cli Client) SendMsg(msgs ...string) {
	for _, msg := range msgs {
		stamp := time.Now().Format(time.Kitchen)
		cli.Channel <- fmt.Sprintf("%s %s", stamp, msg)
	}
}

func (cli *Client) LeaveRoom(msg string) {
	if msg == "" {
		msg = fmt.Sprintf("%s has left the room!", cli.Name)
	}

	cli.Room.RemoveCli(cli, msg)
}

func (cli *Client) EnterRoom(room interfaces.RoomI) {
	room.AddCli(cli)
	cli.Room = room
}

func (cli *Client) Die(npc interfaces.NPCI) {
  deathNotice := fmt.Sprintf("%s was defeated by %s. Their body dissipates.", cli.Name, npc.GetName())
  cli.LeaveRoom(deathNotice)

  cli.SendMsg(fmt.Sprintf("You were defeated by %s.", npc.GetName()))
  time.Sleep(1500 * time.Millisecond)
  cli.EnterRoom(cli.Spawn)
  cli.Health = cli.MaxHealth
  cli.SendMsg(fmt.Sprintf("You find yourself back in a familiar place: %s", cli.Spawn.GetName()))
  time.Sleep(1500 * time.Millisecond)
  cli.SendMsg("")
  cli.Look()
}

func (cli *Client) GetName() string {
	return cli.Name
}

func (cli *Client) GetRoom() interfaces.RoomI {
	return cli.Room
}
