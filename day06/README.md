### Part A

Unoptimized solution. For this solution I created a new slice at the start of each day and updated that. At the end of the day I would set the fish slice to the new slice and repeat the process.

~~~~
$ go test -count 5 -run '^$' -bench . -cpuprofile=v1.cpu.pprof > v1.txt
$ benchstat v1.txt
name      time/op
PartA-12   383µs ± 8%
PartB-12  0.35ns ± 0%
~~~~

I optimized my previous solution by instead mutating the same array, but keeping track of how many elements to iterate across each day.

~~~~
name      time/op
PartA-12   113µs ± 2%
PartB-12  0.35ns ± 0%
~~~~

It looks like a majority of the time being speint is now in the GC. This is likely from having to grow the slice multiple times each day.

### Part B
My "optimized" solution didn't cut it for part B. Instead I took an approach where I used an array to track the number of fishes by day. Each iteration would involve moving the value from one day to the next day, and special-casing day 0. I pre-allocated two arrays that I could use to keep track of the current state of the fish population and the next state of the fish of the fish population. This had the nice benefit of not having to allocate addiitonal memory while running.

~~~~
name      time/op
PartA-12   121µs ±16%
PartB-12  4.25µs ± 1%
~~~~

Benchstat shows that this new approach is much faster than the previous approach despite doing more iterations.