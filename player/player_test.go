package player

import (
  "fmt"
  "os"
  "math/rand"
  "strconv"
  "strings"
  "testing"
  "time"

  "github.com/joelevering/gomud/classes"
  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/mocks"
  "github.com/joelevering/gomud/room"
  "github.com/joelevering/gomud/storage"
)

func TestMain(m *testing.M) {
  os.Mkdir("../test", 0755)
  r := m.Run()
  os.RemoveAll("../test/")
  os.Exit(r)
}

func NewTestPlayer() (*Player, chan string, *mocks.MockQueue) {
  ch := make(chan string)
  q := &mocks.MockQueue{}
  s := storage.LoadStore(fmt.Sprintf("../test/%s.json", strconv.Itoa(rand.Intn(999999))))
  return NewPlayer(ch, q, s), ch, q
}

func Test_CmdSetsCombatEffectWithSkillName(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  p.EnterCombat()
  p.Cmd("bash")
  sk := p.GetCmbSkill()

  if sk == nil || sk.Name != "bash" {
    t.Errorf("Expected to get combat skill 'bash' after commanding 'bash', but got %v", sk)
  }
}

func Test_EnterRoom(t *testing.T) {
  p, ch, q := NewTestPlayer()
  defer close(ch)

  oldRoom := room.Room{Id: 99, Players: []interfaces.PlI{p}}
  p.Room = &oldRoom
  room := room.Room{Id: 101}

  p.EnterRoom(&room)

  if p.GetRoom().GetID() != 101 {
		t.Errorf("Expected player room to be set as %d but it was set as %d", room.GetID(), p.GetRoom().GetID())
  }

  if room.Players[0] != p {
		t.Errorf("Expected player to be the first of the room's pents")
  }

  if len(room.Players) != 1 {
		t.Errorf("Expected room to only have one player, but it has %d", len(room.Players))
  }

  if len(q.Pubs) != 1 || q.Pubs[0] != "pc-enters-101" {
    t.Errorf("Expected entering room to publish pc-enters-101, but it pub'd %s", q.Pubs[0])
  }
}

func Test_LeaveRoom(t *testing.T) {
  p, ch, q := NewTestPlayer()
  defer close(ch)

  oldRoom := room.Room{Id: 666, Players: []interfaces.PlI{p}}
  p.Room = &oldRoom

  p.LeaveRoom("")

  if len(oldRoom.Players) != 0 {
		t.Errorf("Expected oldRoom to have no players, but it has %d", len(oldRoom.Players))
  }

  if len(q.Pubs) != 1 || q.Pubs[0] != "pc-leaves-666" {
    t.Errorf("Expected entering room to publish pc-leaves-666, but it pub'd %s", q.Pubs[0])
  }
}

func Test_SendMsg(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)

  go p.SendMsg("testing SendMsg")

  res := <-ch

  if !strings.Contains(res, "testing SendMsg") {
		t.Error("Expected SendMsg to send 'testing SendMsg' to channel, but it didn't")
  }
}

func Test_List(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  room := &mocks.MockRoom{
    Players: []interfaces.PlI{
      &mocks.MockPlayer{},
    },
  }
  p.Room = room

  go p.List()

  res := <-ch

  // Sends preface
  if !strings.Contains(res, "You look around and see") {
		t.Errorf("Expected List to send 'You look around and see' to the player, but it sent %s", res)
  }

  // Lists players
  if !strings.Contains(res, "Yourself") {
		t.Errorf("Expected List to send 'Yourself' to the player, but it sent %s", res)
  }

  // Lists NPCs
  if !strings.Contains(res, "mock np name (NPC)") {
		t.Errorf("Expected List to send 'mock np name (NPC)' to the player, but it sent %s", res)
  }

  // Lists other players
  if !strings.Contains(res, "mock player") {
		t.Errorf("Expected List to send 'mock player' to the player, but it sent %s", res)
  }
}

