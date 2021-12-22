package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

var testInput = strings.TrimSpace(`3,4,3,1,2`)

func TestPartA(t *testing.T) {
	assert.Equal(t, 5934, PartA(testInput))
}

func TestPartB(t *testing.T) {
	assert.Equal(t, 26984457539, PartB(testInput))
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