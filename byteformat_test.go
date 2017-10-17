package main

import (
  "testing"
)

func TestToBytes(t *testing.T) {
  var humanString string = "1 KB"
  if ToBytes(humanString) != 1000 {
    t.Errorf("Could not convert %s", humanString)
  }
  humanString = "1 KiB"
  if ToBytes("1 KiB") != 1024 {
    t.Errorf("Could not convert %s", humanString)
  }
}

func TestToUnit(t *testing.T) {
  var bytes uint64 = 1234
  var unit string = "KB"
  var expected float64 = 1.234
  if ToUnit(bytes, unit) != expected {
    t.Errorf("Could not convert %v to %s", bytes, unit)
  }
  unit = "KiB"
  expected = 1.205078125
  if ToUnit(bytes, unit) != expected {
    t.Errorf("Could not convert %v to %s", bytes, unit)
  }
}
