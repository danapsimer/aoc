package paper

import (
	"github.com/danapsimer/aoc/2021/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var in = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

var out = `#####
#   #
#   #
#   #
#####
`

func TestPaper_String(t *testing.T) {
	lines := utils.ReadLinesFromReader(strings.NewReader(in))
	paper, instructions, err := LoadPaperAndInstructions(lines)
	if assert.NoError(t, err) {
		for _, instruction := range instructions {
			paper.Fold(instruction)
		}
		assert.Equal(t, out, paper.String())
	}
}
