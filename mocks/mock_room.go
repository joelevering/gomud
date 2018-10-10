package mocks

import "github.com/joelevering/gomud/interfaces"

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
    &MockNPC{
      // Id:     1,
      // Desc:   "Holding a purple crayon",
      // Name:   "Harold",
      // MaxDet: 100,
      // Det:    99,
      // Str:    98,
      // Flo:    97,
      // Exp:    2,
      // Alive:  true,
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
