package player

import (
  "strings"
  "testing"

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
