package player

import (
  "strings"
  "testing"

  "github.com/joelevering/gomud/character"
  "github.com/joelevering/gomud/classes"
  "github.com/joelevering/gomud/mocks"
  "github.com/joelevering/gomud/room"
  "github.com/joelevering/gomud/storage"
)

func Test_NewPlayerHasDefaultAliases(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)

  if p.Aliases["l"] != "look" {
    t.Errorf("Expected default alias 'l' to expand to 'look', but got '%s'", p.Aliases["l"])
  }

  if p.Aliases["a"] != "attack" {
    t.Errorf("Expected default alias 'a' to expand to 'attack', but got '%s'", p.Aliases["a"])
  }
}

func Test_SetAlias(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  p.Store.InitPlayerData(p.GetID())

  go p.SetAlias("gs", "attack slime shove")
  res := <-ch

  if !strings.Contains(res, "Alias set: 'gs' now runs 'attack slime shove'") {
    t.Errorf("Expected confirmation message, but got '%s'", res)
  }

  if p.Aliases["gs"] != "attack slime shove" {
    t.Errorf("Expected p.Aliases['gs'] to be set, but got '%s'", p.Aliases["gs"])
  }

  if p.Store.LoadAliases(p.GetID())["gs"] != "attack slime shove" {
    t.Error("Expected alias to be persisted immediately, but it wasn't")
  }
}

func Test_SetAliasReservedCollision(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  p.Store.InitPlayerData(p.GetID())

  go p.SetAlias("look", "say hi")
  res := <-ch

  if !strings.Contains(res, "'look' is a reserved command and can't be used as an alias name.") {
    t.Errorf("Expected reserved command rejection, but got '%s'", res)
  }

  if _, ok := p.Aliases["look"]; ok {
    t.Error("Expected 'look' to not be added as an alias, but it was")
  }
}

func Test_SetAliasOverwritesDefault(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  p.Store.InitPlayerData(p.GetID())

  go p.SetAlias("l", "look closely")
  <-ch

  if p.Aliases["l"] != "look closely" {
    t.Errorf("Expected default alias 'l' to be reassignable, but got '%s'", p.Aliases["l"])
  }
}

func Test_RemoveAlias(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  p.Store.InitPlayerData(p.GetID())

  go p.RemoveAlias("l")
  res := <-ch

  if !strings.Contains(res, "Alias 'l' removed.") {
    t.Errorf("Expected removal confirmation, but got '%s'", res)
  }

  if _, ok := p.Aliases["l"]; ok {
    t.Error("Expected 'l' to be removed from aliases, but it wasn't")
  }

  if _, ok := p.Store.LoadAliases(p.GetID())["l"]; ok {
    t.Error("Expected removal to be persisted, but 'l' is still in storage")
  }
}

func Test_RemoveAliasNotFound(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  p.Store.InitPlayerData(p.GetID())

  go p.RemoveAlias("nope")
  res := <-ch

  if !strings.Contains(res, "You don't have an alias named 'nope'.") {
    t.Errorf("Expected not-found message, but got '%s'", res)
  }
}

func Test_ListAliasesEmpty(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  p.Aliases = map[string]string{}

  go p.ListAliases()
  res := <-ch

  if !strings.Contains(res, "You don't have any aliases set.") {
    t.Errorf("Expected empty-aliases message, but got '%s'", res)
  }
}

func Test_ListAliasesWithEntries(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)

  go p.ListAliases()

  header := <-ch
  if !strings.Contains(header, "Your aliases:") {
    t.Errorf("Expected 'Your aliases:' header, but got '%s'", header)
  }

  lines := make([]string, 0, len(p.Aliases))
  for i := 0; i < len(p.Aliases); i++ {
    lines = append(lines, <-ch)
  }

  joined := strings.Join(lines, "\n")
  if !strings.Contains(joined, "l -> look") {
    t.Errorf("Expected alias list to contain 'l -> look', but got: %s", joined)
  }
  if !strings.Contains(joined, "a -> attack") {
    t.Errorf("Expected alias list to contain 'a -> attack', but got: %s", joined)
  }
}

func Test_ExpandAliasNoMatch(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)

  res := p.expandAlias("bogus command")
  if res != "bogus command" {
    t.Errorf("Expected unmatched input to pass through unchanged, but got '%s'", res)
  }
}

