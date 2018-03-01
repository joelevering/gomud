package mocks

import "github.com/joelevering/gomud/interfaces"

type MockRoom struct {
	Messages []string
	Clients  []interfaces.CliI
}

func (m *MockRoom) Message(msg string) {
	m.Messages = append(m.Messages, msg)
}

func (m *MockRoom) AddCli(cli interfaces.CliI) {
}

func (m *MockRoom) RemoveCli(cli interfaces.CliI, msg string) {
}

func (m *MockRoom) GetExits() []interfaces.ExitI {
	return []interfaces.ExitI{}
}

func (m *MockRoom) GetNpcs() []interfaces.NPCI {
	return []interfaces.NPCI{}
}

func (m *MockRoom) GetClients() []interfaces.CliI {
	return []interfaces.CliI{}
}

func (m *MockRoom) GetName() string {
	return "Name"
}

func (m *MockRoom) GetDesc() string {
	return "Desc"
}

func (m *MockRoom) GetID() int {
	return 0
}
