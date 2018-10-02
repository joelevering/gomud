package room

type NPC struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"description"`
	MaxHealth int    `json:"max_health"`
	Health    int    `json:"health"`
	Str       int    `json:"strength"`
	End       int    `json:"endurance"`
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

func (n *NPC) SetHealth(newHealth int) {
	n.Health = newHealth
}

func (n *NPC) Die() {
  // Leave room
  // go func to respawn
}
