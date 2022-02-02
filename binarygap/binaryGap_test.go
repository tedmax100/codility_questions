package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldBeZero(t *testing.T) {
	assert.Equal(t, 0, Solution(32))
}

func TestShouldBeSuccess(t *testing.T) {
	assert.Equal(t, 5, Solution(1041))
}
