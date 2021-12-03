package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

var testInput = strings.TrimSpace(`
forward 5
down 5
forward 8
up 3
down 8
forward 2`)

func TestPartA(t *testing.T) {
	assert.Equal(t, 150, PartA(testInput))
}

func BenchmarkPartA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartA(testInput)
	}
}