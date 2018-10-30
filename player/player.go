package player

import (
  "fmt"
  "net"
  "strings"
  "time"
  "unicode/utf8"

  "github.com/joelevering/gomud/character"
  "github.com/joelevering/gomud/classes"
  "github.com/joelevering/gomud/combat"
  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/statfx"
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/storage"
  "github.com/joelevering/gomud/structs"
)

const helpMsg = `
Available commands:
'say <message>' to communicate with people in your room
'move <exit key>' to move to a new room
'look' to see where you are
'look <npc name>' to see more details about an NPC
'list' to see who is currently in your room
'help' to repeat this message
'status' to see details about your character
'attack <npc name>' to start combat
'change <class name>' to change your class
'exit' or 'quit' to log out

Most commands have their first letter as a shortcut
`

type Player struct {
  *character.Character

  Channel chan string
  Queue   interfaces.QueueI
  Room    interfaces.RoomI
  Store   storage.StorageI
  Logout  chan string
}

func NewPlayer(ch chan string, q interfaces.QueueI, s storage.StorageI) *Player {
  return &Player{
    Character: character.NewCharacter(),
    Channel:   ch,
    Queue:     q,
    Store:     s,
    Logout:    make(chan string),
  }
}

func (p Player) StartWriter(conn net.Conn) {
  for msg := range p.Channel {
   fmt.Fprintln(conn, msg)
  }
}

func (p *Player) Init() {
  if !p.Store.StoreExists(p.GetID()) {
    p.Store.InitPlayerData(p.GetID())
    for _, class := range classes.PlayerClasses {
      p.persistClass(class.GetName())
    }
  } else {
    p.loadClass(classes.Conscript)
    p.loadChar()
  }

  go p.regen()
}

func (p *Player) regen() {
  tickTime := 5 * time.Second
  time.Sleep(tickTime)

  for {
    select {
    default:
      if !p.IsInCombat() {
        p.HealPct(0.025)
        p.RefocusPct(0.025)
        p.RecuperatePct(0.025)
      }

      time.Sleep(tickTime)
    case <-p.Logout:
      return
    }
  }
}

func (p *Player) Save() {
  p.persistClass(p.GetClassName())
  p.Store.PersistChar(p.GetName(), p.Character)
}

func (p *Player) Cmd(cmd string) {
  words := strings.Split(cmd, " ")

  if p.IsInCombat() {
    p.useSkill(words[0], true)
    return
  }

  switch strings.ToLower(words[0]) {
  case "ls", "list":
    p.List()
  case "l", "look":
    if len(words) == 1 {
      p.Look()
    } else if len(words) > 1 {
      p.LookNP(words[1])
    }
  case "m", "move":
    if len(words) == 2 {
      p.Move(words[1])
    } else {
      p.SendMsg("I'm not sure where you're trying to go. Try again with a correct exit key.")
    }
  case "h", "help":
    p.SendMsg(helpMsg)
  case "s", "say":
    if len(words) > 1 {
      p.Say(strings.Join(words[1:], " "))
    } else {
      p.SendMsg("If you want to say something, include a message. E.g. 'say hello there!'")
    }
  case "y", "yell":
    if len(words) > 1 {
      p.Yell(strings.Join(words[1:], " "))
    } else {
      p.SendMsg("If you want to yell something, include a message. E.g. 'yell HELLO THERE!'")
    }
  case "a", "attack":
    if len(words) == 2 {
      p.AttackNP(words[1], "")
    } else if len(words) == 3 {
      p.AttackNP(words[1], words[2])
    } else {
      p.SendMsg("I'm not sure how to interpret your attack. Use either 'attack <first name of enemy> <skill name>' or omit the skill name.")
    }
  case "st", "status":
    p.Status()
  case "c", "change":
    if len(words) == 2 {
      p.ChangeClass(words[1])
    } else {
      p.SendMsg("I'm not sure how to interpret your class change. Use 'change <class name>' and try again.")
    }
  default:
    p.SendMsg("I'm not sure what you mean. Type 'help' for assistance.")
  }
}

