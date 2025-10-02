```
$ go run -race main.go

Number not found
==================
WARNING: DATA RACE
Write at 0x00c00019c0f0 by goroutine 7:
  runtime.mapaccess2_fast64()
      /opt/homebrew/Cellar/go/1.25.1/libexec/src/internal/runtime/maps/runtime_fast64_swiss.go:86 +0x25c
  main.updateData()
      /Users/herlon/dev/go-concurrent-map-race/main.go:26 +0x50
  main.main.gowrap1()
      /Users/herlon/dev/go-concurrent-map-race/main.go:16 +0x34

Previous read at 0x00c00019c0f0 by goroutine 8:
  runtime.mapaccess1_fast64()
      /opt/homebrew/Cellar/go/1.25.1/libexec/src/internal/runtime/maps/runtime_fast64_swiss.go:17 +0x24c
  main.readData()
      /Users/herlon/dev/go-concurrent-map-race/main.go:35 +0x40
  main.main.gowrap2()
      /Users/herlon/dev/go-concurrent-map-race/main.go:17 +0x34

Goroutine 7 (running) created at:
  main.main()
      /Users/herlon/dev/go-concurrent-map-race/main.go:16 +0xac

Goroutine 8 (running) created at:
  main.main()
      /Users/herlon/dev/go-concurrent-map-race/main.go:17 +0x34
==================
Number not found
fatal error: concurrent map read and map write
```