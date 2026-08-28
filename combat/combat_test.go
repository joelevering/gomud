package combat

import (
  "strings"
  "testing"

  "github.com/joelevering/gomud/mocks"
)

func Test_StartEndsOnSuccessfulFlee(t *testing.T) {
  pc := mocks.NewMockPlayer()
  pc.WantsFlee = true
  pc.FleeSucceeds = true
  npc := mocks.NewMockNP()
  rm := &mocks.MockRoom{}

  Start(pc, npc, rm)

  found := false
  for _, msg := range rm.Messages {
    if strings.Contains(msg, "flees from the fight") {
      found = true
    }
  }
  if !found {
    t.Errorf("Expected a 'flees from the fight' message, but got %v", rm.Messages)
  }

  if !pc.Fled {
    t.Error("Expected pc.Flee to have been called, but Fled is false")
  }

  if !npc.LeftCombat {
    t.Error("Expected npc to have left combat, but LeftCombat is false")
  }
}

func Test_StartContinuesOnFailedFlee(t *testing.T) {
  pc := mocks.NewMockPlayer()
  pc.WantsFlee = true
  pc.FleeSucceeds = false
  // Marking pc as already "defeated" means npc's very next turn (which it
  // should still get, since the flee attempt failed) ends combat. If a
  // failed flee incorrectly skipped the npc's turn, this test would hang
  // instead of returning, since nothing else would ever end the loop.
  pc.ShouldDie = true
  npc := mocks.NewMockNP()
  rm := &mocks.MockRoom{}

  Start(pc, npc, rm)

  foundFailMsg := false
  for _, msg := range rm.Messages {
    if strings.Contains(msg, "tries to flee, but fails!") {
      foundFailMsg = true
    }
  }
  if !foundFailMsg {
    t.Errorf("Expected a 'tries to flee, but fails!' message, but got %v", rm.Messages)
  }

  if pc.WantsFlee {
    t.Error("Expected the failed attempt to clear WantsFlee, but it's still true")
  }
}
