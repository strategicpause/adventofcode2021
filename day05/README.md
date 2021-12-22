This is a pretty slow solution without any optimizations. I'll likely need to use a different algorithm to optimize this solution.

~~~~
$ go test -count 5 -run '^$' -bench . -cpuprofile=v1.cpu.pprof > v1.txt
$ benchstat v1.txt
name      time/op
PartA-12  5.40µs ± 4%
PartB-12  9.87µs ± 2%
~~~~