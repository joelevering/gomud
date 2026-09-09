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
  FleetFooted = StatusEffect("fleetFooted") // better chance to flee combat
)

// SelfManaged effects age at the point they're actually consumed (see
// Character.ResistAtk for Vulnerable) rather than on Character.TickFx's
// generic per-turn sweep. They're excluded there to avoid double-aging:
// Vulnerable is only ever relevant on the turn its carrier is attacked,
// which can fall on a different turn than the one TickFx runs on for that
// carrier -- ticking it on both would sometimes expire it before it's
// ever used.
var SelfManaged = map[StatusEffect]bool{
  Vulnerable: true,
}
