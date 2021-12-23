package caves

import (
	"github.com/danapsimer/aoc/2021/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var testCaves = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

func TestLoadCaves(t *testing.T) {
	cs, err := LoadCaves(utils.ReadLinesFromReader(strings.NewReader(testCaves)))
	if assert.NoError(t, err) {
		assert.Equal(t, 10, len(cs))
	}
}

func TestCaves_FindPaths(t *testing.T) {
	cs, err := LoadCaves(utils.ReadLinesFromReader(strings.NewReader(testCaves)))
	if assert.NoError(t, err) {
		assert.Equal(t, 10, len(cs))
		assert.Equal(t, 226, len(cs.FindPaths("")))
	}
}

func TestCaves_FindPathWith1DuplicateSmallCave(t *testing.T) {
	cs, err := LoadCaves(utils.ReadLinesFromReader(strings.NewReader(testCaves)))
	if assert.NoError(t, err) {
		assert.Equal(t, 10, len(cs))
		assert.Equal(t, 3509, len(cs.FindPathWith1DuplicateSmallCave()))
	}
}
