package main

import (
  "fmt"
  "log"
  "net"
  "sync/atomic"
  "time"

  "github.com/joelevering/gomud/player"
)

func watchIdle(p *player.Player, conn net.Conn, lastActivity *atomic.Int64, idle IdleDurations) {
  defer func() {
    if r := recover(); r != nil {
      log.Printf("Recovered from panic in watchIdle for %s: %v", p.GetName(), r)
    }
  }()

  ticker := time.NewTicker(idle.CheckInterval)
  defer ticker.Stop()
  warningsSent := 0

  for {
    select {
    case <-p.Logout:
      return
    case <-ticker.C:
      elapsed := time.Duration(time.Now().UnixNano() - lastActivity.Load())
      action, n := idleAction(elapsed, idle.WarnAfter, idle.WarnInterval, idle.KickAfter, warningsSent)

      switch action {
      case ActionWarn:
        warningsSent = n
        idleMin := int((idle.WarnAfter + time.Duration(n-1)*idle.WarnInterval).Minutes())
        remainMin := int(idle.KickAfter.Minutes()) - idleMin
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
