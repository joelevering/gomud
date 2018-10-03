package room

import (
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
  Dead      bool
  Spawn     interfaces.RoomI
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
  n.Spawn = room
}

func (n *NPC) Die() {
  n.Dead = true

  go func() {
    time.Sleep(10 * time.Second)
    n.SetHealth(n.GetMaxHealth())
    n.Dead = false
  }()
}

func (n *NPC) IsAlive() bool {
  return !n.Dead
}
