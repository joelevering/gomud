package main

import (
  "os"
  "path/filepath"
  "testing"
)

func Test_LoadConfiguration_MissingFileFallsBackToDefaults(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")

  cfg := LoadConfiguration(path)

  if *cfg != *DefaultConfiguration() {
    t.Errorf("Expected defaults but got %+v", cfg)
  }
}

func Test_LoadConfiguration_EmptyFileFallsBackToDefaults(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")
  if err := os.WriteFile(path, []byte(""), 0644); err != nil {
    t.Fatal(err)
  }

  cfg := LoadConfiguration(path)

  if *cfg != *DefaultConfiguration() {
    t.Errorf("Expected defaults but got %+v", cfg)
  }
}

func Test_LoadConfiguration_PartialOverrideKeepsDefaultsForOmittedFields(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")
  if err := os.WriteFile(path, []byte(`{"idle":{"kick_after_minutes":30}}`), 0644); err != nil {
    t.Fatal(err)
  }

  cfg := LoadConfiguration(path)

  if cfg.Idle.KickAfterMinutes != 30 {
    t.Errorf("Expected KickAfterMinutes to be 30 but got %d", cfg.Idle.KickAfterMinutes)
  }

  if cfg.Idle.WarnAfterMinutes != 5 {
    t.Errorf("Expected WarnAfterMinutes to still default to 5 but got %d", cfg.Idle.WarnAfterMinutes)
  }

  if cfg.Idle.WarnIntervalMinutes != 5 {
    t.Errorf("Expected WarnIntervalMinutes to still default to 5 but got %d", cfg.Idle.WarnIntervalMinutes)
  }

  if cfg.Idle.CheckIntervalSeconds != 15 {
    t.Errorf("Expected CheckIntervalSeconds to still default to 15 but got %d", cfg.Idle.CheckIntervalSeconds)
  }

  if cfg.DefaultRoomID != 15 {
    t.Errorf("Expected DefaultRoomID to still default to 15 but got %d", cfg.DefaultRoomID)
  }
}

func Test_LoadConfiguration_MalformedJSONFallsBackToDefaults(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")
  if err := os.WriteFile(path, []byte(`{not valid json`), 0644); err != nil {
    t.Fatal(err)
  }

  cfg := LoadConfiguration(path)

  if *cfg != *DefaultConfiguration() {
    t.Errorf("Expected defaults but got %+v", cfg)
  }
}

func Test_LoadConfiguration_DefaultRoomIDOverrideIsRespected(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")
  if err := os.WriteFile(path, []byte(`{"default_room_id":9}`), 0644); err != nil {
    t.Fatal(err)
  }

  cfg := LoadConfiguration(path)

  if cfg.DefaultRoomID != 9 {
    t.Errorf("Expected DefaultRoomID to be 9 but got %d", cfg.DefaultRoomID)
  }
}

func Test_LoadConfiguration_InvalidDefaultRoomIDFallsBackToDefault(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")
  if err := os.WriteFile(path, []byte(`{"default_room_id":-1}`), 0644); err != nil {
    t.Fatal(err)
  }

  cfg := LoadConfiguration(path)

  if cfg.DefaultRoomID != 15 {
    t.Errorf("Expected DefaultRoomID to fall back to 15 but got %d", cfg.DefaultRoomID)
  }
}

func Test_LoadConfiguration_InvalidCheckIntervalFallsBackToDefaults(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")
  if err := os.WriteFile(path, []byte(`{"idle":{"check_interval_seconds":0}}`), 0644); err != nil {
    t.Fatal(err)
  }

  cfg := LoadConfiguration(path)

  if *cfg != *DefaultConfiguration() {
    t.Errorf("Expected defaults but got %+v", cfg)
  }
}
