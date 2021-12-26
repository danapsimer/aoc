package polymers

import (
	"fmt"
	"github.com/danapsimer/aoc/2021/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var in = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

var expectedGrowth = []map[rune]int{
	{'N': 2, 'B': 2, 'C': 2, 'H': 1},    //"NCNBCHB"
	{'N': 2, 'B': 6, 'C': 4, 'H': 1},    //"NBCCNBBBCBHCB"
	{'N': 5, 'B': 11, 'C': 5, 'H': 4},   //"NBBBCNCCNBBNBNBBCHBHHBCHB"
	{'N': 11, 'B': 23, 'C': 10, 'H': 5}, //"NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB",
}

func TestPolymer_ApplyRules(t *testing.T) {
	var polymer [1]Polymer
	var rules map[string]rune
	var err error
	polymer[0], rules, err = LoadPolymerAndRules(utils.ReadLinesFromReader(strings.NewReader(in)))
	if assert.NoError(t, err) {
		for step, expectedOut := range expectedGrowth {
			t.Run(fmt.Sprintf("ApplyRules after %d steps", step+1), func(t *testing.T) {
				counts := polymer[0].ApplyRules(rules, step+1)
				assert.Equal(t, expectedOut, counts)
			})
		}
	}
}

func TestPolymer_FindMostAndLeastCommonElements(t *testing.T) {
	polymer, rules, err := LoadPolymerAndRules(utils.ReadLinesFromReader(strings.NewReader(in)))
	if assert.NoError(t, err) {
		counts := polymer.ApplyRules(rules, 10)
		assert.Equal(t, 1749, counts['B'])
		assert.Equal(t, 161, counts['H'])
	}
}
