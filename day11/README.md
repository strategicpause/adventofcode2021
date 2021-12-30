My first approach used a queue to determine the visit order of the flashing octupus. For this approach I pushed a struct consisting of the (X, Y) pairs to visit. Since I had to both create a new struct and potentially expand the slice, this approach resulted putting a lot of strain on the GC and therefore resulted in a high run time.

~~~~
benchstat v1.txt
name      time/op
PartA-12   193µs ± 2%
PartB-12  0.24ns ± 2%
~~~~

Instead I took an approach that used recursion to explore and update the grid.

~~~~
benchstat v1.txt
name      time/op
PartA-12  68.5µs ± 1%
PartB-12  0.25ns ± 1%
~~~~