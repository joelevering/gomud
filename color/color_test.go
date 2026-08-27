package color

import "testing"

func Test_Dmg(t *testing.T) {
  if got, want := Dmg("hi"), "\x1b[31mhi\x1b[0m"; got != want {
    t.Errorf("Dmg(\"hi\") = %q, want %q", got, want)
  }
}

func Test_Heal(t *testing.T) {
  if got, want := Heal("hi"), "\x1b[32mhi\x1b[0m"; got != want {
    t.Errorf("Heal(\"hi\") = %q, want %q", got, want)
  }
}

func Test_Buff(t *testing.T) {
  if got, want := Buff("hi"), "\x1b[32mhi\x1b[0m"; got != want {
    t.Errorf("Buff(\"hi\") = %q, want %q", got, want)
  }
}

func Test_Debuff(t *testing.T) {
  if got, want := Debuff("hi"), "\x1b[33mhi\x1b[0m"; got != want {
    t.Errorf("Debuff(\"hi\") = %q, want %q", got, want)
  }
}

func Test_Skill(t *testing.T) {
  if got, want := Skill("hi"), "\x1b[36mhi\x1b[0m"; got != want {
    t.Errorf("Skill(\"hi\") = %q, want %q", got, want)
  }
}
