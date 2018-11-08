package player

import (
  "strings"
  "testing"

  "github.com/joelevering/gomud/classes"
  "github.com/joelevering/gomud/skills"
)

func Test_HelpSkillSuccess(t *testing.T) {
  msg := HelpSkill("charge")

  if !strings.Contains(msg, skills.Charge.Desc) {
    t.Errorf("Expected HelpSkill message to include skill (charge) description, but got %s", msg)
  }
}

func Test_HelpSkillFailure(t *testing.T) {
  msg := HelpSkill("chareg")

  if !strings.Contains(msg, "can't find that skill") {
    t.Errorf("Expected HelpSkill to return a message about not knowing a skill for a misspelling, but got %s", msg)
  }
}

func Test_HelpClassSuccess(t *testing.T) {
  msg := HelpClass("conscript")

  if !strings.Contains(msg, classes.Conscript.GetDesc()) {
    t.Errorf("Expected HelpClass message to include class (conscript) description, but got %s", msg)
  }
}

func Test_HelpClassFailure(t *testing.T) {
  msg := HelpClass("conscritp")

  if !strings.Contains(msg, "can't find that class") {
    t.Errorf("Expected HelpClass to return a message about not knowing a class for a misspelling, but got %s", msg)
  }
}
