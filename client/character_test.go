package client

import (
  "testing"

	"github.com/joelevering/gomud/mocks"
)

func Test_GainExpIncreasesExp(t *testing.T) {
  pc := &Character{
    Level: 1,
    Exp: 0,
    NextLvlExp: 100,
  }

  pc.GainExp(10)

  if pc.Exp != 10 {
    t.Errorf("Expected GainExp to give PC 10 exp, but it gave it %d", pc.Exp)
  }

  if pc.Level != 1 {
    t.Errorf("Expected GainExp to not level up PC that didn't hit NextLvlExp, but it leveled to %d", pc.Level)
  }
}

func Test_GainExpLevelsUp(t *testing.T) {
  class := &mocks.MockClass{}
  pc := &Character{
    Level: 1,
    Exp: 99,
    NextLvlExp: 100,
    Health: 1,
    MaxHealth: 100,
    Str: 30,
    End: 50,
    Class: class,
  }

  pc.GainExp(1)

  if pc.Exp != 0 {
    t.Errorf("Expected exp to be zeroes on level up, but it's %d", pc.Exp)
  }
  if pc.Level != 2 {
    t.Errorf("Expected PC to hit level 2 after getting NextLvlExp at level 1, but level is %d", pc.Level)
  }
  if pc.Health != pc.MaxHealth {
    t.Errorf("Expected to heal fully on level up but health is %d/%d", pc.Health, pc.MaxHealth)
  }
  if !(pc.NextLvlExp > 100) {
    t.Error("Expected NextLvlExp to increase on level, but it didn't")
  }
  if pc.MaxHealth <= 100 {
    t.Error("Expected leveling up to raise max health, but it didn't")
  }
  if pc.Str <= 30 || pc.End <= 50 {
    t.Error("Expected leveling up to raise max str and end, but it didn't")
  }
}
