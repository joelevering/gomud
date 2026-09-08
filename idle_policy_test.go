package main

import "testing"
import "time"

func Test_IdleAction_NoneBeforeWarnThreshold(t *testing.T) {
  action, sent := idleAction(4*time.Minute, 5*time.Minute, 5*time.Minute, 20*time.Minute, 0)

  if action != ActionNone {
    t.Errorf("Expected ActionNone but got %v", action)
  }

  if sent != 0 {
    t.Errorf("Expected warningsSent to remain 0 but got %d", sent)
  }
}

func Test_IdleAction_FirstWarningAtThreshold(t *testing.T) {
  action, sent := idleAction(5*time.Minute, 5*time.Minute, 5*time.Minute, 20*time.Minute, 0)

  if action != ActionWarn {
    t.Errorf("Expected ActionWarn but got %v", action)
  }

  if sent != 1 {
    t.Errorf("Expected warningsSent to be 1 but got %d", sent)
  }
}

func Test_IdleAction_NoDuplicateWarningWithinSameInterval(t *testing.T) {
  action, sent := idleAction(7*time.Minute, 5*time.Minute, 5*time.Minute, 20*time.Minute, 1)

  if action != ActionNone {
    t.Errorf("Expected ActionNone but got %v", action)
  }

  if sent != 1 {
    t.Errorf("Expected warningsSent to remain 1 but got %d", sent)
  }
}

func Test_IdleAction_SecondWarningAtNextInterval(t *testing.T) {
  action, sent := idleAction(10*time.Minute, 5*time.Minute, 5*time.Minute, 20*time.Minute, 1)

  if action != ActionWarn {
    t.Errorf("Expected ActionWarn but got %v", action)
  }

  if sent != 2 {
    t.Errorf("Expected warningsSent to be 2 but got %d", sent)
  }
}

func Test_IdleAction_KickTakesPriorityAtThreshold(t *testing.T) {
  action, _ := idleAction(20*time.Minute, 5*time.Minute, 5*time.Minute, 20*time.Minute, 3)

  if action != ActionKick {
    t.Errorf("Expected ActionKick but got %v", action)
  }
}

func Test_IdleAction_DisabledWhenKickAfterNonPositive(t *testing.T) {
  action, _ := idleAction(999*time.Minute, 5*time.Minute, 5*time.Minute, 0, 0)

  if action != ActionNone {
    t.Errorf("Expected ActionNone when disabled but got %v", action)
  }
}
