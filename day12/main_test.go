package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

var testInput = strings.TrimSpace(`start-A
start-b
A-c
A-b
b-d
A-end
b-end`)

func TestPartA(t *testing.T) {
	assert.Equal(t, 10, PartA(testInput))
}

func TestPartB(t *testing.T) {
	assert.Equal(t, 36, PartB(testInput))
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