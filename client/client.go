package client

import (
	"fmt"
  "math"
	"net"
	"strings"
	"time"
  "unicode/utf8"

	"github.com/joelevering/gomud/interfaces"
)

type Client struct {
	Channel   chan string
	Name      string
	Room      interfaces.RoomI
  Level     int
  Exp       int
  ExpToLvl  int
	MaxHealth int
	Health    int
	Str       int
	End       int
  Spawn     interfaces.RoomI
  Character *Character
}

func NewClient(ch chan string) *Client {
	return &Client{
		Channel:   ch,
    Character: &{
      Level:     1,
      ExpToLvl:  10,
      MaxHealth: 200,
      Health:    200,
      Str:       20,
      End:       50,
    }
	}
}

func (cli *Client) SetName(name string) {
  cli.Name = name
  cli.Character.Name = name
}

func (cli *Client) SetSpawn(room interfaces.RoomI) {
  cli.char.SetSpawn(room)
}

func (cli *Client) Spawn() {
	cli.char.Spawn()
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
    if npc.IsAlive() {
		  names = append(names, fmt.Sprintf("%s (NPC)", npc.GetName()))
    }
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

func (cli *Client) Status() {
  header := fmt.Sprintf("~~~~~~~~~~*%s*~~~~~~~~~~", cli.Name)
  cli.SendMsg(header)
  cli.SendMsg(fmt.Sprintf("Level: %d", cli.Level))
  cli.SendMsg(fmt.Sprintf("Experience: %d/%d", cli.Exp, cli.ExpToLvl))
  cli.SendMsg("")
  cli.SendMsg(fmt.Sprintf("Health: %d/%d", cli.Health, cli.MaxHealth))
  cli.SendMsg(fmt.Sprintf("Strength: %d", cli.Str))
  cli.SendMsg(fmt.Sprintf("Endurance: %d", cli.End))
  cli.SendMsg(strings.Repeat("~", utf8.RuneCountInString(header)))
}

func (cli *Client) AttackNPC(npcName string) {
	attack := func(cli *Client, npc interfaces.NPCI) {
		cli.SendMsg(fmt.Sprintf("You attack %s!", npc.GetName()))
		ci := CombatInstance{cli: cli, npc: npc}
		go ci.Start()
	}

	cli.findNpcAndExecute(npcName, "Who are you attacking??", attack)
}

func (cli *Client) findNpcAndExecute(npcName, notFound string, function func(*Client, interfaces.NPCI)) {
	for _, npc := range cli.Room.GetNpcs() {
    if npc.IsAlive() && strings.Contains(strings.ToUpper(npc.GetName()), strings.ToUpper(npcName)) {
			function(cli, npc)
			return
		}
	}

	cli.SendMsg(notFound)
}

func (cli *Client) Move(exitKey string) {
  if cli.char.Move(exitKey) {
    // check if can be moved into char
		cli.Look()
  } else {
	  cli.SendMsg("Where are you trying to go??")
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

func (cli *Client) Defeat(npc interfaces.NPCI) {
  cli.Exp += npc.GetExp()

  if cli.Exp >= cli.ExpToLvl {
    cli.SendMsg(fmt.Sprintf("You gained %d experience and leveled up!", npc.GetExp()))
    cli.levelUp()
  } else {
    toLvl := cli.ExpToLvl - cli.Exp
    cli.SendMsg(fmt.Sprintf("You gained %d experience! You need %d more experience to level up.", npc.GetExp(), toLvl))
  }
}

func (cli *Client) GetName() string {
	return cli.Name
}

func (cli *Client) GetRoom() interfaces.RoomI {
	return cli.Room
}

func (cli *Client) levelUp() {
  // Increase stats
  cli.MaxHealth += 25
  cli.Str += 2
  cli.End += 3

  // Level up and carryover EXP
  cli.Level += 1
  cli.Exp = cli.Exp - cli.ExpToLvl

  // Set new EXP to level
  newExpToLvl := float64(cli.ExpToLvl) * 1.25
  cli.ExpToLvl = int(math.Round(newExpToLvl))

  cli.Health = cli.MaxHealth // Heal

  cli.SendMsg(fmt.Sprintf("You're now level %d!", cli.Level))
}
