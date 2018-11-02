package util

import (
  "time"

  "math/rand"
)

func RandF() float64 {
  rand.Seed(time.Now().UnixNano())
  return rand.Float64()
}

func RandI(min, max int) int {
  absRange := max - min
  if absRange == 0 {
    return min
  }

  rand.Seed(time.Now().UnixNano())
  return rand.Intn(absRange) + min
}
