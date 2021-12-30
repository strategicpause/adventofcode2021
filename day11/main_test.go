package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

var testInput = strings.TrimSpace(`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`)

func TestPartA(t *testing.T) {
	assert.Equal(t, 1656, PartA(testInput))
}

func TestPartB(t *testing.T) {
	assert.Equal(t, 195, PartB(testInput))
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