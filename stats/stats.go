package stats

type Stat string

const(
  Det = Stat("det")
  Stm = Stat("stm")
  Foc = Stat("foc")

  Str = Stat("str")
  Flo = Stat("flo")
  Ing = Stat("ing")
  Kno = Stat("kno")
  Sag = Stat("sag")
)

var names = map[Stat]string{
  Det: "determination",
  Stm: "stamina",
  Foc: "focus",

  Str: "strength",
  Flo: "flow",
  Ing: "ingenuity",
  Kno: "knowledge",
  Sag: "sagacity",
}

// Name returns the player-facing display name for a stat, e.g. "stm" -> "stamina".
func (s Stat) Name() string {
  if name, ok := names[s]; ok {
    return name
  }

  return string(s)
}
