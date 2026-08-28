package player

import (
  "strings"
  "testing"

  "github.com/joelevering/gomud/color"
  "github.com/joelevering/gomud/mocks"
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/structs"
)

// One case per category (damage/heal/buff/debuff/skill) for each of
// ReportAtk and ReportDef, confirming the right color wrapper was applied.

func Test_ReportAtkColorsDamageRed(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportAtk(opp, structs.CmbRep{CmbFx: structs.CmbFx{Dmg: 42}})
  res := <-ch

  want := color.Dmg("mock np name took 42 damage! mock np name has 150/200 health left!")
  if !strings.Contains(res, want) {
    t.Errorf("Expected damage line wrapped in red (%q), but got '%s'", want, res)
  }
}

func Test_ReportAtkColorsHealGreen(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportAtk(opp, structs.CmbRep{CmbFx: structs.CmbFx{Heal: 15}})
  res := <-ch

  want := color.Heal("You healed 15 damage!")
  if !strings.Contains(res, want) {
    t.Errorf("Expected heal line wrapped in green (%q), but got '%s'", want, res)
  }
}

func Test_ReportAtkColorsBuffGreen(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportAtk(opp, structs.CmbRep{Empowered: true})
  res := <-ch

  want := color.Buff("You dealt increased damage!")
  if !strings.Contains(res, want) {
    t.Errorf("Expected buff line wrapped in green (%q), but got '%s'", want, res)
  }
}

func Test_ReportAtkColorsDebuffYellow(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportAtk(opp, structs.CmbRep{Vulnerable: true})
  res := <-ch

  want := color.Debuff("mock np name is vulnerable. They took increased damage!")
  if !strings.Contains(res, want) {
    t.Errorf("Expected debuff line wrapped in yellow (%q), but got '%s'", want, res)
  }
}

func Test_ReportAtkColorsSkillCyan(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportAtk(opp, structs.CmbRep{CmbFx: structs.CmbFx{Skill: skills.Skill{Name: "Shove"}}})
  res := <-ch

  want := color.Skill("You used Shove!")
  if !strings.Contains(res, want) {
    t.Errorf("Expected skill line wrapped in cyan (%q), but got '%s'", want, res)
  }
}

func Test_ReportDefColorsDamageRed(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportDef(opp, structs.CmbRep{CmbFx: structs.CmbFx{Dmg: 30}})
  res := <-ch

  want := color.Dmg("You were attacked for 30 damage! You have 200/200 health left!")
  if !strings.Contains(res, want) {
    t.Errorf("Expected damage line wrapped in red (%q), but got '%s'", want, res)
  }
}

func Test_ReportDefColorsHealGreen(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportDef(opp, structs.CmbRep{CmbFx: structs.CmbFx{Heal: 20}})
  res := <-ch

  want := color.Heal("mock np name healed 20 damage!")
  if !strings.Contains(res, want) {
    t.Errorf("Expected heal line wrapped in green (%q), but got '%s'", want, res)
  }
}

func Test_ReportDefColorsBuffGreen(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportDef(opp, structs.CmbRep{Steeled: true})
  res := <-ch

  want := color.Buff("You steeled yourself, taking lowered damage!")
  if !strings.Contains(res, want) {
    t.Errorf("Expected buff line wrapped in green (%q), but got '%s'", want, res)
  }
}

func Test_ReportDefColorsDebuffYellow(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportDef(opp, structs.CmbRep{Weak: true})
  res := <-ch

  want := color.Debuff("mock np name dealt lowered damage!")
  if !strings.Contains(res, want) {
    t.Errorf("Expected debuff line wrapped in yellow (%q), but got '%s'", want, res)
  }
}

func Test_ReportDefColorsSkillCyan(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  opp := mocks.NewMockNP()

  go p.ReportDef(opp, structs.CmbRep{CmbFx: structs.CmbFx{Skill: skills.Skill{Name: "Charge"}}})
  res := <-ch

  want := color.Skill("mock np name used Charge!")
  if !strings.Contains(res, want) {
    t.Errorf("Expected skill line wrapped in cyan (%q), but got '%s'", want, res)
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

  if strings.Contains(res, color.Reset) {
    t.Errorf("Expected the missed-attack message to be uncolored, but got '%s'", res)
  }
}
