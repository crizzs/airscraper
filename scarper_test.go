package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestToCharMethod(t *testing.T) {
	alphabet := toChar(5)
	assert.Equal(t, alphabet, int32(69), "they should be equal")
}