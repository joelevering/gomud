package color

import "testing"

func Test_DmgIsRed(t *testing.T) {
  if got, want := Dmg("hi"), Red+"hi"+Reset; got != want {
    t.Errorf("Dmg(\"hi\") = %q, want %q", got, want)
  }
}

func Test_HealIsGreen(t *testing.T) {
  if got, want := Heal("hi"), Green+"hi"+Reset; got != want {
    t.Errorf("Heal(\"hi\") = %q, want %q", got, want)
  }
}

func Test_BuffIsGreen(t *testing.T) {
  if got, want := Buff("hi"), Green+"hi"+Reset; got != want {
    t.Errorf("Buff(\"hi\") = %q, want %q", got, want)
  }
}

func Test_DebuffIsYellow(t *testing.T) {
  if got, want := Debuff("hi"), Yellow+"hi"+Reset; got != want {
    t.Errorf("Debuff(\"hi\") = %q, want %q", got, want)
  }
}

func Test_SkillIsCyan(t *testing.T) {
  if got, want := Skill("hi"), Cyan+"hi"+Reset; got != want {
    t.Errorf("Skill(\"hi\") = %q, want %q", got, want)
  }
}
