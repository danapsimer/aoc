package main

import (
	"aoc/2019/intCode"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var in = []string{
	"..#..........",
	"..#..........",
	"#######...###",
	"#.#...#...#.#",
	"#############",
	"..#...#...#..",
	"..#####...^..",
}
var expected = 76

func TestDay17Part1(t *testing.T) {
	assert.Equal(t, expected, Day17Par1(in))
}

func TestDay17Part2(t *testing.T) {
	file, err := os.Open("./day17.in")
	if assert.NoError(t, err) {
		prg := intCode.ReadIntCodeProgram(file)
		assert.EqualValues(t, 0, Day17Part2(prg))
	}
}
