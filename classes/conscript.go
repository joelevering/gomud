package classes

import "github.com/joelevering/gomud/interfaces"

type Conscript struct {}

func (c Conscript) GetName() string {
  return "Conscript"
}

func (c Conscript) LevelUp(cli interfaces.CliI) {
  // Increase stats
  cli.SetMaxHealth(cli.GetMaxHealth() + 25)
  cli.SetStr(cli.GetStr() + 2)
  cli.SetEnd(cli.GetEnd() + 3)
}
