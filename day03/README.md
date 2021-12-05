~~~~
$ benchstat v1.txt
name      time/op
PartA-12  102ns ± 1%
PartB-12  339ns ± 1%

name      alloc/op
PartA-12  0.00B
PartB-12   192B ± 0%

name      allocs/op
PartA-12   0.00
PartB-12   1.00 ± 0%
~~~~

For PartB I took more of a functional approach to avoid code duplication. Originally I was using the `append` function to build my slices in the `FilterRatings`function. Instead I opted to create a fixed size array to act as a buffer for storing the new ratings array for both Oxygen and Carbon Dioxide; this was after noting most of my time was being spent allocating memory in the loops.

Again, if I want to optimize this further I can look at replacing strings.Split with my own custom code.