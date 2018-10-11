package player

import (
	"fmt"
	"net"
	"strings"
	"time"
  "unicode/utf8"

	"github.com/joelevering/gomud/character"
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

type Player struct {
  Channel   chan string
  Queue     interfaces.QueueI
  Room      interfaces.RoomI
  Character *character.Character
  CombatCmd []string
}

func NewPlayer(ch chan string, q interfaces.QueueI) *Player {
  return &Player{
    Channel:   ch,
    Queue:     q,
    Character: character.NewCharacter(),
  }
}

func (p Player) StartWriter(conn net.Conn) {
  for msg := range p.Channel {
		fmt.Fprintln(conn, msg)
  }
}

func (p *Player) Cmd(cmd string) {
  words := strings.Split(cmd, " ")

  if p.Character.IsInCombat() {
    p.SetCombatCmd(words)
    return
  }

  switch strings.ToLower(words[0]) {
  case "":
  case "ls", "list":
		p.List()
  case "l", "look":
		if len(words) == 1 {
			p.Look()
		} else if len(words) > 1 {
			p.LookNP(words[1])
		}
  case "m", "move":
		if len(words) > 1 {
			p.Move(words[1])
		} else {
			p.SendMsg("Where are you trying to go??")
		}
  case "h", "help":
		p.SendMsg(helpMsg)
  case "s", "say":
		p.Say(strings.Join(words[1:], " "))
  case "y", "yell":
		p.Yell(strings.Join(words[1:], " "))
  case "a", "attack":
		p.AttackNP(words[1])
  case "st", "status":
		p.Status()
  case "c", "change":
    p.ChangeClass(words[1])
  default:
		p.SendMsg("I'm not sure what you mean. Type 'help' for assistance.")
  }
}

func (p Player) List() {
  names := []string{fmt.Sprintf("Yourself (%s)", p.GetName())}

  for _, otherP := range p.Room.GetPlayers() {
		if otherP.GetName() != p.GetName() {
			names = append(names, otherP.GetName())
		}
  }

  for _, np := range p.Room.GetNPs() {
    if np.IsAlive() {
		  names = append(names, fmt.Sprintf("%s (NPC)", np.GetName()))
    }
  }

  p.SendMsg(fmt.Sprintf("You look around and see: %s", strings.Join(names, ", ")))
}

func (p Player) Look() {
  p.SendMsg(
		fmt.Sprintf("~~%s~~", p.Room.GetName()),
		p.Room.GetDesc(),
		"",
		"Exits:",
  )

  for _, exit := range p.Room.GetExits() {
		p.SendMsg(fmt.Sprintf("- %s", exit.GetDesc()))
  }

  p.SendMsg("")
  p.List()
}

func (p Player) LookNP(npName string) {
  look := func(p *Player, np interfaces.NPI) {
		p.SendMsg(
			fmt.Sprintf("You look at %s and see:", np.GetName()),
			np.GetDesc(),
		)
  }
  p.findNPAndExecute(npName, "Who are you looking at??", look)
}

func (p *Player) Status() {
  pc := p.Character
  header := fmt.Sprintf("~~~~~~~~~~*%s*~~~~~~~~~~", p.GetName())
  p.SendMsg(header)
  p.SendMsg(fmt.Sprintf("Class: %s", pc.GetClassName()))
  p.SendMsg(fmt.Sprintf("Level: %d", pc.GetLevel()))
  p.SendMsg(fmt.Sprintf("Experience: %d/%d", pc.GetExp(), pc.GetNextLvlExp()))
  p.SendMsg("")
  p.SendMsg(fmt.Sprintf("Determination: %d/%d", pc.GetDet(), pc.GetMaxDet()))
  p.SendMsg(fmt.Sprintf("Stamina: %d/%d", pc.GetStm(), pc.GetMaxStm()))
  p.SendMsg(fmt.Sprintf("Focus: %d/%d", pc.GetFoc(), pc.GetMaxFoc()))
  p.SendMsg("")
  p.SendMsg(fmt.Sprintf("Strength: %d", pc.GetStr()))
  p.SendMsg(fmt.Sprintf("Flow: %d", pc.GetFlo()))
  p.SendMsg(fmt.Sprintf("Ingenuity: %d", pc.GetIng()))
  p.SendMsg(fmt.Sprintf("Knowledge: %d", pc.GetKno()))
  p.SendMsg(fmt.Sprintf("Sagacity: %d", pc.GetSag()))
  p.SendMsg(strings.Repeat("~", utf8.RuneCountInString(header)))
}

func (p *Player) AttackNP(npName string) {
  attack := func(p *Player, np interfaces.NPI) {
		p.SendMsg(fmt.Sprintf("You attack %s!", np.GetName()))
		ci := &CombatInstance{
      p: p,
      np: np,
      pc: p.Character,
      npc: np.GetCharacter(),
    }

		go ci.Start()
  }

  p.findNPAndExecute(npName, "Who are you attacking??", attack)
}

func (p *Player) findNPAndExecute(npName, notFound string, function func(*Player, interfaces.NPI)) {
  for _, np := range p.Room.GetNPs() {
    if np.IsAlive() && strings.Contains(strings.ToUpper(np.GetName()), strings.ToUpper(npName)) {
			function(p, np)
			return
		}
  }

  p.SendMsg(notFound)
}

func (p *Player) Move(exitKey string) {
  for _, exit := range p.Room.GetExits() {
		if strings.ToUpper(exitKey) == strings.ToUpper(exit.GetKey()) {
			p.LeaveRoom(fmt.Sprintf("%s heads to %s!", p.GetName(), exit.GetRoom().GetName()))
			p.EnterRoom(exit.GetRoom())
			p.Look()
			return
		}
  }

  p.SendMsg("Where are you trying to go??")
}

func (p *Player) Say(msg string) {
  if msg != "" {
		p.Room.Message(fmt.Sprintf("%s says \"%s\"", p.GetName(), msg))
  }
}

func (p *Player) Yell(msg string) {
  if msg != "" {
		fullMsg := fmt.Sprintf("%s yells \"%s\"", p.GetName(), msg)
		p.Room.Message(fullMsg)

		for _, exit := range p.Room.GetExits() {
			exit.GetRoom().Message(fullMsg)
		}
  }
}

func (p *Player) ChangeClass(class string) {
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

func (p *Player) SendMsg(msgs ...string) {
  for _, msg := range msgs {
		stamp := time.Now().Format(time.Kitchen)
		p.Channel <- fmt.Sprintf("%s %s", stamp, msg)
  }
}

func (p *Player) LeaveRoom(msg string) {
  if msg == "" {
		msg = fmt.Sprintf("%s has left the room!", p.GetName())
  }

  p.Room.RemovePlayer(p, msg)
  p.Queue.Pub(fmt.Sprintf("pc-leaves-%d", p.Room.GetID()))
}

func (p *Player) EnterRoom(room interfaces.RoomI) {
  room.AddPlayer(p)
  p.Room = room
  p.Queue.Pub(fmt.Sprintf("pc-enters-%d", room.GetID()))
}

func (p *Player) LoseCombat(npc interfaces.CharI) {
  pc := p.Character
  spawn := pc.GetSpawn()

  deathNotice := fmt.Sprintf("%s was defeated by %s. Their body dissipates.", p.GetName(), npc.GetName())
  p.LeaveRoom(deathNotice)

  p.SendMsg(fmt.Sprintf("You were defeated by %s.", npc.GetName()))
  time.Sleep(1500 * time.Millisecond)
  p.EnterRoom(spawn)
  pc.Heal()
  p.SendMsg(fmt.Sprintf("You find yourself back in a familiar place: %s", spawn.GetName()))
  time.Sleep(1500 * time.Millisecond)
  p.SendMsg("")
  p.Look()
}

func (p *Player) WinCombat(loser interfaces.CharI) {
  expGained := loser.GetExpGiven()
  leveledUp := p.Character.GainExp(expGained)

  if leveledUp {
    p.SendMsg(fmt.Sprintf("You gained %d experience and leveled up!", expGained))
    p.SendMsg(fmt.Sprintf("You're now level %d!", p.Character.GetLevel()))
  } else {
    p.SendMsg(fmt.Sprintf("You gained %d experience! You need %d more experience to level up.", expGained, p.Character.ExpToLvl()))
  }
}

func (p *Player) GetName() string {
  return p.Character.GetName()
}

func (p *Player) SetName(name string) {
  p.Character.SetName(name)
}

func (p *Player) GetRoom() interfaces.RoomI {
  return p.Room
}

func (p *Player) GetCombatCmd() []string {
  return p.CombatCmd
}

func (p *Player) SetCombatCmd(cmd []string) {
  p.CombatCmd = cmd
}

func (p *Player) Spawn() {
  p.EnterRoom(p.Character.GetSpawn())
}