func (p Player) List() {
  names := []string{fmt.Sprintf("Yourself (%s)", p.GetName())}

  for _, otherP := range p.Room.GetPlayers() {
    if otherP.GetName() != p.GetName() {
      names = append(names, otherP.GetName())
    }
  }

  for _, np := range p.Room.GetNPs() {
    if np.IsAlive() {
    names = append(names, fmt.Sprintf("%s (NPC)", np.GetName()))
    }
  }

  p.SendMsg(fmt.Sprintf("You look around and see: %s", strings.Join(names, ", ")))
}

func (p Player) Look() {
  p.SendMsg(fmt.Sprintf("~~%s~~", p.Room.GetName()))
  p.SendMsg(strings.Split(p.Room.GetDesc(), "\n")...)
  p.SendMsg("", "Exits:")

  for _, exit := range p.Room.GetExits() {
    p.SendMsg(fmt.Sprintf("- %s", exit.GetDesc()))
  }

  p.SendMsg("")
  p.List()
}

func (p Player) LookNP(npName string) {
  for _, np := range p.Room.GetNPs() {
    if np.IsAlive() && strings.Contains(strings.ToUpper(np.GetName()), strings.ToUpper(npName)) {
      p.SendMsg(
        fmt.Sprintf("You look at %s and see:", np.GetName()),
        np.GetDesc(),
      )

      return
    }
  }

  p.SendMsg("Are you sure they're here??")
}

func (p *Player) Status() {
  header := fmt.Sprintf("~~~~~~~~~~*%s*~~~~~~~~~~", p.GetName())
  p.SendMsg(header)
  p.SendMsg(fmt.Sprintf("Class: %s", p.GetClassName()))
  p.SendMsg(fmt.Sprintf("Level: %d", p.GetLevel()))
  p.SendMsg(fmt.Sprintf("Experience: %d/%d", p.GetExp(), p.GetNextLvlExp()))
  p.SendMsg("")
  p.SendMsg(fmt.Sprintf("Determination: %d/%d", p.GetDet(), p.GetMaxDet()))
  p.SendMsg(fmt.Sprintf("Stamina: %d/%d", p.GetStm(), p.GetMaxStm()))
  p.SendMsg(fmt.Sprintf("Focus: %d/%d", p.GetFoc(), p.GetMaxFoc()))
  p.SendMsg("")
  p.SendMsg(fmt.Sprintf("Strength: %d", p.GetStr()))
  p.SendMsg(fmt.Sprintf("Flow: %d", p.GetFlo()))
  p.SendMsg(fmt.Sprintf("Ingenuity: %d", p.GetIng()))
  p.SendMsg(fmt.Sprintf("Knowledge: %d", p.GetKno()))
  p.SendMsg(fmt.Sprintf("Sagacity: %d", p.GetSag()))
  p.SendMsg(strings.Repeat("~", utf8.RuneCountInString(header)))
}

func (p *Player) AttackNP(npName, skName string) {
  for _, np := range p.Room.GetNPs() {
    if np.IsAlive() && strings.Contains(strings.ToUpper(np.GetName()), strings.ToUpper(npName)) {
      p.useSkill(skName, false)
      go combat.Start(p, np)
      return
    }
  }

  p.SendMsg("Are you sure they're here?")
}

func (p *Player) EnterCombat(opp interfaces.Combatant) {
  p.InCombat = true
  p.SendMsg(fmt.Sprintf("You attack %s!", opp.GetName()))
}

func (p *Player) Move(exitKey string) {
  for _, exit := range p.Room.GetExits() {
    if strings.ToUpper(exitKey) == strings.ToUpper(exit.GetKey()) {
      p.LeaveRoom(fmt.Sprintf("%s heads to %s!", p.GetName(), exit.GetRoom().GetName()))
      p.EnterRoom(exit.GetRoom())
      p.Look()
      return
    }
  }

  p.SendMsg("Where are you trying to go??")
}

func (p *Player) Say(msg string) {
  if msg != "" {
    p.Room.Message(fmt.Sprintf("%s says \"%s\"", p.GetName(), msg))
  }
}

