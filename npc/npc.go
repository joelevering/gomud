package npc

import (
  "fmt"
  "math/rand"
  "time"

  "github.com/joelevering/gomud/character"
  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/stats"
)

type NPC struct {
  Id        int                  `json:"id"`
  Character *character.Character `json:"character"`
  Desc      string               `json:"description"`
  AtkStats  []stats.Stat         `json:"attack_stats"`
  DefStats  []stats.Stat         `json:"defense_stats"`
  Behaviors []*Behavior          `json:"ooc_behavior"`

  Alive     bool
  SpawnRoom interfaces.RoomI
  Room      interfaces.RoomI
}

type Behavior struct {
  Trigger string     `json:"trigger"`
  Actions [][]string `json:"actions"`
  Chance  float64    `json:"chance"`
}

func (n *NPC) Init(room interfaces.RoomI, queue interfaces.QueueI) {
  n.SetSpawn(room)
  n.Character.SetClass()
  n.Character.ResetStats()
  n.Spawn()
  n.SetBehavior(queue)
}

func (n *NPC) GetCharacter() interfaces.CharI {
  return n.Character
}

func (n *NPC) GetName() string {
  return n.Character.GetName()
}

func (n *NPC) GetDesc() string {
  return n.Desc
}

func (n *NPC) SetSpawn(room interfaces.RoomI) {
  n.SpawnRoom = room
}

func (n *NPC) Spawn() {
  n.Room = n.SpawnRoom
  n.Alive = true
}

func (n *NPC) Say(msg string) {
  n.Room.Message(fmt.Sprintf("%s says \"%s\"", n.GetName(), msg))
}

func (n *NPC) Emote(emote string) {
  n.Room.Message(fmt.Sprintf("%s %s", n.GetName(), emote))
}

func (n *NPC) LoseCombat(winner interfaces.CharI) {
  n.Alive = false

  go func() {
    time.Sleep(10 * time.Second)
    n.Character.Heal()
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
