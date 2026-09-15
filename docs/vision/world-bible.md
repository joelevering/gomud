# World Manual

The setting, tone, and lore for the GoMud world. This is the source of truth for anything
world-building work (skills, manual area design, item/NPC naming) should stay consistent with.

## Premise

The world of GoMud (to be named in the future) is one of magic, intrigue, peril, and danger.
It is also one of great opportunity. The players are highly-skilled high-growth-potential
characters that have the ability to combine classes and their skills to make interesting
and powerful combinations. The world's challenging enemies guard fantastic treasures that
can modify and enhance the players' abilities in interesting ways, which gives the game a
focus on discovery and creative character-building.

The goal of the game is to provide a sandbox for players to explore. Social play will be
supported but mainly focus on ad-hoc or coordinating teaming up for combat and dungeon
completion, without deep social systems (e.g. guilds).

## Tone

The tone of the world varies by region. Regions can and often should have their own tone,
especially as they increase in danger. A high-level area, for instance, should have a solid
backstory as to why non-player characters would reside or be exploring there. There should be
a realistic explanation for every element of the world including why a region exists, why
characters should travel there, how people survive there (ecosystem, economics-wise, etc.)
and why dungeons exist as well as where their treasure comes from.

The earlier areas are more placid, somewhat more lighthearted (e.g. Slime Forest), but quickly
the areas become more serious/grim, showing how citizens of the world survive and what they
must endure to do so.JvJJ

## Geography

Known built zones so far (from the world-planning spreadsheets and current `rooms.json`):
Slime Forest, Tableland, Seneca Downs, Morass and their associated dungeons.
The world is currently laid out on a 10x10 spatial grid (see the `*gomud-world-grid*.csv` export used by
`develop-area`). With other areas detailed-but-not-built. Specifically: Jungle South, Jungle North,
Mt. Woe, Petro Pacifica, Frosted Steppe, Vatloks' Domain, Saffron Coast, Vermelico Sea, and Isle
Liveo.

The geography should be grounded in reality and always leave the opportunity for expansion, often
alluding to yet-to-be-built areas in the distance.

## World facts (must stay consistent)

- Magic is real, moderately common, but not well-understood. Learning magic takes discipline and
skill, and thus not as common amongst laypeople, serfs, tradeworkers, children, etc.
- The world is inspired by classic fantasy (e.g. Middle Earth, Narnia, Westeros) but does not
directly copy these settings.
- The world varies greatly in danger from geographic region to geographic region (and sometimes
even intra-region). Some areas are relatively peaceful and some are extremely dangerous.
- Technology is approximately Middle Ages, but some sects or factions may have more advanced
technology (e.g. crude firearms).
- As such, metal is relatively rare and expensive. It is an important comodity globally.
Don't casually imply common metal objects or finishes.

## Factions / history / religion

TBD

## Species / creature types

Creatures are innumerable and varied in types.

Class roster currently spans mundane fauna (Vole, Bird, Goat, Lizard, Boar, ...), Slimes,
Constructs/Golems, and Elementals/Spirits (see `classes/classes.go`), but many others exist.

There is a yet-to-be-determined in-world explanation for how these creatures came to be.

## Naming conventions

Names are generally plain descriptors, with possible reference to fictional in-world elements
e.g. dieties, historical persons, factions, creatures, etc.

Naming conventions can vary geographically or by culture.
