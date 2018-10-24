package statfx

type StatusEffect string

type SEInst struct {
  Effect StatusEffect
  Chance float64
}

const(
  Stun = StatusEffect("stun")
  Surprise = StatusEffect("surprise")
  Conserve = StatusEffect("conserve")
)
