package client

import (
  "math"

	"github.com/joelevering/gomud/classes"
	"github.com/joelevering/gomud/interfaces"
)

type Character struct {
	Name       string
  Class      interfaces.ClassI
  Level      int
  Exp        int
  NextLvlExp int
	MaxHealth  int
	Health     int
	Str        int
	End        int
  InCombat   bool
  Spawn      interfaces.RoomI
}

func NewCharacter() *Character {
  return &Character{
    Class:      classes.Conscript{},
    Level:      1,
    NextLvlExp: 10,
    MaxHealth:  200,
    Health:     200,
    Str:        20,
    End:        50,
  }
}

func (ch *Character) GetName() string {
  return ch.Name
}

func (ch *Character) SetName(name string) {
  ch.Name = name
}

func (ch *Character) GetLevel() int {
  return ch.Level
}

func (ch *Character) GetExp() int {
  return ch.Exp
}

func (ch *Character) GetNextLvlExp() int {
  return ch.NextLvlExp
}

func (ch *Character) GetHealth() int {
  return ch.Health
}

func (ch *Character) SetHealth(health int) {
  ch.Health = health
}

func (ch *Character) GetMaxHealth() int {
  return ch.MaxHealth
}

func (ch *Character) SetMaxHealth(maxHealth int) {
  ch.MaxHealth = maxHealth
}

func (ch *Character) GetStr() int {
  return ch.Str
}

func (ch *Character) SetStr(str int) {
  ch.Str = str
}

func (ch *Character) GetEnd() int {
  return ch.End
}

func (ch *Character) SetEnd(end int) {
  ch.End = end
}

func (ch *Character) GetSpawn() interfaces.RoomI {
  return ch.Spawn
}

func (ch *Character) SetSpawn(spawn interfaces.RoomI) {
  ch.Spawn = spawn
}

func (ch *Character) Heal() {
  ch.Health = ch.MaxHealth
}

func (ch *Character) EnterCombat() {
  ch.InCombat = true
}

func (ch *Character) LeaveCombat() {
  ch.InCombat = false
}

func (ch *Character) IsInCombat() bool {
  return ch.InCombat
}

func (ch *Character) GainExp(exp int) (leveledUp bool) {
  ch.Exp += exp

  if ch.Exp >= ch.NextLvlExp {
    ch.levelUp()
    return true
  }

  return false
}

func (ch *Character) levelUp() {
  // Increase stats based on Class
  statGrowth := ch.Class.GetStatGrowth()
  ch.SetMaxHealth(ch.MaxHealth + statGrowth.Health)
  ch.SetStr(ch.Str + statGrowth.Str)
  ch.SetEnd(ch.End + statGrowth.End)

  // Level up and carryover EXP
  ch.Level += 1
  ch.Exp = ch.Exp - ch.NextLvlExp

  // Set new EXP to level
  newNextLvlExp := float64(ch.NextLvlExp) * 1.25
  ch.NextLvlExp = int(math.Round(newNextLvlExp))

  ch.Heal()
}

func (ch *Character) ExpToLvl() int {
  return ch.NextLvlExp - ch.Exp
}
