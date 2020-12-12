package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var in = `F10
N3
F7
R90
F11`

func TestPart1(t *testing.T) {
	instructions := ReadInstructions(strings.NewReader(in))
	assert.Equal(t, 25, Part1(instructions))
}

func TestPart2(t *testing.T) {
	instructions := ReadInstructions(strings.NewReader(in))
	assert.Equal(t, 286, Part2(instructions))
}

var rotations = []struct {
	x, y, theta, expectedX, expectedY int
}{
	{1, 0, 90, 0, 1},
	{0, 1, 90, -1, 0},
	{-1, 0, 90, 0, -1},
	{0, -1, 90, 1, 0},
	{1, 0, 180, -1, 0},
	{0, 1, 180, 0, -1},
	{-1, 0, 180, 1, 0},
	{0, -1, 180, 0, 1},
	{1, 0, 270, 0, -1},
	{0, 1, 270, 1, 0},
	{-1, 0, 270, 0, 1},
	{0, -1, 270, -1, 0},
}

func TestRotate(t *testing.T) {
	for _, test := range rotations {
		t.Run(fmt.Sprintf("(%d,%d)@%d", test.x, test.y, test.theta), func(t *testing.T) {
			actualX, actualY := Rotate(test.x,test.y,test.theta)
			assert.Equal(t, test.expectedX, actualX)
			assert.Equal(t, test.expectedY, actualY)
		})
	}
}
