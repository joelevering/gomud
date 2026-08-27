package color

const reset = "\x1b[0m"

func Dmg(s string) string    { return wrap("\x1b[31m", s) }
func Heal(s string) string   { return wrap("\x1b[32m", s) }
func Buff(s string) string   { return wrap("\x1b[32m", s) }
func Debuff(s string) string { return wrap("\x1b[33m", s) }
func Skill(s string) string  { return wrap("\x1b[36m", s) }

func wrap(code, s string) string {
  return code + s + reset
}
