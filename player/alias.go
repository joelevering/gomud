package player

import (
  "fmt"
  "strings"

  "github.com/joelevering/gomud/skills"
)

// DefaultAliases mirrors the shorthand letters gomud has always supported
// (e.g. 'l' for 'look'), now expressed as ordinary, player-removable
// aliases instead of hardcoded switch cases.
var DefaultAliases = map[string]string{
  "ls": "list",
  "l":  "look",
  "m":  "move",
  "h":  "help",
  "s":  "say",
  "y":  "yell",
  "em": "emote",
  "a":  "attack",
  "st": "status",
  "cl": "classes",
  "c":  "change",
  "al": "alias",
  "ua": "unalias",
}

// reservedCmds are the long-form command words the switch in Cmd() actually
// dispatches on. These can never be used as an alias name, so a player can't
// shadow a real command. Default alias letters are NOT in this set on
// purpose -- they're just data and can be freely reassigned or removed.
var reservedCmds = []string{
  "list", "look", "move", "help", "say", "yell", "emote", "attack",
  "status", "classes", "change", "alias", "unalias", "exit", "quit",
}

func isReservedCmd(cmd string) bool {
  for _, c := range reservedCmds {
    if c == cmd {
      return true
    }
  }
  return false
}

func copyDefaultAliases() map[string]string {
  aliases := make(map[string]string, len(DefaultAliases))
  for k, v := range DefaultAliases {
    aliases[k] = v
  }
  return aliases
}

func (p *Player) SetAlias(name, expansion string) {
  key := strings.ToLower(name)

  if isReservedCmd(key) {
    p.SendMsg(fmt.Sprintf("'%s' is a reserved command and can't be used as an alias name.", name))
    return
  }

  fields := strings.Fields(expansion)
  if len(fields) == 0 {
    p.SendMsg("An alias needs to run something -- use 'alias <name> <command>'.")
    return
  }

  // A regular command is "verb + args", so only the first word matters
  // (e.g. the NPC name in "attack slime" is unchecked, as expected). A
  // skill name has no such structure -- it's a single multi-word proper
  // noun (e.g. "Desperate Blow"), so it's checked as a whole.
  cmdWord := strings.ToLower(fields[0])
  if !isReservedCmd(cmdWord) && skills.GetSkill(expansion) == nil {
    p.SendMsg(fmt.Sprintf("'%s' isn't a recognized command or skill, so that alias wouldn't do anything.", expansion))
    return
  }

  p.Aliases[key] = expansion
  p.Store.PersistAliases(p.GetID(), p.Aliases)
  p.SendMsg(fmt.Sprintf("Alias set: '%s' now runs '%s'", name, expansion))
}

func (p *Player) RemoveAlias(name string) {
  key := strings.ToLower(name)

  if _, ok := p.Aliases[key]; !ok {
    p.SendMsg(fmt.Sprintf("You don't have an alias named '%s'.", name))
    return
  }

  delete(p.Aliases, key)
  p.Store.PersistAliases(p.GetID(), p.Aliases)
  p.SendMsg(fmt.Sprintf("Alias '%s' removed.", name))
}

func (p *Player) ListAliases() {
  if len(p.Aliases) == 0 {
    p.SendMsg("You don't have any aliases set. Use 'alias <name> <command>' to create one.")
    return
  }

  p.SendMsg("Your aliases:")
  for name, expansion := range p.Aliases {
    p.SendMsg(fmt.Sprintf(" * %s -> %s", name, expansion))
  }
}

func (p *Player) expandAlias(cmd string) string {
  words := strings.Split(cmd, " ")

  expansion, ok := p.Aliases[strings.ToLower(words[0])]
  if !ok {
    return cmd
  }

  if len(words) > 1 {
    return expansion + " " + strings.Join(words[1:], " ")
  }

  return expansion
}

func (p *Player) loadAliases() {
  aliases := p.Store.LoadAliases(p.GetID())
  if aliases != nil {
    p.Aliases = aliases
  }
}
