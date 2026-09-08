package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "time"
)

const Config string = "config.json"

type Configuration struct {
  DefaultRoomID int        `json:"default_room_id"`
  Idle          IdleConfig `json:"idle"`
}

type IdleConfig struct {
  WarnAfterMinutes     int `json:"warn_after_minutes"`
  WarnIntervalMinutes  int `json:"warn_interval_minutes"`
  KickAfterMinutes     int `json:"kick_after_minutes"`
  CheckIntervalSeconds int `json:"check_interval_seconds"`
}

type IdleDurations struct {
  WarnAfter     time.Duration
  WarnInterval  time.Duration
  KickAfter     time.Duration
  CheckInterval time.Duration
}

func (c IdleConfig) Durations() IdleDurations {
  return IdleDurations{
    WarnAfter:     time.Duration(c.WarnAfterMinutes) * time.Minute,
    WarnInterval:  time.Duration(c.WarnIntervalMinutes) * time.Minute,
    KickAfter:     time.Duration(c.KickAfterMinutes) * time.Minute,
    CheckInterval: time.Duration(c.CheckIntervalSeconds) * time.Second,
  }
}

func DefaultConfiguration() *Configuration {
  return &Configuration{
    DefaultRoomID: 15,
    Idle: IdleConfig{
      WarnAfterMinutes:     5,
      WarnIntervalMinutes:  5,
      KickAfterMinutes:     20,
      CheckIntervalSeconds: 15,
    },
  }
}

// LoadConfiguration reads filename if it exists and returns the resulting
// Configuration. A missing or empty file is not an error -- defaults are
// used and logged. Any config.json that *does* specify values is expected
// to specify valid ones; a malformed file or an out-of-range setting is
// treated as a misconfiguration and returned as an error so the caller can
// refuse to start rather than silently running with different settings
// than the operator intended.
func LoadConfiguration(filename string) (*Configuration, error) {
  if _, err := os.Stat(filename); os.IsNotExist(err) {
    log.Printf("%s not found, using default configuration", filename)
    return DefaultConfiguration(), nil
  }

  b, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, fmt.Errorf("reading %s: %w", filename, err)
  }

  if len(b) == 0 {
    log.Printf("%s is empty, using default configuration", filename)
    return DefaultConfiguration(), nil
  }

  cfg := DefaultConfiguration()
  if err := json.Unmarshal(b, cfg); err != nil {
    return nil, fmt.Errorf("parsing %s: %w", filename, err)
  }

  if cfg.Idle.WarnAfterMinutes <= 0 || cfg.Idle.WarnIntervalMinutes <= 0 || cfg.Idle.KickAfterMinutes <= cfg.Idle.WarnAfterMinutes || cfg.Idle.CheckIntervalSeconds <= 0 {
    return nil, fmt.Errorf("invalid idle config in %s: %+v", filename, cfg.Idle)
  }

  if cfg.DefaultRoomID <= 0 {
    return nil, fmt.Errorf("invalid default_room_id in %s: %d", filename, cfg.DefaultRoomID)
  }

  return cfg, nil
}
