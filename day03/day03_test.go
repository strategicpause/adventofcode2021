package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

var testInput = strings.TrimSpace(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`)

func TestPartA(t *testing.T) {
	assert.Equal(t, 198, PartA(testInput))
}

func TestPartB(t *testing.T) {
	assert.Equal(t, 230, PartB(testInput))
}

func BenchmarkPartA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartA(testInput)
	}
}

func BenchmarkPartB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartB(testInput)
	}
}