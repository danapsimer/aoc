package octo

import (
	"fmt"
	"github.com/danapsimer/aoc/2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var octoMap = []string{
	"5483143223",
	"2745854711",
	"5264556173",
	"6141336146",
	"6357385478",
	"4167524645",
	"2176841721",
	"6882881134",
	"4846848554",
	"5283751526",
}

var steps = map[int][]string{
	1: {
		"6594254334",
		"3856965822",
		"6375667284",
		"7252447257",
		"7468496589",
		"5278635756",
		"3287952832",
		"7993992245",
		"5957959665",
		"6394862637",
	},
	2: {
		"8807476555",
		"5089087054",
		"8597889608",
		"8485769600",
		"8700908800",
		"6600088989",
		"6800005943",
		"0000007456",
		"9000000876",
		"8700006848",
	},
	3: {
		"0050900866",
		"8500800575",
		"9900000039",
		"9700000041",
		"9935080063",
		"7712300000",
		"7911250009",
		"2211130000",
		"0421125000",
		"0021119000",
	},
}

func TestNRounds(t *testing.T) {
	for step, result := range steps {
		t.Run(fmt.Sprintf("%d Rounds", step), func(t *testing.T) {
			octopi := LoadOctopi(utils.StringArrayToChannel(octoMap))
			NRounds(step, octopi)
			assert.Equal(t, LoadOctopi(utils.StringArrayToChannel(result)), octopi)
		})
	}
}
