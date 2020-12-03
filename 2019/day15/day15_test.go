package main

import (
	"aoc/2019/intCode"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDay15Part1(t *testing.T) {
	file, err := os.Open("./day15.in")
	if assert.NoError(t, err) {
		prg := intCode.ReadIntCodeProgram(file)
		assert.EqualValues(t, 238, Day15Part1(prg))
	}
}

func TestDay15Part2(t *testing.T) {
	file, err := os.Open("./day15.in")
	if assert.NoError(t, err) {
		prg := intCode.ReadIntCodeProgram(file)
		assert.EqualValues(t, 238, Day15Part2(prg))
	}
}