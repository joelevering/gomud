package interfaces

import (
  "net"

  "github.com/joelevering/gomud/classes"
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/stats"
  "github.com/joelevering/gomud/structs"
)

type QueueI interface {
  Sub(string) chan bool
  Unsub(string, chan bool)
  Pub(string)
}

type RoomI interface {
	Message(string)
	AddPlayer(PlI)
	RemovePlayer(PlI, string)
	GetExits() []ExitI
	GetNPs() []NPI
	GetPlayers() []PlI
	GetName() string
	GetDesc() string
	GetID() int
}

type ExitI interface {
	GetRoom() RoomI
	SetRoom(RoomI)
	GetRoomID() int
	GetKey() string
	GetDesc() string
}

type NPI interface {
  CharI
  Init(RoomI, QueueI)
  SetClass()
  ResetStats()
  GetDesc() string
  Spawn()
  Say(string)
  Emote(string)
  IsAlive() bool
  SetBehavior(QueueI)
  EnterCombat(Combatant)
  ReportAtk(Combatant, structs.CmbRep)
  ReportDef(Combatant, structs.CmbRep)
  WinCombat(Combatant)
  LoseCombat(Combatant)
}

type PlI interface {
  CharI
  StartWriter(net.Conn)
  Init()
  Save()
  GetID() string
  GetRoom() RoomI
  List()
  Look()
  LookNP(string)
  Status()
  AttackNP(string, string)
  Move(string)
  Say(string)
  Yell(string)
  SendMsg(...string)
  LeaveRoom(string)
  EnterRoom(RoomI)
  EnterCombat(Combatant)
  ReportAtk(Combatant, structs.CmbRep)
  ReportDef(Combatant, structs.CmbRep)
  WinCombat(Combatant)
  LoseCombat(Combatant)
}

type Combatant interface {
  EnterCombat(Combatant)
  AtkFx(*structs.CmbRep) structs.CmbFx
  ResistAtk(structs.CmbFx, *structs.CmbRep) structs.CmbFx
  ApplyAtk(structs.CmbFx, *structs.CmbRep)
  ApplyDef(structs.CmbFx, *structs.CmbRep)
  ReportAtk(Combatant, structs.CmbRep)
  ReportDef(Combatant, structs.CmbRep)
  TickFx()
  IsDefeated() bool
  WinCombat(Combatant)
  LoseCombat(Combatant)

  GetName() string
  GetExpGiven() int
  GetDet() int
  GetMaxDet() int
}

type CharI interface {
  GetClassName() string
  GetName() string
  SetName(string)
  GetMaxDet() int
  SetMaxDet(int)
  GetDet() int
  SetDet(int)
  GetMaxStm() int
  SetMaxStm(int)
  GetStm() int
  SetStm(int)
  GetMaxFoc() int
  SetMaxFoc(int)
  GetFoc() int
  SetFoc(int)
  GetStr() int
  SetStr(int)
  GetFlo() int
  SetFlo(int)
  GetIng() int
  SetIng(int)
  GetKno() int
  SetKno(int)
  GetSag() int
  SetSag(int)
  GetAtk() int
  GetDef() int
  SetCmbSkill(*skills.Skill)
  GetLevel() int
  GetExp() int
  GetExpGiven() int
  GetNextLvlExp() int
  GetSpawn() RoomI
  SetSpawn(RoomI)

  FullHeal()
  GainExp(int) bool
  ExpToLvl() int
  TickFx()

  IsInCombat() bool
  AtkFx(*structs.CmbRep) structs.CmbFx
  ResistAtk(structs.CmbFx, *structs.CmbRep) structs.CmbFx
  ApplyAtk(structs.CmbFx, *structs.CmbRep)
  ApplyDef(structs.CmbFx, *structs.CmbRep)
  IsDefeated() bool
}

type ClassI interface {
  GetName() string
  GetStatGrowth() classes.StatGrowth
  GetAtkStats() []stats.Stat
  GetDefStats() []stats.Stat
  GetSkill(string) *skills.Skill
}
