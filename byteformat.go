package main

// NOTE: This code is written from a TDD perspective and is not yet complete.

import (
  "fmt"
  "strconv"
  "strings"
)

func ToBytes(human_readable string) float64 {
  var result float64
  var parts = strings.Split(human_readable, " ")
  f, _ := strconv.ParseFloat(parts[0], 64)
  if parts[1] == "KB" {
    result = f * float64(1000)
  } else if parts[1] == "KiB" {
    result = f * float64(1024)
  }
  return result
}

func ToUnit(bytes uint64, unit string) float64 {
  var result float64
  var binary float64 = 1024
  var decimal float64 = 1000
  if unit == "KB" {
    result = float64(bytes) / decimal
  } else if unit == "KiB" {
    result = float64(bytes) / binary
  }
  return result
}

func main() {
  fmt.Println("Byte format")
}