func Test_ExpandAliasExactMatch(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)

  res := p.expandAlias("l")
  if res != "look" {
    t.Errorf("Expected 'l' to expand to 'look', but got '%s'", res)
  }
}

func Test_ExpandAliasWithExtraArgs(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  p.Aliases["gs"] = "attack slime"

  res := p.expandAlias("gs shove")
  if res != "attack slime shove" {
    t.Errorf("Expected extra args to be appended to expansion, but got '%s'", res)
  }
}

func Test_CmdExpandsAliasOutOfCombat(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  room := &mocks.MockRoom{Name: "Aliased Room"}
  p.Room = room

  go p.Cmd("l")

  res := <-ch
  if !strings.Contains(res, "~~Aliased Room~~") {
    t.Errorf("Expected 'l' alias to trigger Look(), but got '%s'", res)
  }

  // Drain the rest of Look()'s output (desc, exits header, blank, list) so
  // the goroutine doesn't block sending on a channel this test then closes.
  for i := 0; i < 5; i++ {
    <-ch
  }
}

func Test_CmdExpandsAliasInCombat(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  p.Class = classes.Minder
  p.Classes[classes.Tier2] = classes.Minder
  p.Level = character.MaxLevel
  p.Aliases["gs"] = "shove"
  defer close(ch)
  go p.EnterCombat(&mocks.MockNP{})
  <-ch // "You attack %s!"

  go p.Cmd("gs")
  res := <-ch

  if !strings.Contains(res, "Preparing Shove") {
    t.Errorf("Expected 'Preparing Shove', but got %s", res)
  }

  sk := p.CmbSkill
  if sk == nil || sk.Name != "Shove" {
    t.Errorf("Expected combat skill 'Shove' after aliased command 'gs', but got %v", sk)
  }
}

func Test_AliasPersistsAcrossReload(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  name := "alias-persist-test"
  p.SetName(name)
  p.Init()

  go p.SetAlias("gs", "attack slime shove")
  <-ch // confirmation message

  ch2 := make(chan string)
  p2 := NewPlayer(ch2, p.Queue, p.Store)
  defer close(ch2)
  p2.SetName(name)
  p2.SetSpawn(&room.Room{})
  p2.Init()

  if p2.Aliases["gs"] != "attack slime shove" {
    t.Errorf("Expected alias to persist across reload, but got '%s'", p2.Aliases["gs"])
  }
}

func Test_LoadAliasesFallsBackToDefaultsForLegacySave(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  name := "legacy-player"
  p.SetName(name)

  // Simulate a save file that predates the aliases feature: no "aliases" key at all.
  s := p.Store.(*storage.Storage)
  s.PlayersData[name] = &storage.PlayerData{
    Character: storage.CharStats{Room: -1, Spawn: -1},
  }

  p.Init()

  if p.Aliases["l"] != "look" {
    t.Errorf("Expected legacy save with no aliases to fall back to defaults, but got '%s'", p.Aliases["l"])
  }
}

// Cmd() wiring for the 'alias'/'unalias' commands themselves (argument-count
// branches and end-to-end dispatch), as opposed to the SetAlias/RemoveAlias/
// ListAliases methods tested directly above.

func Test_CmdSetsAliasViaAliasCommand(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  p.Store.InitPlayerData(p.GetID())

  go p.Cmd("alias gs attack slime shove")
  res := <-ch

  if !strings.Contains(res, "Alias set: 'gs' now runs 'attack slime shove'") {
    t.Errorf("Expected confirmation message, but got '%s'", res)
  }

  if p.Aliases["gs"] != "attack slime shove" {
    t.Errorf("Expected 'gs' alias to be set via Cmd(), but got '%s'", p.Aliases["gs"])
  }
}

func Test_CmdListsAliasesViaAliasCommand(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)

  go p.Cmd("alias")

  header := <-ch
  if !strings.Contains(header, "Your aliases:") {
    t.Errorf("Expected 'Your aliases:' header via Cmd(), but got '%s'", header)
  }

  for i := 0; i < len(p.Aliases); i++ {
    <-ch
  }
}

