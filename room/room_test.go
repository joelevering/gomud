package room

import (
  "fmt"
  "sync"
  "testing"

  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/mocks"
)

type namedPlayer struct {
  *mocks.MockPlayer
  name string
}

func (n *namedPlayer) GetName() string { return n.name }

func newNamedPlayer(name string) interfaces.PlI {
  return &namedPlayer{MockPlayer: mocks.NewMockPlayer(), name: name}
}

func Test_RemovePlayer(t *testing.T) {
  p1 := newNamedPlayer("one")
  p2 := newNamedPlayer("two")
  rm := &Room{Players: []interfaces.PlI{p1, p2}}

  rm.RemovePlayer(p1, "one left")

  if len(rm.Players) != 1 {
    t.Fatalf("Expected 1 player left, got %d", len(rm.Players))
  }

  if rm.Players[0] != p2 {
    t.Errorf("Expected remaining player to be p2")
  }
}

func Test_ConcurrentAddRemovePlayers_NoRace(t *testing.T) {
  rm := &Room{}

  var wg sync.WaitGroup
  players := make([]interfaces.PlI, 50)
  for i := range players {
    players[i] = newNamedPlayer(fmt.Sprintf("player-%d", i))
  }

  for _, pl := range players {
    wg.Add(1)
    go func(pl interfaces.PlI) {
      defer wg.Done()
      rm.AddPlayer(pl)
    }(pl)
  }
  wg.Wait()

  if len(rm.GetPlayers()) != len(players) {
    t.Fatalf("Expected %d players after concurrent adds, got %d", len(players), len(rm.GetPlayers()))
  }

  for _, pl := range players {
    wg.Add(1)
    go func(pl interfaces.PlI) {
      defer wg.Done()
      rm.RemovePlayer(pl, "left")
    }(pl)
  }

  for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      rm.GetPlayers()
      rm.Message("noise")
    }()
  }

  wg.Wait()

  if len(rm.GetPlayers()) != 0 {
    t.Errorf("Expected 0 players after concurrent removes, got %d", len(rm.GetPlayers()))
  }
}
