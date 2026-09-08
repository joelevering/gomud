package main

import "time"

type IdleAction int

const (
  ActionNone IdleAction = iota
  ActionWarn
  ActionKick
)

// warningsSent tracks how many warning "slots" have already fired, so the
// same warning isn't resent every tick within one interval.
func idleAction(elapsed, warnAfter, warnInterval, kickAfter time.Duration, warningsSent int) (IdleAction, int) {
  if kickAfter <= 0 || warnAfter <= 0 || warnInterval <= 0 {
    return ActionNone, warningsSent
  }

  if elapsed >= kickAfter {
    return ActionKick, warningsSent
  }

  if elapsed < warnAfter {
    return ActionNone, warningsSent
  }

  slot := int((elapsed-warnAfter)/warnInterval) + 1
  if slot > warningsSent {
    return ActionWarn, slot
  }

  return ActionNone, warningsSent
}
