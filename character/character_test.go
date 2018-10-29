package character

import (
  "testing"

  "github.com/joelevering/gomud/mocks"
  "github.com/joelevering/gomud/statfx"
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
    Det: 1,
    MaxDet: 100,
    Str: 30,
    Flo: 50,
    Class: class,
  }

  pc.GainExp(1)

  if pc.Exp != 0 {
    t.Errorf("Expected exp to be zeroes on level up, but it's %d", pc.Exp)
  }
  if pc.Level != 2 {
    t.Errorf("Expected PC to hit level 2 after getting NextLvlExp at level 1, but level is %d", pc.Level)
  }
  if pc.Det != pc.MaxDet {
    t.Errorf("Expected to heal fully on level up but detrmination is %d/%d", pc.Det, pc.MaxDet)
  }
  if !(pc.NextLvlExp > 100) {
    t.Error("Expected NextLvlExp to increase on level, but it didn't")
  }
  if pc.MaxDet <= 100 {
    t.Error("Expected leveling up to raise max determination, but it didn't")
  }
  if pc.Str <= 30 || pc.Flo <= 50 {
    t.Error("Expected leveling up to raise max str and flo, but it didn't")
  }
}

func Test_TickFxLowersFxDuration(t *testing.T) {
  ch := NewCharacter()
  fxI := statfx.SEInst{
    Effect: statfx.Weak,
    Duration: 1,
  }
  ch.addFx(fxI)

  ch.TickFx()

  fxRes := ch.Fx[statfx.Weak]
  if fxRes.Duration != 0 {
    t.Errorf("Expected TickFx to lower duration, but duration is %d", fxRes.Duration)
  }
}

func Test_TickFxRemoves0DurationFx(t *testing.T) {
  ch := NewCharacter()
  fxI := statfx.SEInst{
    Effect: statfx.Weak,
    Duration: 0,
  }
  ch.addFx(fxI)

  ch.TickFx()

  fxRes := ch.Fx[statfx.Weak]
  if fxRes != nil {
    t.Error("Expected TickDuration to remove 0-duration fx")
  }
}
