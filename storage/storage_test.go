package storage

import (
  "fmt"
  "testing"
)

// Aliases must survive an actual process restart, not just being shared
// in-memory between two Player instances pointed at the same *Storage. This
// exercises the real write-to-disk / read-from-disk path via a fresh
// LoadStore() call, which catches JSON tag/field mistakes that in-memory
// sharing wouldn't.
func Test_AliasesSurviveDiskRoundTrip(t *testing.T) {
  filename := fmt.Sprintf("%s/store.json", t.TempDir())
  s := LoadStore(filename)
  pID := "roundtrip-player"
  s.InitPlayerData(pID)

  aliases := map[string]string{"gs": "attack slime shove", "l": "look"}
  s.PersistAliases(pID, aliases)

  reloaded := LoadStore(filename)

  got := reloaded.LoadAliases(pID)
  if got["gs"] != "attack slime shove" || got["l"] != "look" {
    t.Errorf("Expected aliases to survive a fresh LoadStore() from disk, but got %v", got)
  }
}
