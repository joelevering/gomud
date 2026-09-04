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
  go func(ch chan string) {
    defer close(ch)
    who = confirmName(p, serverConn, claimName)
  }(ch)

  var messages []string
  for msg := range ch {
    messages = append(messages, msg)
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
