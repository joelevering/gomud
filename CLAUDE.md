# GoMud

A MUD (text-based multiplayer game) server written in Go.

## Vision docs

Before doing world-building, mechanics, or scope-affecting work, check:

- `docs/vision/world-bible.md` — setting, tone, lore, factions, established world facts
- `docs/vision/mechanics.md` — combat/leveling/economy rules and established numeric conventions
- `docs/vision/hosting-goals.md` — target player count, hosting setup, what "done" looks like

These are kept out of this file on purpose — CLAUDE.md loads on every session, so it stays
short. Read the specific doc(s) above when a task touches that area; don't rely on memory
of them from an earlier session.

## Conventions

- `data/rooms.json` is the live/canonical world file the server reads. `data/rooms.draft.json`
  is a staging file for in-progress area work — see `.claude/skills/develop-area/`. Never write
  world-building output directly to `rooms.json`.
- NPC `class` values must be registered in `classes.ByName` (`classes/classes.go`) before use —
  there's no fallback, so an unregistered class name panics the server on spawn.

## Skills

- `develop-area` — drafts a new world zone's rooms from planning spreadsheets. See
  `.claude/skills/develop-area/SKILL.md`.
