~~~~
name      time/op
PartA-12   423ns ± 9%
PartB-12  1.33µs ± 7%
~~~~
For PartB the summation funciton looked like

~~~~
sum := 0
for i := 0; i <= n; i++ {
	sum += i
}
return sum
~~~~
However, I realized this could be replaced with something much more efficient:
~~~~
return int(float64(n) / 2.0 * float64(1 + n))
~~~~

This resulted in the following speed up
~~~~
name      time/op
PartA-12  370ns ± 1%
PartB-12  499ns ± 0%
~~~~