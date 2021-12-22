package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

var testInput = strings.TrimSpace(`16,1,2,0,4,2,7,1,2,14`)

func TestPartA(t *testing.T) {
	assert.Equal(t, 37, PartA(testInput))
}

func TestPartB(t *testing.T) {
	assert.Equal(t, 168, PartB(testInput))
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