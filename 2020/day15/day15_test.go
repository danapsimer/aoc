package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpeakNumbers(t *testing.T) {
	lastNumber := SpeakNumbers([]int{0,3,6})
	assert.Equal(t,436, lastNumber)
}
