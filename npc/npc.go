package npc

import (
  "fmt"
  "math/rand"
  "time"

  "github.com/joelevering/gomud/interfaces"
)

type NPC struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"description"`
	MaxHealth int    `json:"max_health"`
	Health    int    `json:"health"`
	Str       int    `json:"strength"`
	End       int    `json:"endurance"`
  Exp       int    `json:"experience"`
  Alive     bool
  SpawnRoom interfaces.RoomI
  Room      interfaces.RoomI
  Behaviors []*Behavior `json:"ooc_behavior"`
}

type Behavior struct {
  Trigger string     `json:"trigger"`
  Actions [][]string `json:"actions"`
  Chance  float64    `json:"chance"`
}

func (n *NPC) GetName() string {
	return n.Name
}

func (n *NPC) GetDesc() string {
	return n.Desc
}

func (n *NPC) GetHealth() int {
	return n.Health
}

func (n *NPC) GetMaxHealth() int {
	return n.MaxHealth
}

func (n *NPC) GetEnd() int {
	return n.End
}

func (n *NPC) GetStr() int {
	return n.Str
}

func (n *NPC) GetExp() int {
	return n.Exp
}

func (n *NPC) SetHealth(newHealth int) {
	n.Health = newHealth
}

func (n *NPC) SetSpawn(room interfaces.RoomI) {
  n.SpawnRoom = room
}

func (n *NPC) Spawn() {
  n.Room = n.SpawnRoom
  n.Alive = true
}

func (n *NPC) Say(msg string) {
  n.Room.Message(fmt.Sprintf("%s says \"%s\"", n.Name, msg))
}

func (n *NPC) Emote(emote string) {
  n.Room.Message(fmt.Sprintf("%s %s", n.Name, emote))
}

func (n *NPC) Die() {
  n.Alive = false

  go func() {
    time.Sleep(10 * time.Second)
    n.SetHealth(n.GetMaxHealth())
    n.Alive = true
  }()
}

func (n *NPC) IsAlive() bool {
  return n.Alive
}

func (n *NPC) SetBehavior(queue interfaces.QueueI) {
	for _, b := range n.Behaviors {
    topic := fmt.Sprintf("%s-%d", b.Trigger, n.Room.GetID())
    ch := queue.Sub(topic)
    go func(n interfaces.NPCI, ch chan bool, chance float64, actions [][]string) {
      for range ch {
        if (rand.Float64() <= chance) {
          time.Sleep(100 * time.Millisecond) // Otherwise triggers on enter happen during room desc
          for _, action := range actions {
            switch action[0] {
            case "say":
              n.Say(action[1])
            case "emote":
              n.Emote(action[1])
            }
          }
        }
      }
    }(n, ch, b.Chance, b.Actions)
  }
}
