~~~~
$ benchstat v1.txt
name      time/op
PartA-12  103ns ± 4%
PartB-12  879ns ± 0%
~~~~

For PartB I took more of a functional approach to avoid code duplication. Originally I was using the `append` function to build my slices in the `FilterRatings`function. Instead I opted to create a fixed size array to get under 1uS as a lot of time was being spent allocating memory.