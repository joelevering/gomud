package interfaces

import (
  "net"

  "github.com/joelevering/gomud/classes"
  "github.com/joelevering/gomud/stats"
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
	GetNpcs() []NPCI
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

type NPCI interface {
  Init(RoomI, QueueI)
  GetCharacter() CharI
	GetName() string
	GetDesc() string
  SetSpawn(RoomI)
  Spawn()
  Say(string)
  Emote(string)
  LoseCombat(CharI)
  IsAlive() bool
  SetBehavior(QueueI)
}

type PlI interface {
	StartWriter(net.Conn)
  GetName() string
  SetName(string)
	GetRoom() RoomI
  GetCombatCmd() []string
  SetCombatCmd([]string)
	List()
	Look()
	LookNPC(string)
  Status()
	AttackNPC(string)
	Move(string)
	Say(string)
	Yell(string)
	SendMsg(...string)
	LeaveRoom(string)
	EnterRoom(RoomI)
  LoseCombat(CharI)
  WinCombat(CharI)
}

type CharI interface {
  SetClass()
  ResetStats()
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
  GetLevel() int
  GetExp() int
  GetExpGiven() int
  GetNextLvlExp() int
  GetSpawn() RoomI
  SetSpawn(RoomI)

  Heal()
  EnterCombat()
  LeaveCombat()
  IsInCombat() bool
  GainExp(int) bool
  ExpToLvl() int
}

type ClassI interface {
  GetName() string
  GetStatGrowth() classes.StatGrowth
  GetAtkStats() []stats.Stat
  GetDefStats() []stats.Stat
}
