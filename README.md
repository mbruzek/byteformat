# byteformat

This is a Go package that converts and formats large binary (GiB)
and decimal (GB) byte values to human readable format.

It would be useful to converting human readable values to bytes and
back to other orders of magnitude.

## Usage:

To convert a human readable string to bytes:
```
byteformat.ToBytes("1 KB")     // 1000
byteformat.ToBytes("1 KiB")    // 1024
byteformat.ToBytes("500MB")    // 500000000
byteformat.ToBytes("500MiB")   // 524288000
byteformat.ToBytes("1.21GB")   // 1210000000
byteformat.ToBytes("1.21GiB")  // 1299227607.04
```

To convert a value to a specific unit:
```
byteformat.ToUnit(1234, "KB")     // 1.234
byteformat.ToUnit(1234, "KiB")    // 1.205078125
byteformat.ToUnit(567890, "MB")   // 0.56789
byteformat.ToUnit(567890, "MiB")  // 0.541582108
```

## Notice
I am writing this from a TDD perspective, the code is not yet complete.
