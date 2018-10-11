package mocks

import "github.com/joelevering/gomud/interfaces"

type MockRoom struct {
  Messages   []string
  Players    []interfaces.PlI
  Exits      []interfaces.ExitI
  AddedPlayer   interfaces.PlI
  RemovedPlayer interfaces.PlI
  NPs       []interfaces.NPI
  Name       string
}

func (m *MockRoom) Message(msg string) {
  m.Messages = append(m.Messages, msg)
}

func (m *MockRoom) AddPlayer(player interfaces.PlI) {
  m.AddedPlayer = player
}

func (m *MockRoom) RemovePlayer(player interfaces.PlI, msg string) {
  m.RemovedPlayer = player
}

func (m *MockRoom) GetExits() []interfaces.ExitI {
  return m.Exits
}

func (m *MockRoom) GetNPs() []interfaces.NPI {
  if (len(m.NPs) != 0) {
    return m.NPs
  }

  return []interfaces.NPI{
    &MockNP{},
  }
}

func (m *MockRoom) GetPlayers() []interfaces.PlI {
  return m.Players
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
