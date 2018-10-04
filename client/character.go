package client

import(
	"fmt"
	"strings"

	"github.com/joelevering/gomud/interfaces"
)

type Character struct {
  Name string
	Room interfaces.RoomI
  Spawn interfaces.RoomI
  Level     int
  Exp       int
  ExpToLvl  int
	MaxHealth int
	Health    int
	Str       int
	End       int
}

func (c *Character) GetName() string {
  return c.Name
}

func (c *Character) Say(msg string) {
	if msg != "" {
		c.Room.Message(fmt.Sprintf("%s says \"%s\"", c.Name, msg))
	}
}

func (c *Character) Yell(msg string) {
	if msg != "" {
		fullMsg := fmt.Sprintf("%s yells \"%s\"", c.Name, msg)
		cli.Room.Message(fullMsg)

		for _, exit := range c.Room.GetExits() {
			exit.GetRoom().Message(fullMsg)
		}
	}
}

func (c *Character) Move(exitKey string) bool {
	for _, exit := range c.Room.GetExits() {
		if strings.ToUpper(exitKey) == strings.ToUpper(exit.GetKey()) {
			c.LeaveRoom(fmt.Sprintf("%s heads to %s!", c.Name, exit.GetRoom().GetName()))
			c.EnterRoom(exit.GetRoom())
			return true
		}
	}

  return false
}

func (c *Character) SetSpawn(room interfaces.RoomI) {
  c.Spawn = room
}

func (c *Character) Spawn() {
  c.EnterRoom(c.Spawn)
}

func (c *Character) EnterRoom(room interfaces.RoomI) {
	room.AddChar(c)
	c.Room = room
}

func (c *Character) LeaveRoom(msg string) {
	if msg == "" {
		msg = fmt.Sprintf("%s has left the room!", c.Name)
	}

	c.Room.RemoveChar(c, msg)
}

func (c *Character) Die(npc interfaces.NPCI) {
  deathNotice := fmt.Sprintf("%s was defeated by %s. Their body dissipates.", c.Name, npc.GetName())
  c.LeaveRoom(deathNotice)

  cli.SendMsg(fmt.Sprintf("You were defeated by %s.", npc.GetName()))
  time.Sleep(1500 * time.Millisecond)
  c.EnterRoom(c.Spawn)
  c.Health = c.MaxHealth
  cli.SendMsg(fmt.Sprintf("You find yourself back in a familiar place: %s", cli.Spawn.GetName()))
  time.Sleep(1500 * time.Millisecond)
  cli.SendMsg("")
  cli.Look()
}
