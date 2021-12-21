package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransform(t *testing.T) {
	assert.Equal(t, 5764801, Transform(7, 8))
	assert.Equal(t, 17807724, Transform(7, 11))
}

func TestFindLoopSize(t *testing.T) {
	assert.Equal(t, 8, FindLoopSize(5764801))
	assert.Equal(t, 11, FindLoopSize(17807724))
}
