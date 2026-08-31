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
  // should still get, since the flee attempt failed) ends combat. Since a
  // failed flee now leaves WantsFlee set (retrying is automatic, no
  // re-typing 'flee'), this is also what keeps the test from looping
  // forever instead of returning.
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

  if !pc.WantsFlee {
    t.Error("Expected a failed roll to leave WantsFlee set so combat keeps retrying, but it's false")
  }
}

func Test_StartAnnouncesTirednessOnOutOfStaminaFlee(t *testing.T) {
  pc := mocks.NewMockPlayer()
  pc.WantsFlee = true
  pc.FleeSucceeds = false
  pc.FleeOutOfStamina = true
  // Same trick as above: guarantees the npc's turn ends combat, proving it
  // still ran even though the pc's flee attempt failed.
  pc.ShouldDie = true
  npc := mocks.NewMockNP()
  rm := &mocks.MockRoom{}

  Start(pc, npc, rm)

  foundTiredMsg := false
  for _, msg := range rm.Messages {
    if strings.Contains(msg, "is too tired to flee!") {
      foundTiredMsg = true
    }
  }
  if !foundTiredMsg {
    t.Errorf("Expected an 'is too tired to flee!' message, but got %v", rm.Messages)
  }
}
