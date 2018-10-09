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
	AddCli(CliI)
	RemoveCli(CliI, string)
	GetExits() []ExitI
	GetNpcs() []NPCI
	GetClients() []CliI
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
	GetName() string
	GetDesc() string
	GetMaxDet() int
	GetDet() int
	SetDet(int)
	GetStr() int
  GetAtk() int
  GetDef() int
  GetExp() int
  SetSpawn(RoomI)
  Spawn()
  Say(string)
  Emote(string)
  LoseCombat(CharI)
  IsAlive() bool
  SetBehavior(QueueI)
}

type CliI interface {
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
  LoseCombat(NPCI)
  WinCombat(NPCI)
}

type CharI interface {
  GetClassName() string
	GetName() string
  SetName(string)
  GetDet() int
  SetDet(int)
  GetMaxDet() int
  SetMaxDet(int)
  GetMaxStm() int
  SetMaxStm(int)
  GetMaxFoc() int
  SetMaxFoc(int)
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
