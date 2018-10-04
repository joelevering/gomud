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
  Die()
  IsAlive() bool
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
	GetName() string
	GetRoom() RoomI
}

type CharI interface {
  GetName() string
  SendMsg(string)
  IsHidden() bool // sub for IsAlive on NPC -- could be used for PC admins
}
