package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

var (
	inputs = []string{
		`#########
#b.A.@.a#
#########`,
		`########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`,
		`########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`,
		`#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`,
		`########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################`,
	}
	outputs      = []int{8, 86, 132, 136, 81}
	outputsOrder = []string{
		"ab",
		"abcdef",
		"bacdfeg",
		"afbjgnhdloepcikm",
		"acfidgbeh",
	}
)

func TestDay18Part1(t *testing.T) {
	for test := range inputs {
		t.Run(strconv.Itoa(test), func(t *testing.T) {
			grid := ReadGrid(strings.NewReader(inputs[test]))
			steps, order := Day18Part1(grid)
			assert.Equal(t, outputs[test], steps)
			if outputsOrder[test] != order {
				t.Logf("order was different than suggested: %s, expected: %s", order, outputsOrder[test])
			}
		})
	}
}
