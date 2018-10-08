package classes

import(
  "testing"

  "github.com/joelevering/gomud/mocks"
)

func Test_LevelUpIncreasesStats(t *testing.T) {
  pc := &mocks.MockCharacter{}
  conscript := &Conscript{}

  conscript.LevelUp(pc)

  if pc.SetEndArg <= pc.GetEnd() || pc.SetStrArg <= pc.GetStr() {
    t.Error("Expected END and STR to increase on level up but they didn't")
  }

  if pc.SetMaxHealthArg <= pc.GetMaxHealth() {
    t.Error("Expceted max health to increase on level up, but it didn't")
  }
}
