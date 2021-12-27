This was actually a pretty fun one. Rather than creating some n x m grid to represent the different heights, I opted to use a map to manage these points. The key is a one-to-one mapping from the (x, y) 2-D coordinate system to a 1-D value (x * 1000 + y) which allows me to quickly look up values in the map. To find neighbors I generate the keys for the corresponding (x, y) values and see if they're present in the map.

PartA was fairly trivial (after the data structure has been generated) - I simply iterate over the entries in the map and see if the given value is less than all of its neighbors. PartB was a bit more interesting as I iterate over the values and find an entry that has not yet been explored and is not equal to 9 (the highest value). Once I find that value I use a BFS to explore all of the neighbors in the basin and return the size (Excluding the entries with the height value of 9).

~~~~
benchstat v1.txt
name      time/op
PartA-12  9.27µs ± 1%
PartB-12  10.7µs ± 0%
~~~~