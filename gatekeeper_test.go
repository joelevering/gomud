package main

import (
  "testing"

  "github.com/joelevering/gomud/player"
)

func Test_IsNameAvailable_TrueWhenFree(t *testing.T) {
  gk := &Gatekeeper{
    state:        &GameState{Players: map[string]*player.Player{}},
    pendingNames: map[string]bool{},
  }

  if !gk.isNameAvailable("Alice") {
    t.Error("Expected 'Alice' to be available")
  }
}

func Test_IsNameAvailable_FalseWhenLoggedIn(t *testing.T) {
  gk := &Gatekeeper{
    state:        &GameState{Players: map[string]*player.Player{"Alice": nil}},
    pendingNames: map[string]bool{},
  }

  if gk.isNameAvailable("Alice") {
    t.Error("Expected 'Alice' to be unavailable while logged in")
  }
}

func Test_IsNameAvailable_FalseWhenReserved(t *testing.T) {
  gk := &Gatekeeper{
    state:        &GameState{Players: map[string]*player.Player{}},
    pendingNames: map[string]bool{"Alice": true},
  }

  if gk.isNameAvailable("Alice") {
    t.Error("Expected 'Alice' to be unavailable while reserved")
  }
}

func Test_KeepTheGate_ClaimNameIsAtomicAndReleasable(t *testing.T) {
  claimName := make(chan nameClaimRequest)
  releaseName := make(chan string)

  gk := &Gatekeeper{
    entering:     make(chan *player.Player),
    leaving:      make(chan *player.Player),
    claimName:    claimName,
    releaseName:  releaseName,
    state:        &GameState{Players: map[string]*player.Player{}},
    pendingNames: map[string]bool{},
  }

  go gk.KeepTheGate()

  reply1 := make(chan bool)
  claimName <- nameClaimRequest{name: "Alice", reply: reply1}
  if !<-reply1 {
    t.Fatal("Expected first claim on 'Alice' to succeed")
  }

  reply2 := make(chan bool)
  claimName <- nameClaimRequest{name: "Alice", reply: reply2}
  if <-reply2 {
    t.Fatal("Expected second claim on 'Alice' to fail while first is pending")
  }

  releaseName <- "Alice"

  reply3 := make(chan bool)
  claimName <- nameClaimRequest{name: "Alice", reply: reply3}
  if !<-reply3 {
    t.Fatal("Expected claim on 'Alice' to succeed again after release")
  }
}
