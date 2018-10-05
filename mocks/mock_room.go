package mocks

import (
	"github.com/joelevering/gomud/interfaces"
	"github.com/joelevering/gomud/npc"
)

type MockRoom struct {
	Messages   []string
	Clients    []interfaces.CliI
	Exits      []interfaces.ExitI
	AddedCli   interfaces.CliI
	RemovedCli interfaces.CliI
  NPCs       []interfaces.NPCI
	Name       string
}

func (m *MockRoom) Message(msg string) {
	m.Messages = append(m.Messages, msg)
}

func (m *MockRoom) AddCli(cli interfaces.CliI) {
	m.AddedCli = cli
}

func (m *MockRoom) RemoveCli(cli interfaces.CliI, msg string) {
	m.RemovedCli = cli
}

func (m *MockRoom) GetExits() []interfaces.ExitI {
	return m.Exits
}

func (m *MockRoom) GetNpcs() []interfaces.NPCI {
  if (len(m.NPCs) != 0) {
    return m.NPCs
  }

	return []interfaces.NPCI{
		&npc.NPC{
			Id:        1,
			Name:      "Harold",
			Desc:      "Holding a purple crayon",
			MaxHealth: 100,
			Health:    99,
			Str:       98,
			End:       97,
      Exp:       2,
      Alive:     true,
		},
	}
}

func (m *MockRoom) GetClients() []interfaces.CliI {
	return m.Clients
}

func (m *MockRoom) GetName() string {
	return m.Name
}

func (m *MockRoom) GetDesc() string {
	return "Desc"
}

func (m *MockRoom) GetID() int {
	return 0
}
