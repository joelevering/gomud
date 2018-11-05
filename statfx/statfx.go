package statfx

type StatusEffect string

type SEInst struct {
  Effect   StatusEffect
  Chance   float64
  Duration int
}

const(
  // Negative
  Stun          = StatusEffect("stun")
  Surprise      = StatusEffect("surprise")
  Weak          = StatusEffect("weak") // less damage dealt
  Vulnerable    = StatusEffect("vulnerable") // more damage taken
  Concentration = StatusEffect("concentration") // can only attack

  // Positive
  Conserve = StatusEffect("conserve") // reduced stamina consumption
  Dodging  = StatusEffect("dodging") // avoid attacks
  Redirecting = StatusEffect("redirect") // take partial damage and the reduced dmg goes to enemy
  Empowered = StatusEffect("empowered") // more damage dealt
  Steeled = StatusEffect("steeled") // less damage taken
)
