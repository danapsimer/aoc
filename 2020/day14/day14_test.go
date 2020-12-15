package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var program = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

func TestReadAndRunProgram(t *testing.T) {
	memory := RunProgramPart1(ReadProgram(strings.NewReader(program)))
	assert.Equal(t, uint64(64), memory[8])
	assert.Equal(t, uint64(101), memory[7])
	assert.Equal(t, 2, len(memory))
}

var program2 = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

func TestRunProgramPart2(t *testing.T) {
	memory := RunProgramPart2(ReadProgram(strings.NewReader(program2)))
	assert.Equal(t, uint64(100), memory[58])
	assert.Equal(t, uint64(100), memory[59])
	assert.Equal(t, uint64(1), memory[16])
	assert.Equal(t, uint64(1), memory[17])
	assert.Equal(t, uint64(1), memory[18])
	assert.Equal(t, uint64(1), memory[19])
	assert.Equal(t, uint64(1), memory[24])
	assert.Equal(t, uint64(1), memory[25])
	assert.Equal(t, uint64(1), memory[26])
	assert.Equal(t, uint64(1), memory[27])
	assert.Equal(t, 10, len(memory))

}