func Test_CmdAliasTooFewArgs(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)

  go p.Cmd("alias x")
  res := <-ch

  if !strings.Contains(res, "Use 'alias <name> <command>' to create an alias") {
    t.Errorf("Expected usage message, but got '%s'", res)
  }

  if _, ok := p.Aliases["x"]; ok {
    t.Error("Expected malformed 'alias' command to not create an alias, but it did")
  }
}

func Test_CmdUnaliasesViaUnaliasCommand(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  p.Store.InitPlayerData(p.GetID())

  go p.Cmd("unalias l")
  res := <-ch

  if !strings.Contains(res, "Alias 'l' removed.") {
    t.Errorf("Expected removal confirmation via Cmd(), but got '%s'", res)
  }

  if _, ok := p.Aliases["l"]; ok {
    t.Error("Expected 'l' to be removed via Cmd(), but it wasn't")
  }
}

func Test_CmdUnaliasWrongArgCount(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)

  go p.Cmd("unalias")
  res := <-ch

  if !strings.Contains(res, "Use 'unalias <name>' to remove an alias.") {
    t.Errorf("Expected usage message, but got '%s'", res)
  }
}

// Alias expansion must not create a loophole around skill restrictions like
// OOCOnly -- aliasing a restricted skill name shouldn't let it slip through
// mid-combat when typing the skill name directly wouldn't.

func Test_CmdAliasDoesNotBypassOOCRestrictionInCombat(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  p.Level = character.MaxLevel
  p.Aliases["ch"] = "charge"
  defer close(ch)
  go p.EnterCombat(&mocks.MockNP{})
  <-ch // "You attack %s!"

  go p.Cmd("ch")
  res := <-ch

  if !strings.Contains(res, "You cannot use 'Charge' in combat!") {
    t.Errorf("Expected aliased OOC-only skill to still be rejected in combat, but got '%s'", res)
  }

  if p.CmbSkill != nil {
    t.Errorf("Expected no skill to be set when using an alias for a restricted skill in combat, but got %s", p.CmbSkill.Name)
  }
}

// Each player must get their own independent copy of the default aliases --
// otherwise NewPlayer() could accidentally share (and let players corrupt)
// one global map instance across every connected player.

func Test_NewPlayerAliasesAreIndependent(t *testing.T) {
  p1, ch1, _ := NewTestPlayer()
  defer close(ch1)
  p2, ch2, _ := NewTestPlayer()
  defer close(ch2)

  p1.Aliases["l"] = "look very closely"

  if p2.Aliases["l"] != "look" {
    t.Errorf("Expected p2's aliases to be unaffected by mutating p1's, but got '%s'", p2.Aliases["l"])
  }

  if DefaultAliases["l"] != "look" {
    t.Errorf("Expected the shared DefaultAliases map to be unaffected by player mutation, but got '%s'", DefaultAliases["l"])
  }
}

// Alias names should be case-insensitive on set, expand, and remove -- 'GS',
// 'gs', and 'Gs' should all refer to the same alias.

func Test_SetAliasIsCaseInsensitive(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  p.Store.InitPlayerData(p.GetID())

  go p.SetAlias("GS", "attack slime shove")
  <-ch

  if p.Aliases["gs"] != "attack slime shove" {
    t.Errorf("Expected alias name to be normalized to lowercase, but got %v", p.Aliases)
  }

  if _, ok := p.Aliases["GS"]; ok {
    t.Error("Expected alias key to be stored lowercase, not with its original casing")
  }
}

func Test_ExpandAliasIsCaseInsensitive(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)

  res := p.expandAlias("L")
  if res != "look" {
    t.Errorf("Expected uppercase 'L' to match the 'l' alias, but got '%s'", res)
  }

  p.Aliases["gs"] = "attack slime"
  res = p.expandAlias("Gs shove")
  if res != "attack slime shove" {
    t.Errorf("Expected mixed-case alias invocation to match, but got '%s'", res)
  }
}

func Test_RemoveAliasIsCaseInsensitive(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  p.Store.InitPlayerData(p.GetID())

  go p.RemoveAlias("L")
  res := <-ch

  if !strings.Contains(res, "Alias 'L' removed.") {
    t.Errorf("Expected removal confirmation, but got '%s'", res)
  }

  if _, ok := p.Aliases["l"]; ok {
    t.Error("Expected 'l' to be removed regardless of the case used to remove it")
  }
}
