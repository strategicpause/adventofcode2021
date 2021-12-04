~~~~
benchstat v1.txt
name      time/op
PartA-12  147ns ± 0%
PartB-12  155ns ± 7%
~~~~

There's still an opportunity to improve performance by removing the call to `strings.Split` for the new line characters.