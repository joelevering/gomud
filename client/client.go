package client

import (
	"fmt"
	"net"
	"strings"
	"time"
  "unicode/utf8"

	"github.com/joelevering/gomud/interfaces"
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

type Client struct {
	Channel   chan string
  Queue     interfaces.QueueI
	Room      interfaces.RoomI
  Character interfaces.CharI
  CombatCmd []string
}

func NewClient(ch chan string, q interfaces.QueueI) *Client {
	return &Client{
		Channel:   ch,
    Queue:     q,
    Character: NewCharacter(),
	}
}

func (cli Client) StartWriter(conn net.Conn) {
	for msg := range cli.Channel {
		fmt.Fprintln(conn, msg)
	}
}

func (cli *Client) Cmd(cmd string) {
	words := strings.Split(cmd, " ")

  if cli.Character.IsInCombat() {
    cli.SetCombatCmd(words)
    return
  }

	switch strings.ToLower(words[0]) {
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
	case "st", "status":
		cli.Status()
  case "c", "change":
    cli.ChangeClass(words[1])
	default:
		cli.SendMsg("I'm not sure what you mean. Type 'help' for assistance.")
	}
}

func (cli Client) List() {
	names := []string{fmt.Sprintf("Yourself (%s)", cli.GetName())}

	for _, otherCli := range cli.Room.GetClients() {
		if otherCli.GetName() != cli.GetName() {
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
  pc := cli.Character
  header := fmt.Sprintf("~~~~~~~~~~*%s*~~~~~~~~~~", cli.GetName())
  cli.SendMsg(header)
  cli.SendMsg(fmt.Sprintf("Class: %s", pc.GetClassName()))
  cli.SendMsg(fmt.Sprintf("Level: %d", pc.GetLevel()))
  cli.SendMsg(fmt.Sprintf("Experience: %d/%d", pc.GetExp(), pc.GetNextLvlExp()))
  cli.SendMsg("")
  cli.SendMsg(fmt.Sprintf("Determination: %d/%d", pc.GetDet(), pc.GetMaxDet()))
  cli.SendMsg(fmt.Sprintf("Strength: %d", pc.GetStr()))
  cli.SendMsg(fmt.Sprintf("Flow: %d", pc.GetFlo()))
  cli.SendMsg(fmt.Sprintf("Ingenuity: %d", pc.GetIng()))
  cli.SendMsg(fmt.Sprintf("Knowledge: %d", pc.GetKno()))
  cli.SendMsg(fmt.Sprintf("Sagacity: %d", pc.GetSag()))
  cli.SendMsg(strings.Repeat("~", utf8.RuneCountInString(header)))
}

func (cli *Client) AttackNPC(npcName string) {
	attack := func(cli *Client, npc interfaces.NPCI) {
		cli.SendMsg(fmt.Sprintf("You attack %s!", npc.GetName()))
		ci := &CombatInstance{
      cli: cli,
      pc: cli.Character,
      npc: npc,
    }

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
	for _, exit := range cli.Room.GetExits() {
		if strings.ToUpper(exitKey) == strings.ToUpper(exit.GetKey()) {
			cli.LeaveRoom(fmt.Sprintf("%s heads to %s!", cli.GetName(), exit.GetRoom().GetName()))
			cli.EnterRoom(exit.GetRoom())
			cli.Look()
			return
		}
	}

	cli.SendMsg("Where are you trying to go??")
}

func (cli *Client) Say(msg string) {
	if msg != "" {
		cli.Room.Message(fmt.Sprintf("%s says \"%s\"", cli.GetName(), msg))
	}
}

func (cli *Client) Yell(msg string) {
	if msg != "" {
		fullMsg := fmt.Sprintf("%s yells \"%s\"", cli.GetName(), msg)
		cli.Room.Message(fullMsg)

		for _, exit := range cli.Room.GetExits() {
			exit.GetRoom().Message(fullMsg)
		}
	}
}

func (cli *Client) ChangeClass(class string) {
	switch strings.ToLower(class) {
	case "conscript":
    // something
	case "athlete":
    // something
  case "charmer":
    // something
  case "augur":
    // something
  case "sophist":
    // something
  }
}

func (cli *Client) SendMsg(msgs ...string) {
	for _, msg := range msgs {
		stamp := time.Now().Format(time.Kitchen)
		cli.Channel <- fmt.Sprintf("%s %s", stamp, msg)
	}
}

func (cli *Client) LeaveRoom(msg string) {
	if msg == "" {
		msg = fmt.Sprintf("%s has left the room!", cli.GetName())
	}

	cli.Room.RemoveCli(cli, msg)
  cli.Queue.Pub(fmt.Sprintf("pc-leaves-%d", cli.Room.GetID()))
}

func (cli *Client) EnterRoom(room interfaces.RoomI) {
	room.AddCli(cli)
	cli.Room = room
  cli.Queue.Pub(fmt.Sprintf("pc-enters-%d", room.GetID()))
}

func (cli *Client) LoseCombat(npc interfaces.NPCI) {
  pc := cli.Character
  spawn := pc.GetSpawn()

  deathNotice := fmt.Sprintf("%s was defeated by %s. Their body dissipates.", cli.GetName(), npc.GetName())
  cli.LeaveRoom(deathNotice)

  cli.SendMsg(fmt.Sprintf("You were defeated by %s.", npc.GetName()))
  time.Sleep(1500 * time.Millisecond)
  cli.EnterRoom(spawn)
  pc.Heal()
  cli.SendMsg(fmt.Sprintf("You find yourself back in a familiar place: %s", spawn.GetName()))
  time.Sleep(1500 * time.Millisecond)
  cli.SendMsg("")
  cli.Look()
}

func (cli *Client) WinCombat(loser interfaces.NPCI) {
  expGained := loser.GetExp()
  leveledUp := cli.Character.GainExp(expGained)

  if leveledUp {
    cli.SendMsg(fmt.Sprintf("You gained %d experience and leveled up!", expGained))
    cli.SendMsg(fmt.Sprintf("You're now level %d!", cli.Character.GetLevel()))
  } else {
    cli.SendMsg(fmt.Sprintf("You gained %d experience! You need %d more experience to level up.", expGained, cli.Character.ExpToLvl()))
  }
}

func (cli *Client) GetName() string {
  return cli.Character.GetName()
}

func (cli *Client) SetName(name string) {
  cli.Character.SetName(name)
}

func (cli *Client) GetRoom() interfaces.RoomI {
	return cli.Room
}

func (cli *Client) GetCombatCmd() []string {
	return cli.CombatCmd
}

func (cli *Client) SetCombatCmd(cmd []string) {
  cli.CombatCmd = cmd
}

func (cli *Client) Spawn() {
  cli.EnterRoom(cli.Character.GetSpawn())
}
