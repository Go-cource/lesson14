package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 8.0, Abs(-8))
}
func TestRetTrue(t *testing.T) {
	assert.True(t, true, RetTrue())
}
