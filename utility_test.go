package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestAdd(t *testing.T) {
	output := Add(3, 5)
	assert.Equal(t, output, 8)
}
