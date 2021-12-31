package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

var testInput = strings.TrimSpace(`6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`)

func TestPartA(t *testing.T) {
	assert.Equal(t, 17, PartA(testInput))
}

func TestPartB(t *testing.T) {
	assert.Equal(t, 1, PartB(testInput))
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