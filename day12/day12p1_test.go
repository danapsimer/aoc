package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBits_Copy(t *testing.T) {
	b1 := NewBitsFromString("#.#.#")
	b2 := b1.Copy()
	assert.Equal(t,b1,b2)
}

func TestBits_Set(t *testing.T) {
	b1 := NewBitsFromString(".")
	b1.Set(-2,true)
	t.Logf("%s", b1.String())
	b1.Set(7, true)
	assert.Equal(t, 10, b1.Len())
	assert.Equal(t, 7, b1.Last())
	assert.Equal(t, -2, b1.First())
}

func TestBits_ContainsAt(t *testing.T) {
	b1 := NewBitsFromString(".....#.#.#....")
	b2 := NewBitsFromString("#.#")

	assert.True(t, b1.ContainsAt(5, b2))
	assert.True(t, b1.ContainsAt(7, b2))
	assert.False(t, b1.ContainsAt(1, b2))
	assert.False(t, b1.ContainsAt(4, b2))
}
