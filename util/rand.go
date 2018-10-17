package util

import (
  "time"

  "math/rand"
)

func RandF() float64 {
  rand.Seed(time.Now().UnixNano())
  return rand.Float64()
}
