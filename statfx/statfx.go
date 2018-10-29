package statfx

type StatusEffect string

type SEInst struct {
  Effect   StatusEffect
  Chance   float64
  Duration int
}

const(
  // Negative
  Stun       = StatusEffect("stun")
  Surprise   = StatusEffect("surprise")
  Weak       = StatusEffect("weak")
  Vulnerable = StatusEffect("vulnerable")

  // Positive
  Conserve   = StatusEffect("conserve")
)
