package main

import (
	"aoc/2019/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGCD(t *testing.T) {
	assert.EqualValues(t, 2, utils.GCD(2, 4))
	assert.EqualValues(t, 2, utils.GCD(2, 6))
	assert.EqualValues(t, 8, utils.GCD(13, 24))
}

func TestReduceSlope(t *testing.T) {
	dx, dy := ReduceSlope(24, 8)
	assert.EqualValues(t, 3, dx)
	assert.EqualValues(t, 1, dy)
	dx, dy = ReduceSlope(24, -8)
	assert.EqualValues(t, 3, dx)
	assert.EqualValues(t, -1, dy)
	dx, dy = ReduceSlope(24, 13)
	assert.EqualValues(t, 24, dx)
	assert.EqualValues(t, 13, dy)
	dx, dy = ReduceSlope(4, 1)
	assert.EqualValues(t, 4, dx)
	assert.EqualValues(t, 1, dy)

}

var (
	testMaps = []string{
		`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`,
		`#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`,
		`.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`,
		`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`,
	}
	bestPositionAndCount = [][]int{
		{5, 8, 33},
		{1, 2, 35},
		{6, 3, 41},
		{11, 13, 210},
	}
)

func TestFindBestLocation(t *testing.T) {
	for i := 0; i < len(testMaps); i++ {
		g := ReadMap(strings.NewReader(testMaps[i]))
		best := g.FindBestLocation()
		assert.EqualValues(t, bestPositionAndCount[i][0], best.x)
		assert.EqualValues(t, bestPositionAndCount[i][1], best.y)
		assert.EqualValues(t, bestPositionAndCount[i][2], len(best.visible))
	}
}

func TestTheta(t *testing.T) {
	assert.EqualValues(t, 0.0, theta(29, 28, 29, 27))
	assert.EqualValues(t, 0.7853981633974483, theta(29, 28, 31, 26))
	assert.EqualValues(t, 1.5707963267948966, theta(29, 28,32, 28))
	assert.EqualValues(t, 2.356194490192345, theta(29, 28,32, 31))
	assert.EqualValues(t, 2.819842099193151, theta(29, 28,30, 31))
	assert.EqualValues(t, 3.6052402625905993, theta(29, 28,28, 30))
	assert.EqualValues(t, 4.124386376837123, theta(29, 28,26, 30))
	assert.EqualValues(t, 4.514993420534809, theta(29, 28,24, 29))
	assert.EqualValues(t, 4.71238898038469, theta(29, 28,23, 28))
	assert.EqualValues(t, 5.117280766669773, theta(29, 28,22, 25))
	assert.EqualValues(t, 5.81953769817878, theta(29, 28,27, 24))

}
