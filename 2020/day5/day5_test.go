package main

import (
	"strconv"
	"testing"
	"github.com/stretchr/testify/assert"
)

var parseBoardingPassTests = []struct {
	BPString string
	Row, Col uint8
	Id boardingPass
} {
	{ "BFFFBBFRRR", 70, 7, 567},
	{ "FFFBBBFRRR", 14, 7, 119},
	{ "BBFFBBFRLL", 102, 4, 820},
}

func TestParseBoardingPass(t *testing.T) {
	for idx, test := range parseBoardingPassTests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			bp := ParseBoardingPass(test.BPString)
			assert.Equal(t, test.Id, bp)
			assert.Equal(t, test.Row, bp.row())
			assert.Equal(t, test.Col, bp.col())
		})
	}
}
