# profiling
recording profiling info command:
```bash
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
```

# profile visualization
```bash
# for web
go toos pprof -web path/to/binary path/to/profiling/data/by/go/test

# for cli 
go toos pprof -text path/to/binary path/to/profiling/data/by/go/test

# to get longest run N functions
go toos pprof -text -nodecount=N path/to/binary path/to/profiling/data/by/go/test

# you can pipe cli output to a file
```

