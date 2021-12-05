package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

var testInput = strings.TrimSpace(``)

func TestPartA(t *testing.T) {
	assert.Equal(t, 0, PartA(testInput))
}

func TestPartB(t *testing.T) {
	assert.Equal(t, 0, PartB(testInput))
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