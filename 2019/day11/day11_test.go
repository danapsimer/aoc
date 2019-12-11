package main

import (
	"aoc/2019/intCode"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPaintRobot(t *testing.T) {
	file, err := os.Open("./day11.in")
	if assert.NoError(t, err) {
		prg := intCode.ReadIntCodeProgram(file)
		panels := PaintRobot(false, prg.Copy())
		assert.EqualValues(t, 2238, len(panels))
		panels = PaintRobot(true, prg.Copy())
		assert.EqualValues(t, ``, Render(panels))
	}
}
