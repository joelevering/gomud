package interfaces

import (
  "net"

  "github.com/joelevering/gomud/classes"
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
	GetHealth() int
	GetMaxHealth() int
	GetEnd() int
	GetStr() int
  GetExp() int
	SetHealth(int)
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
	GetName() string
  SetName(string)
  GetHealth() int
  SetHealth(int)
  GetMaxHealth() int
  SetMaxHealth(int)
  GetStr() int
  SetStr(int)
  GetEnd() int
  SetEnd(int)
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
}
