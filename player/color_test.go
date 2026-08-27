package player

import (
  "strings"
  "testing"

  "github.com/joelevering/gomud/mocks"
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/structs"
)

// One case per category (damage/heal/buff/debuff/skill) for each of
// ReportAtk and ReportDef, confirming the right ANSI wrapper was applied.

func Test_ReportAtkColorsDamageRed(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportAtk(opp, structs.CmbRep{CmbFx: structs.CmbFx{Dmg: 42}})
  res := <-ch

  if !strings.Contains(res, "\x1b[31mmock np name took 42 damage!") || !strings.Contains(res, "\x1b[0m") {
    t.Errorf("Expected damage line wrapped in red, but got '%s'", res)
  }
}

func Test_ReportAtkColorsHealGreen(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportAtk(opp, structs.CmbRep{CmbFx: structs.CmbFx{Heal: 15}})
  res := <-ch

  if !strings.Contains(res, "\x1b[32mYou healed 15 damage!\x1b[0m") {
    t.Errorf("Expected heal line wrapped in green, but got '%s'", res)
  }
}

func Test_ReportAtkColorsBuffGreen(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportAtk(opp, structs.CmbRep{Empowered: true})
  res := <-ch

  if !strings.Contains(res, "\x1b[32mYou dealt increased damage!\x1b[0m") {
    t.Errorf("Expected buff line wrapped in green, but got '%s'", res)
  }
}

func Test_ReportAtkColorsDebuffYellow(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportAtk(opp, structs.CmbRep{Vulnerable: true})
  res := <-ch

  if !strings.Contains(res, "\x1b[33mmock np name is vulnerable. They took increased damage!\x1b[0m") {
    t.Errorf("Expected debuff line wrapped in yellow, but got '%s'", res)
  }
}

func Test_ReportAtkColorsSkillCyan(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportAtk(opp, structs.CmbRep{CmbFx: structs.CmbFx{Skill: skills.Skill{Name: "Shove"}}})
  res := <-ch

  if !strings.Contains(res, "\x1b[36mYou used Shove!\x1b[0m") {
    t.Errorf("Expected skill line wrapped in cyan, but got '%s'", res)
  }
}

func Test_ReportDefColorsDamageRed(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportDef(opp, structs.CmbRep{CmbFx: structs.CmbFx{Dmg: 30}})
  res := <-ch

  if !strings.Contains(res, "\x1b[31mYou were attacked for 30 damage!") || !strings.Contains(res, "\x1b[0m") {
    t.Errorf("Expected damage line wrapped in red, but got '%s'", res)
  }
}

func Test_ReportDefColorsHealGreen(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportDef(opp, structs.CmbRep{CmbFx: structs.CmbFx{Heal: 20}})
  res := <-ch

  if !strings.Contains(res, "\x1b[32mmock np name healed 20 damage!\x1b[0m") {
    t.Errorf("Expected heal line wrapped in green, but got '%s'", res)
  }
}

func Test_ReportDefColorsBuffGreen(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportDef(opp, structs.CmbRep{Steeled: true})
  res := <-ch

  if !strings.Contains(res, "\x1b[32mYou steeled yourself, taking lowered damage!\x1b[0m") {
    t.Errorf("Expected buff line wrapped in green, but got '%s'", res)
  }
}

func Test_ReportDefColorsDebuffYellow(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportDef(opp, structs.CmbRep{Weak: true})
  res := <-ch

  if !strings.Contains(res, "\x1b[33mmock np name dealt lowered damage!\x1b[0m") {
    t.Errorf("Expected debuff line wrapped in yellow, but got '%s'", res)
  }
}

func Test_ReportDefColorsSkillCyan(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportDef(opp, structs.CmbRep{CmbFx: structs.CmbFx{Skill: skills.Skill{Name: "Charge"}}})
  res := <-ch

  if !strings.Contains(res, "\x1b[36mmock np name used Charge!\x1b[0m") {
    t.Errorf("Expected skill line wrapped in cyan, but got '%s'", res)
  }
}

// ReportDef's missed-attack line previously passed opp.GetName() as a
// second, un-formatted argument to SendMsg (which sends each argument as
// its own message rather than Sprintf-ing them) -- it should now send a
// single, correctly formatted, uncolored line.
func Test_ReportDefFormatsMissedAttackCorrectly(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportDef(opp, structs.CmbRep{Missed: true})
  res := <-ch

  if !strings.Contains(res, "mock np name missed their attack!") {
    t.Errorf("Expected 'mock np name missed their attack!', but got '%s'", res)
  }

  if strings.Contains(res, "%s") {
    t.Errorf("Expected the missed-attack message to be formatted, but it still contains a literal '%%s': '%s'", res)
  }

  if strings.Contains(res, "\x1b[") {
    t.Errorf("Expected the missed-attack message to be uncolored, but got '%s'", res)
  }
}
