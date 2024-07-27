// ex01 Focuses on profiling and benchmarking
// Using profiled data I noticed that we were dry running QuickSort even when array was sorted
// so I only fixed dry runs in MinCoins2Optimized and it resultet in 2-3 times performance win.
//
// # profiling
//
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
//
// * profile visualization
//
// * for web:
// go toos pprof -web path/to/binary path/to/profiling/data/by/go/test
//
// * for cli:
// go toos pprof -text path/to/binary path/to/profiling/data/by/go/test
//
// * to get longest run N functions:
// go toos pprof -text -nodecount=N path/to/binary path/to/profiling/data/by/go/test
//
// * you can pipe cli output to a file
package ex01
