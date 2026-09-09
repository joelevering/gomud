package color

const (
  Reset        = "\x1b[0m"
  Red          = "\x1b[31m"
  Green        = "\x1b[32m"
  Yellow       = "\x1b[33m"
  Blue         = "\x1b[34m"
  Magenta      = "\x1b[35m"
  Cyan         = "\x1b[36m"
  BrightYellow = "\x1b[93m"
)

func Dmg(s string) string       { return wrap(Red, s) }
func Heal(s string) string      { return wrap(Green, s) }
func Buff(s string) string      { return wrap(Green, s) }
func Debuff(s string) string    { return wrap(Yellow, s) }
func Skill(s string) string     { return wrap(Cyan, s) }
func RoomTitle(s string) string { return wrap(Magenta, s) }
func Exit(s string) string      { return wrap(Blue, s) }
func ClassName(s string) string { return wrap(BrightYellow, s) }

func wrap(code, s string) string {
  return code + s + Reset
}
