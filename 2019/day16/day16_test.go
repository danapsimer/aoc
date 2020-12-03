package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	inputs = [][]byte{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{8, 0, 8, 7, 1, 2, 2, 4, 5, 8, 5, 9, 1, 4, 5, 4, 6, 6, 1, 9, 0, 8, 3, 2, 1, 8, 6, 4, 5, 5, 9, 5},
		{1, 9, 6, 1, 7, 8, 0, 4, 2, 0, 7, 2, 0, 2, 2, 0, 9, 1, 4, 4, 9, 1, 6, 0, 4, 4, 1, 8, 9, 9, 1, 7},
		{6, 9, 3, 1, 7, 1, 6, 3, 4, 9, 2, 9, 4, 8, 6, 0, 6, 3, 3, 5, 9, 9, 5, 9, 2, 4, 3, 1, 9, 8, 7, 3},
	}
	iterations = []int{
		4, 100, 100, 100,
	}
	outputs = []string{
		"01029498",
		"24176176",
		"73745418",
		"52432133",
	}

	part2Inputs = [][]byte{
		{0, 3, 0, 3, 6, 7, 3, 2, 5, 7, 7, 2, 1, 2, 9, 4, 4, 0, 6, 3, 4, 9, 1, 5, 6, 5, 4, 7, 4, 6, 6, 4},
		{0, 2, 9, 3, 5, 1, 0, 9, 6, 9, 9, 9, 4, 0, 8, 0, 7, 4, 0, 7, 5, 8, 5, 4, 4, 7, 0, 3, 4, 3, 2, 3},
		{0, 3, 0, 8, 1, 7, 7, 0, 8, 8, 4, 9, 2, 1, 9, 5, 9, 7, 3, 1, 1, 6, 5, 4, 4, 6, 8, 5, 0, 5, 1, 7},
	}
	part2Outputs = []string{
		"84462026",
		"78725270",
		"53553731",
	}
)

func TestDay16Part1(t *testing.T) {
	for testNum := range inputs {
		t.Run(fmt.Sprintf("%d", testNum), func(t *testing.T) {
			assert.EqualValues(t, outputs[testNum], Day16Part1(inputs[testNum], iterations[testNum], 0))
		})
	}
}

func TestDay16Part2(t *testing.T) {
	for testNum := range part2Inputs {
		t.Run(fmt.Sprintf("%d", testNum), func(t *testing.T) {
			assert.EqualValues(t, part2Outputs[testNum], Day16Part2(part2Inputs[testNum], 100))
		})
	}
}
