package nonplayer

import (
  "fmt"
  "log"
  "time"

  "github.com/joelevering/gomud/character"
  "github.com/joelevering/gomud/classes"
  "github.com/joelevering/gomud/combat"
  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
  "github.com/joelevering/gomud/stats"
  "github.com/joelevering/gomud/structs"
  "github.com/joelevering/gomud/util"
)

type NonPlayer struct {
  *character.Character   `json:"character"`

  Id           int            `json:"id"`
  Desc         string         `json:"description"`
  ClassName    string         `json:"class"`
  AtkStats     []stats.Stat   `json:"attack_stats"`
  DefStats     []stats.Stat   `json:"defense_stats"`
  Behaviors    []*Behavior    `json:"behavior"`
  CmbBehaviors []*CmbBehavior `json:"combat_behavior"`
  Alive        bool
}

type Behavior struct {
  Trigger string     `json:"trigger"`
  Actions [][]string `json:"actions"`
  Chance  float64    `json:"chance"`
}

type CmbBehavior struct {
  Skill  string  `json:"skill"`
  Chance float64 `json:"chance"`
}

func (n *NonPlayer) Init(room interfaces.RoomI, queue interfaces.QueueI) {
  n.Fx = make(map[statfx.StatusEffect]*statfx.SEInst)
  n.Dots = make(map[statfx.DotType]*statfx.DotInst)

  n.SetSpawn(room)
  n.SetClass()
  n.ResetStats()
  n.Spawn()
  n.SetBehavior(queue)
}

func (n *NonPlayer) SetClass() {
  n.Class = classes.ByName[n.ClassName]
}

func (n *NonPlayer) ResetStats() {
  lvl := n.Level
  exp := n.Exp

  n.Level = 0
  for i := 0; i < lvl; i++ {
    n.LevelUp()
  }

  n.Exp = exp

  if n.GetMaxStm() == 0 {
    n.SetMaxStm(100)
    n.SetStm(100)
  }

  if n.GetMaxFoc() == 0 {
    n.SetMaxFoc(100)
    n.SetFoc(100)
  }
}

func (n *NonPlayer) GetDesc() string {
  return n.Desc
}

func (n *NonPlayer) Spawn() {
  n.Room = n.GetSpawn()
  n.Alive = true
}

func (n *NonPlayer) Say(msg string) {
  n.Room.Message(fmt.Sprintf("%s says \"%s\"", n.GetName(), msg))
}

func (n *NonPlayer) Emote(emote string) {
  n.Room.Message(fmt.Sprintf("%s %s", n.GetName(), emote))
}

func (n *NonPlayer) EnterCombat(opp interfaces.Combatant) {
  n.InCombat = true
  go n.setCmbBehavior()
}

func (n *NonPlayer) setCmbBehavior() {
  for {
    if !n.IsInCombat() {
      break
    }

    for _, cb := range n.CmbBehaviors {
      if (util.RandF() <= cb.Chance) {
        sk := skills.GetSkill(cb.Skill)
        if sk != nil {
          n.SetCmbSkill(sk)
          break
        } else {
          log.Printf("%s - Can't find skill %s", n.GetName(), cb.Skill)
        }
      }
    }

    time.Sleep(combat.TickTime)
  }
}

func (n *NonPlayer) ReportAtk(_ interfaces.Combatant, _ structs.CmbRep) {}

func (n *NonPlayer) ReportDef(_ interfaces.Combatant, _ structs.CmbRep) {}

func (n *NonPlayer) WinCombat(_ interfaces.Combatant) {
  n.LeaveCombat()
  n.FullHeal()
}

func (n *NonPlayer) LoseCombat(_ interfaces.Combatant) {
  n.LeaveCombat()

  n.Alive = false

  go func() {
    time.Sleep(10 * time.Second)
    n.FullHeal()
    n.Alive = true
  }()
}

func (n *NonPlayer) IsAlive() bool {
  return n.Alive
}

func (n *NonPlayer) SetBehavior(queue interfaces.QueueI) {
  for _, b := range n.Behaviors {
    topic := fmt.Sprintf("%s-%d", b.Trigger, n.Room.GetID())
    ch := queue.Sub(topic)
    go func(n interfaces.NPI, ch chan bool, chance float64, actions [][]string) {
      for range ch {
        if (util.RandF() <= chance) {
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
