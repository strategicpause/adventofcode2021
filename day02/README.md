~~~~
benchstat v1.txt
name      time/op
PartA-12  148ns ± 1%
~~~~

There's still an opportunity to improve performance by removing the call to `strings.Split` for the new line characters.