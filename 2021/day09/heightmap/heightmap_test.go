package heightmap

import (
	"fmt"
	"github.com/danapsimer/aoc/2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var scenarios = []struct {
	Map            []string
	ExpectedOutput int
}{
	{
		Map: []string{
			"2199943210",
			"3987894921",
			"9856789892",
			"8767896789",
			"9899965678",
		},
		ExpectedOutput: 1134,
	},
}

func TestHeightMap_FindBasins(t *testing.T) {
	for _, scenario := range scenarios {
		t.Run(fmt.Sprintf("FindBasins %v", scenario.Map), func(t *testing.T) {
			hm := LoadHeightMap(utils.StringArrayToChannel(scenario.Map))
			assert.Equal(t, scenario.ExpectedOutput, hm.FindBasins())
		})
	}
}

var calcBasinSizeScenarios = []struct {
	Map          []string
	X, Y         int
	ExpectedSize int
}{
	{
		Map: []string{
			"2199943210",
			"3987894921",
			"9856789892",
			"8767896789",
			"9899965678",
		},
		X: 1, Y: 0,
		ExpectedSize: 3,
	},
	{
		Map: []string{
			"2199943210",
			"3987894921",
			"9856789892",
			"8767896789",
			"9899965678",
		},
		X: 9, Y: 0,
		ExpectedSize: 9,
	},
	{
		Map: []string{
			"2199943210",
			"3987894921",
			"9856789892",
			"8767896789",
			"9899965678",
		},
		X: 2, Y: 2,
		ExpectedSize: 14,
	},
	{
		Map: []string{
			"2199943210",
			"3987894921",
			"9856789892",
			"8767896789",
			"9899965678",
		},
		X: 6, Y: 4,
		ExpectedSize: 9,
	},
}

func TestHeightMap_CalcBasinSize(t *testing.T) {
	for _, scenario := range calcBasinSizeScenarios {
		t.Run(fmt.Sprintf("CalcBasinSize %v %d,%d", scenario.Map, scenario.X, scenario.Y), func(t *testing.T) {
			hm := LoadHeightMap(utils.StringArrayToChannel(scenario.Map))
			hm.CalcBasinSize(scenario.X, scenario.Y)
		})
	}
}
