package npc

import (
  "fmt"
  "math/rand"
  "time"

  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/stats"
)

type NPC struct {
	Id        int              `json:"id"`
	Name      string           `json:"name"`
	Desc      string           `json:"description"`
	MaxDet    int              `json:"max_health"`
	Det       int              `json:"health"`
	Str       int              `json:"strength"`
	Flo       int              `json:"flow"`
	Ing       int              `json:"ingenuity"`
	Kno       int              `json:"knowledge"`
	Sag       int              `json:"sagacity"`
  AtkStats  []stats.Stat     `json:"attack_stats"`
  DefStats  []stats.Stat     `json:"defense_stats"`
  Exp       int              `json:"experience"`
  Behaviors []*Behavior      `json:"ooc_behavior"`

  Alive     bool
  SpawnRoom interfaces.RoomI
  Room      interfaces.RoomI
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

func (n *NPC) GetDet() int {
	return n.Det
}

func (n *NPC) GetMaxDet() int {
	return n.MaxDet
}

func (n *NPC) GetStr() int {
	return n.Str
}

func (n *NPC) GetAtk() int {
  atk := 0

  for _, stat := range n.AtkStats {
    switch stat {
    case stats.Str:
      atk += n.Str
    case stats.Flo:
      atk += n.Flo
    case stats.Ing:
      atk += n.Ing
    case stats.Kno:
      atk += n.Kno
    case stats.Sag:
      atk += n.Sag
    }
  }

  return atk
}

func (n *NPC) GetDef() int {
  def := 0

  for _, stat := range n.DefStats {
    switch stat {
    case stats.Str:
      def += n.Str
    case stats.Flo:
      def += n.Flo
    case stats.Ing:
      def += n.Ing
    case stats.Kno:
      def += n.Kno
    case stats.Sag:
      def += n.Sag
    }
  }

  return def
}

func (n *NPC) GetExp() int {
	return n.Exp
}

func (n *NPC) SetDet(newDet int) {
	n.Det = newDet
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

func (n *NPC) LoseCombat(winner interfaces.CharI) {
  n.Alive = false

  go func() {
    time.Sleep(10 * time.Second)
    n.SetDet(n.GetMaxDet())
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
