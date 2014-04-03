raidcheck [![Build Status](https://drone.io/github.com/PreetamJinka/raidcheck/status.png)](https://drone.io/github.com/PreetamJinka/raidcheck/latest)
===
A (tiny) Go package to detect degraded RAID arrays using `/proc/mdstat`.

Usage
---
```go
degraded, err := raidcheck.Degraded()
```

License
---
MIT
