package main

import (
	"testing"
	"strings"
)

var testInput = strings.TrimSpace(`
199
200
208
210
200
207
240
269
260
263
`)

func BenchmarkPartA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartA(testInput)
	}
}