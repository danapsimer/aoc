package deck

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var testInput = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

func TestReadGame(t *testing.T) {

	game := ReadGame(strings.NewReader(testInput))
	assert.Equal(t, 2, len(game))
	assert.Equal(t, 5, len(*game[0]))
	assert.Equal(t, 5, len(*game[1]))
}
