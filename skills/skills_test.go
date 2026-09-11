package skills

import (
  "testing"

  "github.com/joelevering/gomud/stats"
)

func Test_GetSkill(t *testing.T) {
  sk := GetSkill("charge")

  if sk != Charge {
    t.Errorf("Expected GetSkill with 'charge' to find skill but got %s", sk.Name)
  }
}

func Test_SkillCostString(t *testing.T) {
  sk := &Skill{CostAmt: 10, CostType: stats.Stm}

  if sk.CostString() != "10 stamina" {
    t.Errorf("Expected CostString to be '10 stamina', but got %s", sk.CostString())
  }
}
