package stats

import "testing"

func Test_StatName(t *testing.T) {
  if Stm.Name() != "stamina" {
    t.Errorf("Expected Stm.Name() to be 'stamina', but got %s", Stm.Name())
  }

  if Foc.Name() != "focus" {
    t.Errorf("Expected Foc.Name() to be 'focus', but got %s", Foc.Name())
  }
}

func Test_StatNameFallsBackToRawValue(t *testing.T) {
  unknown := Stat("xyz")

  if unknown.Name() != "xyz" {
    t.Errorf("Expected unknown Stat.Name() to fall back to its raw value, but got %s", unknown.Name())
  }
}
