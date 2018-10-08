package interfaces

import "net"

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
  Die(CliI)
  IsAlive() bool
  SetBehavior(QueueI)
}

type CliI interface {
	StartWriter(net.Conn)
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
  Die(NPCI)
  Defeat(NPCI)
	GetName() string
	GetRoom() RoomI
  GetHealth() int
  SetHealth(int)
  GetMaxHealth() int
  GetStr() int
  GetEnd() int
  GetCombatCmd() []string
}

type QueueI interface {
  Sub(string) chan bool
  Unsub(string, chan bool)
  Pub(string)
}