func (p *Player) Yell(msg string) {
  if msg != "" {
    fullMsg := fmt.Sprintf("%s yells \"%s\"", p.GetName(), msg)
    p.Room.Message(fullMsg)

    for _, exit := range p.Room.GetExits() {
      exit.GetRoom().Message(fullMsg)
    }
  }
}

func (p *Player) ChangeClass(class string) {
  switch strings.ToLower(class) {
  case "conscript":
    p.Save()
    p.loadClass(classes.Conscript)
  case "athlete":
    p.Save()
    p.loadClass(classes.Athlete)
  case "charmer":
    p.Save()
    p.loadClass(classes.Charmer)
  case "augur":
    p.Save()
    p.loadClass(classes.Augur)
  case "sophist":
    p.Save()
    p.loadClass(classes.Sophist)
  }
}


func (p *Player) SendMsg(msgs ...string) {
  stamp := time.Now().Format(time.Kitchen)

  for _, msg := range msgs {
    p.Channel <- fmt.Sprintf("%s %s", stamp, msg)
  }
}

func (p *Player) LeaveRoom(msg string) {
  if msg == "" {
    msg = fmt.Sprintf("%s has left the room!", p.GetName())
  }

  p.Room.RemovePlayer(p, msg)
  p.Queue.Pub(fmt.Sprintf("pc-leaves-%d", p.Room.GetID()))
}

func (p *Player) EnterRoom(room interfaces.RoomI) {
  room.AddPlayer(p)
  p.Room = room
  p.Queue.Pub(fmt.Sprintf("pc-enters-%d", room.GetID()))
}

func (p *Player) ReportAtk(opp interfaces.Combatant, rep structs.CmbRep) {
  if rep.Stunned {
    p.SendMsg("You were unable to attack!")
  }

  if rep.Skill.Name != "" {
    p.SendMsg(fmt.Sprintf("You used %s!", rep.Skill.Name))
  }

  if rep.Missed {
    p.SendMsg("Your attack missed!")
  }

  if rep.Surprised != (structs.SurpriseRep{}) {
    if rep.Surprised.Stunned {
      p.SendMsg(fmt.Sprintf("You surprised %s! They're stunned!", opp.GetName()))
    }
    if rep.Surprised.LowerAtk {
      p.SendMsg(fmt.Sprintf("You surprised %s! They're off-balance!", opp.GetName()))
    }
    if rep.Surprised.LowerDef {
      p.SendMsg(fmt.Sprintf("You surprised %s! They're vulnerable!", opp.GetName()))
    }
  }

  if rep.Heal > 0 {
    p.SendMsg(fmt.Sprintf("You healed %d damage!", rep.Heal))
  }

  if rep.LowerAtk {
    p.SendMsg("You dealt lowered damage!")
  }

  if rep.LowerDef {
    p.SendMsg("You dealt increased damage!")
  }

  if rep.Dmg > 0 {
    if opp.GetDet() == 0 {
      p.SendMsg(fmt.Sprintf("%s took %d damage!", opp.GetName(), rep.Dmg))
    } else {
      p.SendMsg(fmt.Sprintf("%s took %d damage! %s has %d/%d health left!", opp.GetName(), rep.Dmg, opp.GetName(), opp.GetDet(), opp.GetMaxDet()))
    }
  }

  if len(rep.SFx) > 0 {
    for _, e := range rep.SFx {
      switch e.Effect {
      case statfx.Stun:
        p.SendMsg(fmt.Sprintf("%s was stunned!", opp.GetName()))
      }
    }
  }
}

