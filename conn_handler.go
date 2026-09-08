package main

import (
  "bufio"
  "fmt"
  "net"
  "strings"
  "sync/atomic"
  "time"

  "github.com/joelevering/gomud/player"
)

type ConnHandler struct {
  entering    chan *player.Player
  leaving     chan *player.Player
  claimName   chan nameClaimRequest
  releaseName chan string
  state       *GameState
  idle        IdleDurations
}

func (handler *ConnHandler) Handle(conn net.Conn) {
  defer conn.Close()

  ch := make(chan string)
  p := player.NewPlayer(ch, handler.state.Queue, handler.state.Store)
  go p.StartWriter(conn)

  name, ok := confirmName(p, conn, handler.claimName)
  if !ok {
    close(ch)
    return
  }
  p.SetName(name)

  reachedLogin := false
  defer func() {
    if !reachedLogin {
      handler.releaseName <- name
    }
  }()

  p.Init()
  handler.entering <- p
  reachedLogin = true

  lastActivity := &atomic.Int64{}
  lastActivity.Store(time.Now().UnixNano())
  go watchIdle(p, conn, lastActivity, handler.idle)

  input := bufio.NewScanner(conn)
  for input.Scan() {
    lastActivity.Store(time.Now().UnixNano())
    txt := input.Text()
    if txt == "exit" || txt == "quit" {
      p.SendMsg("Are you sure you want to quit? ('Y' to confirm)")
      input.Scan()
      if strings.ToUpper(input.Text()) == "Y" {
        p.SendMsg("OK! See you next time!")
        break
      }

      p.SendMsg("OK, keeping you logged in. ('Y' would have logged you out)")
    } else {
      p.Cmd(txt)
    }
  }

  handler.leaving <- p
}

// confirmName prompts the connection for a name and claims it, retrying on
// a rejected (already-logged-in) name. It returns ok=false if the
// connection disconnects before a name is successfully claimed -- Scan()
// returns false immediately (non-blocking) once the underlying conn hits
// EOF/an error, so this must be checked rather than assumed true, or a
// disconnect during the prompt spins the loop forever.
func confirmName(p *player.Player, conn net.Conn, claimName chan<- nameClaimRequest) (string, bool) {
  var confirmed, who string
  input := bufio.NewScanner(conn)

  for {
    for strings.ToUpper(confirmed) != "Y" {
      p.SendMsg("Who are you?")
      if !input.Scan() {
        return "", false
      }
      who = input.Text()

      p.SendMsg(fmt.Sprintf("Are you sure you want to be called \"%s\"? ('Y' to confirm)", who))
      if !input.Scan() {
        return "", false
      }
      confirmed = input.Text()
    }

    reply := make(chan bool)
    claimName <- nameClaimRequest{name: who, reply: reply}
    if <-reply {
      return who, true
    }

    p.SendMsg(fmt.Sprintf("Sorry, \"%s\" is already logged in. Please choose a different name.", who))
    confirmed = ""
  }
}
