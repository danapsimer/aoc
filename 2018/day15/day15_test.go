package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var (
	tests = []string{
		`#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`,
		`#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`,
		`#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`,
		`#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`,
		`#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`,
		`#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`,
		readFileToString("input.txt"),
		`################################
####.#######..G..########.....##
##...........G#..#######.......#
#...#...G.....#######..#......##
########.......######..##.E...##
########......G..####..###....##
#...###.#.....##..G##.....#...##
##....#.G#....####..##........##
##..#....#..#######...........##
#####...G.G..#######...G......##
#########.GG..G####...###......#
#########.G....EG.....###.....##
########......#####...##########
#########....#######..##########
#########G..#########.##########
#########...#########.##########
######...G..#########.##########
#G###......G#########.##########
#.##.....G..#########..#########
#............#######...#########
#...#.........#####....#########
#####.G..................#######
####.....................#######
####.........E..........########
#####..........E....E....#######
####....#.......#...#....#######
####.......##.....E.#E...#######
#####..E...####.......##########
########....###.E..E############
#########.....##################
#############.##################
################################`,
	}
	results      = []int{36334, 39514, 27755, 28944, 18740, 27730, 201856, 215168}
	results2     = []int{29064, 31284, 3478, 6474, 1140, 4988, 0, 52374}
	attackPowers = []int{4, 4, 15, 12, 34, 15, 12, 16}
)

func TestBoard_RunBattle(t *testing.T) {
	for i, test := range tests {
		board := LoadBoard(strings.NewReader(test))
		result := board.RunBattle()
		assert.Equal(t, results[i], result, "test #%d failed, expected %d but found %d", i, results[i], result)
	}
}

func TestBoard_RunBattleForMyInput(t *testing.T) {
	i := 6
	board := LoadBoard(strings.NewReader(tests[i]))
	result := board.RunBattle()
	assert.Equal(t, results[i], result, "test #%d failed, expected %d but found %d", i, results[i], result)
}

func TestBoard_FindAttackPowerToWin(t *testing.T) {
	for i, test := range tests {
		ElfDamage = 4
		result := FindAttachPowerToWin(test)
		assert.Equal(t, results2[i], result, "test #%d failed, expected %d but found %d", i, results2[i], result)
		assert.Equal(t, attackPowers[i], ElfDamage, "test #%d failed, expected %d attack power but found %d", i, attackPowers[i], ElfDamage)
	}
}

func TestBoard_FindAttackPowerToWinForMyInput(t *testing.T) {
	i := 6
	result := FindAttachPowerToWin(tests[i])
	assert.Equal(t, results2[i], result, "test #%d failed, expected %d but found %d", i, results2[i], result)
	assert.Equal(t, attackPowers[i], ElfDamage, "test #%d failed, expected %d attack power but found %d", i, attackPowers[i], ElfDamage)
}