func Test_Look(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  room := &mocks.MockRoom{
    Name: "Name",
    Exits: []interfaces.ExitI{
      &room.Exit{
        Desc: "You can go with this",
      },
      &room.Exit{
        Desc: "You can go with that",
      },
    },
  }
  p.Room = room

  go p.Look()

  res := <-ch
  if !strings.Contains(res, "~~Name~~") {
    t.Errorf("Expected room name 'Name' but got %s", res)
  }

  res = <-ch
  if !strings.Contains(res, "Desc") {
    t.Errorf("Expected room desc 'Desc' but got %s", res)
  }

  res = <-ch // Should be just time stamp (e.g. blank line sent)

  res = <-ch
  if !strings.Contains(res, "Exits:") {
    t.Errorf("Expected Look to send 'Exits:', but got %s", res)
  }

  res = <-ch
  if !strings.Contains(res, "You can go with this") {
    t.Errorf("Expected Look to send 'You can go with this', but got %s", res)
  }

  res = <-ch
  if !strings.Contains(res, "You can go with that") {
    t.Errorf("Expected Look to send 'You can go with that', but got %s", res)
  }

  res = <-ch // Should be just time stamp (e.g. blank line sent)

  res = <-ch
  if !strings.Contains(res, "You look around and see") {
    t.Errorf("Expected Look to send list, but got %s", res)
  }
}

func Test_LookNPWithNPName(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  room := &mocks.MockRoom{}
  p.Room = room

  go p.LookNP("mock np")

  res := <-ch
  if !strings.Contains(res, "You look at mock np name and see:") {
		t.Errorf("Expected 'You look at mock np name and see:' but got unexpected LookNPC result '%s'", res)
  }
  res = <-ch
  if !strings.Contains(res, "mock np desc") {
		t.Errorf("Expected 'mock np desc' but got unexpected LookNPC result '%s'", res)
  }
}

func Test_LookNPWithNoNP(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  room := &mocks.MockRoom{}
  p.Room = room

  go p.LookNP("missingno")

  res := <-ch
  if !strings.Contains(res, "Who are you looking at??") {
		t.Errorf("Expected 'Who are you looking at??' with unknown NPC, but got '%s'", res)
  }
}

func Test_Say(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  room := &mocks.MockRoom{}
  p.Room = room

  p.Say("testing Say")

  if !strings.Contains(room.Messages[0], "testing Say") {
		t.Error("Expected Say to send 'testing Say' to the room, but it didn't")
  }
}

func Test_Yell(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  adjacentRoom := &mocks.MockRoom{}
  room := &mocks.MockRoom{
		Exits: []interfaces.ExitI{
			&room.Exit{
				Room: adjacentRoom,
			},
		},
  }
  p.Room = room

  p.Yell("TESTING YELL")

  if !strings.Contains(adjacentRoom.Messages[0], "TESTING YELL") {
		t.Error("Expected Yell to send 'TESTING YELL' to adjacent rooms, but it didn't")
  }
}

func Test_MoveWithAccurateExitKey(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  adjacentRoom := &mocks.MockRoom{
		Name: "Adjacent Room",
  }
  room := &mocks.MockRoom{
		Exits: []interfaces.ExitI{
			&room.Exit{
				Room: adjacentRoom,
				Key:  "o",
			},
		},
  }
  p.Room = room

  go p.Move("o")

  res := <-ch

  if room.RemovedPlayer != p {
		t.Error("Expected player to be removed from initial room, but it was not")
  }

  if adjacentRoom.AddedPlayer != p {
		t.Error("Expected player to be added to adjacent room, but it was not")
  }

  if !strings.Contains(res, "~~Adjacent Room~~") {
		t.Errorf("Expected room name 'Name' but got %s", res)
  }

  // If the above test passes, assume it's 'Look'-ing and clear the channel before closing
  for i := 0; i < 5; i++ {
		res = <-ch
  }
}

func Test_MoveWithInaccurateExitKey(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  room := &mocks.MockRoom{}
  p.Room = room

  go p.Move("o")

  res := <-ch

  if !strings.Contains(res, "Where are you trying to go??") {
		t.Errorf("Expected 'Where are you trying to go??' with unknown move key, but got '%s'", res)
  }
}

