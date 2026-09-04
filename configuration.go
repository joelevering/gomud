package main

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "os"
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

func LoadConfiguration(filename string) *Configuration {
  f, err := os.OpenFile(filename, os.O_CREATE, 0644)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  b, err := ioutil.ReadAll(f)
  if err != nil {
    panic(err)
  }

  cfg := DefaultConfiguration()
  if len(b) == 0 {
    return cfg
  }

  if err := json.Unmarshal(b, cfg); err != nil {
    log.Printf("Error parsing %s, falling back to defaults: %v", filename, err)
    return DefaultConfiguration()
  }

  if cfg.Idle.WarnAfterMinutes <= 0 || cfg.Idle.WarnIntervalMinutes <= 0 || cfg.Idle.KickAfterMinutes <= cfg.Idle.WarnAfterMinutes || cfg.Idle.CheckIntervalSeconds <= 0 {
    log.Printf("Invalid idle config in %s, falling back to defaults", filename)
    cfg.Idle = DefaultConfiguration().Idle
  }

  if cfg.DefaultRoomID <= 0 {
    cfg.DefaultRoomID = DefaultConfiguration().DefaultRoomID
  }

  return cfg
}
