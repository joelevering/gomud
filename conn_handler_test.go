package main

import (
  "fmt"
  "net"
  "path/filepath"
  "strings"
  "testing"

  "github.com/joelevering/gomud/mocks"
  "github.com/joelevering/gomud/player"
  "github.com/joelevering/gomud/storage"
)

func Test_ConfirmName_RetriesWhenNameIsTaken(t *testing.T) {
  serverConn, clientConn := net.Pipe()
  defer serverConn.Close()
  defer clientConn.Close()

  ch := make(chan string)
  q := &mocks.MockQueue{}
  s := storage.LoadStore(filepath.Join(t.TempDir(), "store.json"))
  p := player.NewPlayer(ch, q, s)

  claimName := make(chan nameClaimRequest)
  go func() {
    first := <-claimName
    first.reply <- false // "Bob" is already logged in

    second := <-claimName
    second.reply <- true // "Alice" is free
  }()

  go func() {
    for _, line := range []string{"Bob", "Y", "Alice", "Y"} {
      fmt.Fprintln(clientConn, line)
    }
  }()

  var who string
  var ok bool
  go func(ch chan string) {
    defer close(ch)
    who, ok = confirmName(p, serverConn, claimName)
  }(ch)

  var messages []string
  for msg := range ch {
    messages = append(messages, msg)
  }

  if !ok {
    t.Fatal("Expected confirmName to return ok=true once a name is successfully claimed")
  }

  if who != "Alice" {
    t.Errorf("Expected confirmName to return 'Alice' but got %q", who)
  }

  found := false
  for _, msg := range messages {
    if strings.Contains(msg, "already logged in") {
      found = true
    }
  }

  if !found {
    t.Errorf("Expected a rejection message for the taken name, got messages: %v", messages)
  }
}

func Test_ConfirmName_ReturnsFalseWhenClientDisconnectsDuringPrompt(t *testing.T) {
  serverConn, clientConn := net.Pipe()
  defer serverConn.Close()
  clientConn.Close() // client disconnects before ever answering "Who are you?"

  ch := make(chan string)
  q := &mocks.MockQueue{}
  s := storage.LoadStore(filepath.Join(t.TempDir(), "store.json"))
  p := player.NewPlayer(ch, q, s)

  claimName := make(chan nameClaimRequest)

  var who string
  var ok bool
  go func(ch chan string) {
    defer close(ch)
    who, ok = confirmName(p, serverConn, claimName)
  }(ch)

  for range ch {
    // drain SendMsg output until confirmName gives up
  }

  if ok {
    t.Errorf("Expected confirmName to return ok=false on disconnect, but got ok=true (who=%q)", who)
  }

  if who != "" {
    t.Errorf("Expected an empty name on disconnect but got %q", who)
  }
}
