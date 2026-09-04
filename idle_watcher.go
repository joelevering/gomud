package main

import (
  "fmt"
  "log"
  "net"
  "sync/atomic"
  "time"

  "github.com/joelevering/gomud/player"
)

func watchIdle(p *player.Player, conn net.Conn, lastActivity *atomic.Int64, warnAfter, warnInterval, kickAfter, checkInterval time.Duration) {
  defer func() {
    if r := recover(); r != nil {
      log.Printf("Recovered from panic in watchIdle for %s: %v", p.GetName(), r)
    }
  }()

  ticker := time.NewTicker(checkInterval)
  defer ticker.Stop()
  warningsSent := 0

  for {
    select {
    case <-p.Logout:
      return
    case <-ticker.C:
      elapsed := time.Duration(time.Now().UnixNano() - lastActivity.Load())
      action, n := idleAction(elapsed, warnAfter, warnInterval, kickAfter, warningsSent)

      switch action {
      case ActionWarn:
        warningsSent = n
        idleMin := int((warnAfter + time.Duration(n-1)*warnInterval).Minutes())
        remainMin := int(kickAfter.Minutes()) - idleMin
        p.SendMsg(fmt.Sprintf("You've been idle for %d minutes. You will be disconnected in %d minutes if you remain idle.", idleMin, remainMin))
      case ActionKick:
        log.Printf("Kicking idle player: %s", p.GetName())
        p.SendMsg("You've been idle too long and are being disconnected.")
        time.Sleep(1500 * time.Millisecond)
        conn.Close()
        return
      }
    }
  }
}
