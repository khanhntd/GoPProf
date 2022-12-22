package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSolve(t *testing.T) {
	c, _ := os.ReadFile("1.txt")
	res := solve(c)
	expected := 5
	assert.Equal(t, res, expected)
}
func BenchmarkSolve(b *testing.B) {
	c, err := os.ReadFile("1.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		solve(c)
	}
}
