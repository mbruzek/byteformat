package main

import (
  "testing"
)

func TestToBytes(t *testing.T) {
  if ToBytes("500MB") != 500000000 {
    t.Error("Could not convert 500MB")
  }
  if ToBytes("500MiB") != 524288000 {
    t.Error("Could not convert 500MiB")
  }
  if ToBytes("1.21GiB") != 1299227607.04 {
    t.Errorf("Could not convert 1.21GiB")
  }
  if ToBytes("1.21GB") != 1210000000 {
    t.Errorf("Could not convert 1.21GB")
  }
}

func TestToUnit(t *testing.T) {
  if ToUnit(1234, "KB") != 1.234 {
    t.Errorf("Could not convert 1234 to KB")
  }
  if ToUnit(1234, "KiB") != 1.204078 {
    t.Errorf("Could not convert 1234 to KiB")
  }
  if ToUnit(567890, "MB") != 0.56789 {
    t.Errorf("Could not convert 567890 to MB")
  }
  if ToUnit(567890, "MiB") != 0.541582108 {
    t.Errorf("Could not convert 567890 to MiB")
  }
}
