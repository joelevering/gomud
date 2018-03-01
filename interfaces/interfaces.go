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
	SetHealth(int)
}

type CliI interface {
	StartWriter(net.Conn)
	List()
	Look()
	LookNPC(string)
	AttackNPC(string)
	Move(string)
	Say(string)
	Yell(string)
	SendMsg(...string)
	LeaveRoom(string)
	EnterRoom(RoomI)
	GetName() string
	GetRoom() RoomI
}
