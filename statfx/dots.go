package statfx

type DotType string

type DotFx struct {
  Type        DotType
  Chance      float64
  DmgPct      float64
  DurationMin int
  DurationMax int
}

type DotInst struct {
  Type     DotType
  Dmg      int
  Duration int
}

const(
  Bleed = DotType("bleed")
)
