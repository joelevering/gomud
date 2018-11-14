package player

import (
  "fmt"
  "log"
  "net"
  "strings"
  "time"
  "unicode/utf8"

  "github.com/joelevering/gomud/character"
  "github.com/joelevering/gomud/classes"
  "github.com/joelevering/gomud/combat"
  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/room"
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
  "github.com/joelevering/gomud/storage"
  "github.com/joelevering/gomud/structs"
)

type Player struct {
  *character.Character

  Channel chan string
  Queue   interfaces.QueueI
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
    for _, class := range classes.StartingClasses {
      p.persistClass(class.GetName())
    }
  } else {
    p.loadClass(classes.Conscript)
    p.loadChar()
  }

  if p.GetSpawn() == nil {
    p.SetSpawn(room.RoomStore.Default)
  }

  if p.Room == nil {
    p.Spawn()
  } else {
    p.EnterRoom(p.Room)
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
  p.log("Save")
  p.persistClass(p.GetClassName())
  p.Store.PersistChar(p.GetName(), p.Character)
}

func (p *Player) Cmd(cmd string) {
  p.log(fmt.Sprintf("Cmd %s", cmd))
  if p.IsInCombat() {
    p.useSkill(cmd)
    return
  }

  words := strings.Split(cmd, " ")

  switch strings.ToLower(words[0]) {
  case "ls", "list":
    p.List()
  case "l", "look":
    if len(words) == 1 {
      p.Look()
    } else if len(words) > 1 {
      p.LookTarget(strings.Join(words[1:], " "))
    }
  case "m", "move":
    if len(words) == 2 {
      p.Move(words[1])
    } else {
      p.SendMsg("I'm not sure where you're trying to go. Try again with a correct exit key.")
    }
  case "h", "help":
    helpMsg := Help(words)
    for _, ln := range strings.Split(helpMsg, "\n") {
      p.SendMsg(ln)
    }
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
    } else if len(words) > 2 {
      skName := strings.Join(words[2:], " ")
      p.AttackNP(words[1], skName)
    } else {
      p.SendMsg("I'm not sure how to interpret your attack. Use either 'attack <first name of enemy> <skill name>' or omit the skill name.")
    }
  case "st", "status":
    p.Status()
  case "cl", "classes":
    p.ListClasses()
  case "c", "change":
    if len(words) > 1 {
      if words[1] == "subclass" || words[1] == "sc" {
        p.ChangeSubclass(words[2])
      } else {
        p.ChangeClass(words[1])
      }
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
      if otherP.IsInCombat() {
        names = append(names, fmt.Sprintf("%s (in combat)", otherP.GetName()))
      } else {
        names = append(names, otherP.GetName())
      }
    }
  }

  for _, np := range p.Room.GetNPs() {
    if np.IsAlive() {
      if np.IsInCombat() {
        names = append(names, fmt.Sprintf("%s (NPC, in combat)", np.GetName()))
      } else {
        names = append(names, fmt.Sprintf("%s (NPC)", np.GetName()))
      }
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

func (p Player) LookTarget(name string) {
  for _, np := range p.Room.GetNPs() {
    if np.IsAlive() && strings.Contains(strings.ToUpper(np.GetName()), strings.ToUpper(name)) {
      p.SendMsg(
        fmt.Sprintf("You look at %s and see:", np.GetName()),
        np.GetDesc(),
      )
      if np.IsInCombat() {
        p.SendMsg("They're in the middle of a fight!")
      }

      return
    }
  }

  for _, pl := range p.Room.GetPlayers() {
    if strings.Contains(strings.ToUpper(pl.GetName()), strings.ToUpper(name)) {
      p.SendMsg(fmt.Sprintf("You look at %s and see:", pl.GetName()))

      if pl.IsInCombat() {
        p.SendMsg(fmt.Sprintf("A level %d %s in the middle of a fight!.", pl.GetLevel(), pl.GetClassName()))
      } else {
        p.SendMsg(fmt.Sprintf("A level %d %s.", pl.GetLevel(), pl.GetClassName()))
      }

      return
    }
  }

  p.SendMsg("Are you sure they're here??")
}

func (p *Player) Status() {
  header := fmt.Sprintf("~~~~~~~~~~*%s*~~~~~~~~~~", p.GetName())
  p.SendMsg(header)
  p.SendMsg(fmt.Sprintf("Class: %s", p.GetClassName()))
  if p.IsMaxLevel() {
    p.SendMsg(fmt.Sprintf("Level: %d (MAX)", character.MaxLevel))
    p.SendMsg(fmt.Sprintf("Experience: %d/MAX", p.GetExp()))
  } else {
    p.SendMsg(fmt.Sprintf("Level: %d", p.GetLevel()))
    p.SendMsg(fmt.Sprintf("Experience: %d/%d", p.GetExp(), p.GetNextLvlExp()))
  }
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
  p.SendMsg("")
  p.SendMsg("Skills:")
  for _, sk := range p.GetSkills() {
    p.SendMsg(fmt.Sprintf(" * %s", sk.Name))
  }
  p.SendMsg(strings.Repeat("~", utf8.RuneCountInString(header)))
}

func (p *Player) ListClasses() {
  p.SendMsg("Your Classes:")
  for name, stats := range p.loadClasses() {
    p.SendMsg("")

    subheader := classes.Find(name).GetDesc()

    shCount := utf8.RuneCountInString(subheader)
    nameCount := utf8.RuneCountInString(name)
    tildes := strings.Repeat("~", (shCount - nameCount)/2)

    header := (tildes + "*" + name + "*" + tildes)

    p.SendMsg(header)
    p.SendMsg(subheader)
    p.SendMsg("")

    var lvl, maxDet, exp, nextLvlExp int

    if name == p.GetClassName() {
      lvl = p.GetLevel()
      maxDet = p.GetMaxDet()
      exp = p.GetExp()
      nextLvlExp = p.GetNextLvlExp()
    } else {
      lvl = stats.Lvl
      maxDet = stats.MaxDet
      exp = stats.Exp
      nextLvlExp = stats.NextLvlExp
    }

    if lvl == character.MaxLevel {
      p.SendMsg(fmt.Sprintf("Level: %d (MAX)", lvl))
      p.SendMsg(fmt.Sprintf("Max. Determination: %d", maxDet))
      p.SendMsg(fmt.Sprintf("Experience: %d/MAX", exp))
    } else {
      p.SendMsg(fmt.Sprintf("Level: %d", lvl))
      p.SendMsg(fmt.Sprintf("Max. Determination: %d", maxDet))
      p.SendMsg(fmt.Sprintf("Experience: %d/%d", exp, nextLvlExp))
    }

    p.SendMsg(strings.Repeat("~", utf8.RuneCountInString(header)))
  }
}

func (p *Player) AttackNP(npName, skName string) {
  for _, np := range p.Room.GetNPs() {
    if np.IsAlive() && strings.Contains(strings.ToUpper(np.GetName()), strings.ToUpper(npName)) {
      p.useSkill(skName)
      go combat.Start(p, np, p.Room)
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
  p.Save()
  cl := classes.Find(class)
  if cl == nil {
    p.SendMsg(fmt.Sprintf("Couldn't find class '%s'. Please use `classes` to confirm the spelling.", strings.Title(class)))
    return
  }

  if p.classIsLocked(cl.GetName()) {
    p.SendMsg(fmt.Sprintf("You haven't unlocked '%s' yet!", cl.GetName()))
    return
  }

  p.loadClass(cl)
  p.SendMsg(fmt.Sprintf("Changed to %s!", p.Class.GetName()))
}

func (p *Player) ChangeSubclass(clName string) {
  cl := classes.Find(clName)
  if cl == nil {
    p.SendMsg(fmt.Sprintf("Couldn't find subclass '%s'. Please use `classes` to confirm the spelling.", strings.Title(clName)))
    return
  }

  if cl.GetTier() >= p.Class.GetTier() {
    p.SendMsg(fmt.Sprintf("%s is in or above your current Tier. Change your main class with `change <class name>`.", cl.GetName()))
    p.SendMsg(fmt.Sprintf("You're still a %s", p.GetHybridClassName()))
    return
  }

  pCl := p.loadClasses()[cl.GetName()]
  if pCl.Lvl == 0 {
    p.SendMsg(fmt.Sprintf("You need to unlock %s to use it as a subclass.", cl.GetName()))
    p.SendMsg(fmt.Sprintf("You're still a %s", p.GetHybridClassName()))
    return
  } else if pCl.Lvl != character.MaxLevel {
    p.SendMsg(fmt.Sprintf("You need to reach maximum level with %s to use it as a subclass.", cl.GetName()))
    p.SendMsg(fmt.Sprintf("You're still a %s", p.GetHybridClassName()))
    return
  }

  p.Classes[cl.GetTier()] = cl
  p.SendMsg(fmt.Sprintf("Set %s as a subclass.", cl.GetName()))
  p.SendMsg(fmt.Sprintf("You're now a %s", p.GetHybridClassName()))
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

  if rep.Concentrating {
    p.SendMsg("You are concentrating on your enemy.")
  }

  if rep.Skill.Name != "" {
    if rep.Concentrating {
      p.SendMsg(fmt.Sprintf("You were unable to use %s!", rep.Skill.Name))
    } else if rep.FollowUpReq != "" {
      p.SendMsg(fmt.Sprintf("%s failed! It has to follow up %s", rep.Skill.Name, rep.FollowUpReq))
    } else {
      p.SendMsg(fmt.Sprintf("You used %s!", rep.Skill.Name))
    }
  }

  if rep.Missed {
    p.SendMsg("Your attack missed!")
  }

  if rep.Dodged {
    p.SendMsg(fmt.Sprintf("%s dodged your attack!", opp.GetName()))
  }

  if rep.Redirected {
    p.SendMsg(fmt.Sprintf("%s redirected your attack!", opp.GetName()))
  }

  if rep.Heal > 0 {
    p.SendMsg(fmt.Sprintf("You healed %d damage!", rep.Heal))
  }

  if rep.StmRec > 0 {
    p.SendMsg(fmt.Sprintf("You recovered %d stamina!", rep.StmRec))
  }

  if rep.Empowered {
    p.SendMsg("You dealt increased damage!")
  }

  if rep.Weak {
    p.SendMsg("You dealt lowered damage!")
  }

  if rep.Steeled {
    p.SendMsg(fmt.Sprintf("%s steeled themselves, taking lowered damage!", opp.GetName()))
  }

  if rep.Vulnerable {
    p.SendMsg(fmt.Sprintf("%s is vulnerable. They took increased damage!", opp.GetName()))
  }

  if rep.Dmg > 0 {
    if opp.GetDet() == 0 {
      p.SendMsg(fmt.Sprintf("%s took %d damage!", opp.GetName(), rep.Dmg))
    } else {
      p.SendMsg(fmt.Sprintf("%s took %d damage! %s has %d/%d health left!", opp.GetName(), rep.Dmg, opp.GetName(), opp.GetDet(), opp.GetMaxDet()))
    }
  }

  if rep.SelfDmg > 0 {
    p.SendMsg(fmt.Sprintf("You took %d damage from the redirection! You have %d/%d health left!", rep.SelfDmg, p.GetDet(), p.GetMaxDet()))
  }

  for _, e := range rep.SFx {
    switch e.Effect {
    case statfx.Stun:
      p.SendMsg(fmt.Sprintf("%s was stunned!", opp.GetName()))
    }
  }

  for _, e := range rep.Dots {
    switch e.Type {
    case statfx.Bleed:
      p.SendMsg(fmt.Sprintf("%s started bleeding!", opp.GetName()))
    case statfx.Fire:
      p.SendMsg(fmt.Sprintf("%s caught on fire!", opp.GetName()))
    }
  }

  if rep.Surprised != (structs.SurpriseRep{}) {
    if rep.Surprised.Stunned {
      p.SendMsg(fmt.Sprintf("You surprised %s! They're stunned!", opp.GetName()))
    }
    if rep.Surprised.Weak {
      p.SendMsg(fmt.Sprintf("You surprised %s! They're off-balance!", opp.GetName()))
    }
    if rep.Surprised.Vulnerable {
      p.SendMsg(fmt.Sprintf("You surprised %s! They're vulnerable!", opp.GetName()))
    }
  }
}

func (p *Player) ReportDef(opp interfaces.Combatant, rep structs.CmbRep) {
  if rep.Skill.Name != "" {
    p.SendMsg(fmt.Sprintf("%s used %s!", opp.GetName(), rep.Skill.Name))
  }

  for _, d := range rep.DotDmgs {
    p.SendMsg(fmt.Sprintf("%s took %d %s damage!", opp.GetName(), d.Dmg, d.Type))
  }

  if rep.Missed {
    p.SendMsg("%s missed their attack!", opp.GetName())
  }

  if rep.Dodged {
    p.SendMsg("You dodged the attack!")
  }

  if rep.Redirected {
    p.SendMsg("You redirected the attack!")
  }

  if rep.Heal > 0 {
    p.SendMsg(fmt.Sprintf("%s healed %d damage!", opp.GetName(), rep.Heal))
  }

  if rep.StmRec > 0 {
    p.SendMsg(fmt.Sprintf("%s recovered stamina!", opp.GetName()))
  }

  if rep.Empowered {
    p.SendMsg(fmt.Sprintf("%s dealt increased damage!", opp.GetName()))
  }

  if rep.Weak {
    p.SendMsg(fmt.Sprintf("%s dealt lowered damage!", opp.GetName()))
  }

  if rep.Steeled {
    p.SendMsg("You steeled yourself, taking lowered damage!")
  }

  if rep.Vulnerable {
    p.SendMsg("You're vulnerable. You took increased damage!")
  }

  if rep.Dmg > 0 {
    p.SendMsg(fmt.Sprintf("You were attacked for %d damage! You have %d/%d health left!", rep.Dmg, p.GetDet(), p.GetMaxDet()))
  }

  if rep.SelfDmg > 0 {
    p.SendMsg(fmt.Sprintf("%s took %d damage from the redirection!", opp.GetName(), rep.SelfDmg))
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
  p.log(fmt.Sprintf("LoseCombat vs %s", winner.GetName()))
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
  p.log(fmt.Sprintf("WinCombat vs %s", loser.GetName()))
  p.LeaveCombat()

  p.SendMsg(fmt.Sprintf("%s is defeated!", loser.GetName()))

  expGained := loser.GetExpGiven()
  p.GainExp(expGained)
}

func (p *Player) GainExp(exp int) {
  oldLvl := p.GetLevel()

  p.Character.GainExp(exp)

  if p.GetLevel() > oldLvl {
    p.SendMsg(fmt.Sprintf("You gained %d experience and leveled up!", exp))
    p.SendMsg(fmt.Sprintf("You're now level %d!", p.GetLevel()))
    newSk := p.Class.SkillForLvl(p.Level)
    if newSk != nil {
      p.SendMsg(fmt.Sprintf("You gained the skill '%s'!", newSk.Name))
    }
    if p.IsMaxLevel() {
      p.SendMsg(fmt.Sprintf("You've reached the maximum level with the %s class!", p.GetClassName()))

      p.unlockNewClasses()
    }
  } else {
    p.SendMsg(fmt.Sprintf("You gained %d experience! You need %d more experience to level up.", exp, p.ExpToLvl()))
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
  p.Classes[class.GetTier()] = class
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

  curRm := room.RoomStore.Find(l.Room)
  if curRm != nil {
    p.Room = curRm
  }

  spawn := room.RoomStore.Find(l.Spawn)
  if spawn != nil {
    p.SetSpawn(spawn)
  }

  if l.Class != "" {
    cl := classes.Find(l.Class)
    p.loadClass(cl)
  }
}

func (p *Player) unlockNewClasses() {
  classStats := p.loadClasses()

  for _, cl := range classes.PlayerClasses {
    if cl.GetTier() != p.Class.GetTier() + 1 {
      continue
    }

    var reqsMet int
    for _, req := range cl.GetReqs() {
      if req.GetName() == p.GetClassName() {
        reqsMet += 1
        continue
      }

      loadedClass := classStats[req.GetName()]
      if loadedClass.Lvl == character.MaxLevel {
        reqsMet += 1
      }
    }

    if reqsMet == len(cl.GetReqs()) {
      p.persistClass(cl.GetName())
      p.SendMsg(fmt.Sprintf("You unlocked a new class: '%s'!", cl.GetName()))
    }
  }
}

func (p *Player) useSkill (skName string) {
  if skName == "" { return }

  var sk *skills.Skill

  for t, cl := range p.Classes {
    if cl == nil { continue }

    if t < p.Class.GetTier() {
      sk = cl.GetSkill(skName, character.MaxLevel)
    } else { // it's the main class
      sk = cl.GetSkill(skName, p.Level)
    }

    if sk != nil {
      if p.IsInCombat() && sk.IsOOCOnly() {
        p.SendMsg(fmt.Sprintf("You cannot use '%s' in combat!", sk.Name))
        return
      }

      p.SetCmbSkill(sk)
      p.SendMsg(fmt.Sprintf("Preparing %s", sk.Name))

      return
    }
  }

  p.SendMsg(fmt.Sprintf("You don't know how to prepare '%s'!", skName))
}

func (p *Player) loadClasses() map[string]storage.ClassStats {
  return p.Store.LoadClasses(p.GetID())
}

func (p *Player) classIsLocked(clName string) bool {
  return p.loadClasses()[clName].Lvl == 0
}

func (p *Player) log(msg string) {
  log.Print(fmt.Sprintf("%s - %d - %s", p.GetName(), p.GetRoom().GetID(), msg))
}
