package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var testImg = []uint16{
	0b0011010010,
	0b1100100000,
	0b1000110010,
	0b1111010001,
	0b1101101110,
	0b1100010111,
	0b0101010011,
	0b0010000100,
	0b1110001010,
	0b0011100111,
}

var testImgRL1 = []uint16{
	0b0100111110,
	0b0101111010,
	0b1110001001,
	0b1001011001,
	0b1000010110,
	0b0001101101,
	0b0100010000,
	0b1010110000,
	0b1101110101,
	0b1001101000,
}

func TestNewTile(t *testing.T) {
	id := 2311
	tile := NewTile(id, testImg)
	assert.Equal(t, id, tile.Id)
	assert.EqualValues(t, testImg, tile.Img)
	assert.Equal(t, uint16(0b0011010010), tile.Boarder[North])
	assert.Equal(t, uint16(0b0001011001), tile.Boarder[East])
	assert.Equal(t, uint16(0b0011100111), tile.Boarder[South])
	assert.Equal(t, uint16(0b0111110010), tile.Boarder[West])
}

func TestTile_Rotate(t *testing.T) {
	tile := NewTile(0, testImg)
	tile.Rotate()
	assert.EqualValues(t, testImgRL1, tile.Img)
}

var testInput = `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...

`

func TestReadTiles(t *testing.T) {
	tiles := ReadTiles(strings.NewReader(testInput))
	assert.Equal(t, 9, len(tiles))
}

func TestFindCorners(t *testing.T) {
	tiles := ReadTiles(strings.NewReader(testInput))
	FindEdges(tiles)
	corners := FindCorners(tiles)
	assert.Equal(t, 4, len(corners))
}

func TestOrientTiles(t *testing.T) {
	tiles := ReadTiles(strings.NewReader(testInput))
	FindEdges(tiles)
	ulc := OrientTiles(tiles)
	if assert.NotNil(t, ulc) {

	}
}
