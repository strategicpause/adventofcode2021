Inspired by Felix Geisendörfer's solution at https://felixge.de/2021/12/01/advent-of-go-profiling-2021-day-1-1/. This has mostly been an exercise in understanding what tools are available for profiling code in go.

~~~~
$ go test -count 5 -run '^$' -bench . -cpuprofile=v1.cpu.pprof > v1.txt
$ benchstat v1.txt
name      time/op
PartA-12  97.1ns ± 3%
~~~~
