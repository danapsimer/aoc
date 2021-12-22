package diagnostic

import (
	"fmt"
	"github.com/danapsimer/aoc/2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var scenarios = []struct {
	input                                []string
	expectedOxyRating, expectedCo2Rating uint
}{
	{
		input: []string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		},
		expectedOxyRating: 23,
		expectedCo2Rating: 10,
	},
}

func TestCalculateRatings(t *testing.T) {
	for _, scenario := range scenarios {
		t.Run(fmt.Sprintf("input = %v", scenario.input), func(t *testing.T) {
			in, width, err := ReadInputFromChannel(utils.StringArrayToChannel(scenario.input))
			if assert.NoError(t, err) {
				o, c := CalculateRatings(in, width)
				assert.Equal(t, scenario.expectedOxyRating, o)
				assert.Equal(t, scenario.expectedCo2Rating, c)
			}
		})
	}
}
