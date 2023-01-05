package main

import (
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ms = sync.Mutex{}
var labels = map[string]string{"label": "label2"}

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