func (p *Player) ReportDef(opp interfaces.Combatant, rep structs.CmbRep) {
  if rep.Skill.Name != "" {
    p.SendMsg(fmt.Sprintf("%s used %s!", opp.GetName(), rep.Skill.Name))
  }

  if rep.Missed {
    p.SendMsg("%s missed their attack!", opp.GetName())
  }

  if rep.Heal > 0 {
    p.SendMsg(fmt.Sprintf("%s healed %d damage!", opp.GetName(), rep.Heal))
  }

  if rep.LowerAtk {
    p.SendMsg("You took lowered damage!")
  }

  if rep.LowerDef {
    p.SendMsg("You took increased damage!")
  }

  if rep.Dmg > 0 {
    p.SendMsg(fmt.Sprintf("You took %d damage! You have %d/%d health left!", rep.Dmg, p.GetDet(), p.GetMaxDet()))
  }

  if len(rep.SFx) > 0 {
    for _, e := range rep.SFx {
      switch e.Effect {
      case statfx.Stun:
        p.SendMsg("You were stunned into inaction!")
      case statfx.Surprise:
        p.SendMsg("You were surprised by the attack!")
      }
    }
  }
}

func (p *Player) LoseCombat(winner interfaces.Combatant) {
  p.LeaveCombat()

  spawn := p.GetSpawn()

  deathNotice := fmt.Sprintf("%s was defeated by %s. Their body dissipates.", p.GetName(), winner.GetName())
  p.LeaveRoom(deathNotice)

  p.SendMsg(fmt.Sprintf("You were defeated by %s.", winner.GetName()))
  time.Sleep(1500 * time.Millisecond)
  p.EnterRoom(spawn)
  p.FullHeal()
  p.SendMsg(fmt.Sprintf("You find yourself back in a familiar place: %s", spawn.GetName()))
  time.Sleep(1500 * time.Millisecond)
  p.SendMsg("")
  p.Look()
}

func (p *Player) WinCombat(loser interfaces.Combatant) {
  p.LeaveCombat()

  p.SendMsg(fmt.Sprintf("%s is defeated!", loser.GetName()))

  expGained := loser.GetExpGiven()
  leveledUp := p.GainExp(expGained)

  if leveledUp {
    p.SendMsg(fmt.Sprintf("You gained %d experience and leveled up!", expGained))
    p.SendMsg(fmt.Sprintf("You're now level %d!", p.GetLevel()))
  } else {
    p.SendMsg(fmt.Sprintf("You gained %d experience! You need %d more experience to level up.", expGained, p.ExpToLvl()))
  }
}

func (p *Player) Spawn() {
  p.EnterRoom(p.GetSpawn())
}

// Getters and Setters

func (p *Player) GetID() string {
  return p.GetName()
}

func (p *Player) GetRoom() interfaces.RoomI {
  return p.Room
}

// Private

func (p *Player) persistClass(className string) {
  stats := storage.ClassStats{
    Lvl: p.GetLevel(),
    MaxDet: p.GetMaxDet(),
    Exp: p.GetExp(),
    NextLvlExp: p.GetNextLvlExp(),
  }

  p.Store.PersistClass(p.GetID(), className, stats)
}

func (p *Player) loadClass(class *classes.Class) {
  stats := p.Store.LoadStats(p.GetID(), class.GetName())

  p.Class = class
  p.Level = stats.Lvl
  p.MaxDet = stats.MaxDet
  if p.GetDet() > p.MaxDet {
    p.SetDet(p.MaxDet)
  }
  p.Exp = stats.Exp
  p.NextLvlExp = stats.NextLvlExp
}

func (p *Player) loadChar() {
  l := p.Store.LoadChar(p.GetID())

  p.Det = l.Det
  p.MaxStm = l.MaxStm
  p.Stm = l.Stm
  p.MaxFoc = l.Foc
  p.Str = l.Str
  p.Flo = l.Flo
  p.Ing = l.Ing
  p.Kno = l.Kno
  p.Sag = l.Sag
}

func (p *Player) useSkill (skName string, inCombat bool) {
  sk := p.Class.GetSkill(skName)
  if sk != nil {
    if inCombat && sk.Rstcn == skills.OOCOnly {
      p.SendMsg(fmt.Sprintf("You cannot use '%s' in combat!", sk.Name))
      return
    }

    p.SetCmbSkill(sk)
    p.SendMsg(fmt.Sprintf("Preparing %s", sk.Name))
  }
}
