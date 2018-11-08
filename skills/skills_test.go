package skills

import "testing"

func Test_GetSkill(t *testing.T) {
  sk := GetSkill("charge")

  if sk != Charge {
    t.Errorf("Expected GetSkill with 'charge' to find skill but got %s", sk.Name)
  }
}
