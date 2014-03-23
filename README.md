raidcheck
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
