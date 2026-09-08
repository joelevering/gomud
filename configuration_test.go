package main

import (
  "os"
  "path/filepath"
  "testing"
)

func Test_LoadConfiguration_MissingFileFallsBackToDefaults(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")

  cfg, err := LoadConfiguration(path)
  if err != nil {
    t.Fatalf("Expected no error but got %v", err)
  }

  if *cfg != *DefaultConfiguration() {
    t.Errorf("Expected defaults but got %+v", cfg)
  }
}

func Test_LoadConfiguration_MissingFileIsNotCreated(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")

  if _, err := LoadConfiguration(path); err != nil {
    t.Fatalf("Expected no error but got %v", err)
  }

  if _, err := os.Stat(path); !os.IsNotExist(err) {
    t.Error("Expected LoadConfiguration to not create config.json when it's absent")
  }
}

func Test_LoadConfiguration_EmptyFileFallsBackToDefaults(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")
  if err := os.WriteFile(path, []byte(""), 0644); err != nil {
    t.Fatal(err)
  }

  cfg, err := LoadConfiguration(path)
  if err != nil {
    t.Fatalf("Expected no error but got %v", err)
  }

  if *cfg != *DefaultConfiguration() {
    t.Errorf("Expected defaults but got %+v", cfg)
  }
}

func Test_LoadConfiguration_PartialOverrideKeepsDefaultsForOmittedFields(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")
  if err := os.WriteFile(path, []byte(`{"idle":{"kick_after_minutes":30}}`), 0644); err != nil {
    t.Fatal(err)
  }

  cfg, err := LoadConfiguration(path)
  if err != nil {
    t.Fatalf("Expected no error but got %v", err)
  }

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

func Test_LoadConfiguration_MalformedJSONErrors(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")
  if err := os.WriteFile(path, []byte(`{not valid json`), 0644); err != nil {
    t.Fatal(err)
  }

  if _, err := LoadConfiguration(path); err == nil {
    t.Error("Expected an error for malformed JSON but got nil")
  }
}

func Test_LoadConfiguration_DefaultRoomIDOverrideIsRespected(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")
  if err := os.WriteFile(path, []byte(`{"default_room_id":9}`), 0644); err != nil {
    t.Fatal(err)
  }

  cfg, err := LoadConfiguration(path)
  if err != nil {
    t.Fatalf("Expected no error but got %v", err)
  }

  if cfg.DefaultRoomID != 9 {
    t.Errorf("Expected DefaultRoomID to be 9 but got %d", cfg.DefaultRoomID)
  }
}

func Test_LoadConfiguration_InvalidDefaultRoomIDErrors(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")
  if err := os.WriteFile(path, []byte(`{"default_room_id":-1}`), 0644); err != nil {
    t.Fatal(err)
  }

  if _, err := LoadConfiguration(path); err == nil {
    t.Error("Expected an error for an invalid default_room_id but got nil")
  }
}

func Test_LoadConfiguration_InvalidCheckIntervalErrors(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")
  if err := os.WriteFile(path, []byte(`{"idle":{"check_interval_seconds":0}}`), 0644); err != nil {
    t.Fatal(err)
  }

  if _, err := LoadConfiguration(path); err == nil {
    t.Error("Expected an error for an invalid check_interval_seconds but got nil")
  }
}

func Test_LoadConfiguration_KickAfterNotGreaterThanWarnAfterErrors(t *testing.T) {
  path := filepath.Join(t.TempDir(), "config.json")
  if err := os.WriteFile(path, []byte(`{"idle":{"warn_after_minutes":10,"kick_after_minutes":5}}`), 0644); err != nil {
    t.Fatal(err)
  }

  if _, err := LoadConfiguration(path); err == nil {
    t.Error("Expected an error when kick_after_minutes <= warn_after_minutes but got nil")
  }
}
