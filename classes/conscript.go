package classes

import "github.com/joelevering/gomud/interfaces"

type Conscript struct {}

func (c Conscript) GetName() string {
  return "Conscript"
}

func (c Conscript) LevelUp(ch interfaces.CharI) {
  // Increase stats
  ch.SetMaxHealth(ch.GetMaxHealth() + 25)
  ch.SetStr(ch.GetStr() + 2)
  ch.SetEnd(ch.GetEnd() + 3)
}