func Test_LoseCombat(t *testing.T) {
  p, ch, _ := NewTestPlayer()

  origRoom := &mocks.MockRoom{ Name: "origin" }
  p.Room = origRoom

  spawn := &mocks.MockRoom{ Name: "spawn" }
  pc := p.Character
  pc.SetSpawn(spawn)
  pc.SetDet(1)

  go func (ch chan string) {
    defer close(ch)
    p.LoseCombat(p.Room.GetNPs()[0])
  }(ch)

  res := <-ch

  time.Sleep(1600 * time.Millisecond) // matches sleep in code

  if !strings.Contains(res, "You were defeated by mock np name") {
    t.Errorf("Expected 'You were defeated by mock np name' on death, but got '%s'", res)
  }

  if strings.Contains(res, "was defeated by mock char name") {
    t.Error("Expected player to not receive death notice to room, but it did")
  }

  if p.Room != spawn {
    t.Errorf("Expected to be moved to spawn on death but moved to '%s' instead", p.Room.GetName())
  }

  if pc.GetDet() != pc.GetMaxDet() {
    t.Error("Expected PC to be healed on combat loss/respawn, but it wasn't")
  }
}

func Test_WinCombatEndsCombatAndGivesExp(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)

  room := &mocks.MockRoom{}
  p.Room = room

  go func (ch chan string) {
    p.WinCombat(p.Room.GetNPs()[0])
  }(ch)

  res := <-ch

  if !strings.Contains(res, "You gained 2 experience!") { // hardcoded mock char exp
    t.Errorf("Expected 'You gained 2 experience' on defeating, but got '%s'", res)
  }
  if p.GetExp() != 2 {
    t.Errorf("Expected exp to be 2 but got %d", p.GetExp())
  }
  if !strings.Contains(res, "You need 8 more experience to level up.") {
    t.Errorf("Expected 'You need 8 more experience to level up' on defeating, but got '%s'", res)
  }
  if p.Level != 1 {
    t.Error("Expected PC not to level up, but it did")
  }
}

func Test_WinCombatLevelsUpPC(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  rm := &mocks.MockRoom{}
  p.Room = rm
  p.GainExp(p.NextLvlExp - 1)


  go func (ch chan string) {
    defer close(ch)
    p.WinCombat(p.Room.GetNPs()[0])
  }(ch)

  res := <-ch // Exp gain
  if !strings.Contains(res, "leveled up!") {
    t.Errorf("Expected 'leveled up!'' on defeating, but got '%s'", res)
  }

  res = <-ch // Level up
  if !strings.Contains(res, "You're now level 2!") {
    t.Errorf("Expected 'You're now level 2!'' on defeating, but got '%s'", res)
  }
}

func Test_ChangeClassResetsStats(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  p.Init()
  p.GainExp(p.NextLvlExp + 1)

  p.ChangeClass("athlete")

  if p.GetClassName() != classes.Athlete.GetName() {
    t.Errorf("Expected class name to be Athlete on change, but it's %s", p.GetClassName())
  }

  if p.GetLevel() != 1 {
    t.Errorf("Expected player level to be 1 after class change, but got %d", p.GetLevel())
  }

  if p.GetMaxDet() != 200  {
    t.Errorf("Expected Max Determination to reset to 200 on class change, but it's %d", p.GetMaxDet())
  }

  if p.GetDet() != 200 {
    t.Errorf("Expected Determination to lower to new max on class change, but it's %d", p.GetDet())
  }

  if p.GetExp() != 0 || p.GetNextLvlExp() != 10 {
    t.Errorf("Expected exp to reset to 0/10 on class change, but it's %d/%d", p.GetExp(), p.GetNextLvlExp())
  }
}

func Test_ChangeClassKeepsDet(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  p.Init()
  p.SetDet(33)

  p.ChangeClass("athlete")

  if p.GetDet() != 33 {
    t.Errorf("Expected determination to remain 33 on class change, but it's %d", p.GetDet())
  }
}

func Test_ClassSavesPersistAcrossLikeNames(t *testing.T) {
  p, ch, _ := NewTestPlayer()
  defer close(ch)
  name := "name"
  p.SetName(name)
  p.Init()
  p.GainExp(p.NextLvlExp)
  p.ChangeClass("athlete") // to persist conscript

  ch = make(chan string)
  p2 := NewPlayer(ch, p.Queue, p.Store)
  defer close(ch)
  p2.SetName(name)
  p2.Init()

  if p2.Level != 2 {
    t.Errorf("Expected player class level to persist across like-named characters, but level is %d", p2.GetLevel())
  }
}
