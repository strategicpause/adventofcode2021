package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

var testInput = strings.TrimSpace(`2199943210
3987894921
9856789892
8767896789
9899965678`)

func TestPartA(t *testing.T) {
	assert.Equal(t, 15, PartA(testInput))
}

func TestPartB(t *testing.T) {
	assert.Equal(t, 1134, PartB(testInput))
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