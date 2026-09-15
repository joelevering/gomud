# Mechanics

Established gameplay-numbers and systems conventions. Anything generating or balancing
content (rooms, NPCs, items) should match these rather than inventing new curves.

> Draft status: **[confirmed]** entries are pulled from `develop-area/SKILL.md`, where they'd
> been established through actual world-building passes but only lived inside one skill's
> instructions. Everything else is a placeholder.

## Leveling & world pacing

- **[confirmed]** Room density: roughly 3–7 rooms per character level covered by a zone; 5
  (the mid-point) is the default sizing target absent other guidance.
- **[confirmed]** Town/hub sub-areas are not level-scaled — flat 5–10 rooms regardless of the
  surrounding zone's level span, since a town's size reflects its role as a hub, not danger.
- **[confirmed]** NPC density: roughly 65–75% of rooms in a built zone contain an NPC; pure
  connector/path rooms are often bare.
- **[confirmed]** NPC level should track perceived power (what its name/class/description
  suggest), not distance from the zone entrance. Escalation within a zone should come from
  introducing tougher creature types deeper in, not uniformly inflating weak creatures' levels.

## Combat & rewards

- **[confirmed]** Baseline `exp_given = level * 5`.
- **[confirmed]** A deliberate elite/mini-boss NPC can carry a 2x–3x multiplier on the
  baseline exp. A zone's capstone boss can go further, roughly 2x–10x, reflecting that bosses
  are meant to be a disproportionate payoff relative to reaching them.
  - Exception: Slime Castle's boss exp numbers are a deliberate joke specific to that area —
    don't use them as a scaling reference for other zones.
- **[confirmed]** Class growth-stat tiering (`Det` = the class's growth stat, alongside a
  primary stat):
  - Weak/common fauna: ~12–18 Det, 3–6 primary stat
  - Moderate: ~20–28 Det, 6–9 primary stat
  - Strong/dungeon-tier: ~28–38 Det, 9–13 primary stat
  - Boss/elite: ~55–70+ Det, 18–25+ primary stat

## Systems not yet built

- **[confirmed]** Conditional/gated exits and attackable-on-sight NPC behavior don't exist yet
  in the engine. Until they do, a boss can only be narratively gated (e.g. placing a guardian
  NPC in the room before it) — not mechanically enforced.

<!-- TODO: Anything else on the roadmap that world-building or mechanics work should design
*toward* even though it can't be enforced yet? -->

## Economy

TBD
<!-- TODO: Currency, item rarity/value curve, crafting — anything established? -->

## Character progression

- Classes all level to 10.
- Characters can switch classes at-will (outside of combat).
- Not implemented: Classes have a tier, starting with tier 1 for the base classes. At least 3 tiers
of classes will exist, if not more.
- Not implemented: Once an additional tier of class is available to the player, they can equip one
class from each tier available to them. For instance, if a player has access to the Tier 1 class
Sophist and then gains access to the Tier 2 class Orator, they can simultaneously equip both classes,
allowing access to both classes skills and stats.
- Not implemented: Classes are unlockable beyond the starting 5, with various hidden requirements to
each.
- Classes cover several non-stereotypical archetypes including Middle Ages influence military, debate,
connivery, etc.

## PvP / social systems

PvP is not curretnly planned. Social systems are low on the priority list.
