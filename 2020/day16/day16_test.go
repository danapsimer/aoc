package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var testNotes = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

func TestReadNotes(t *testing.T) {
	notes := ReadNotes(strings.NewReader(testNotes))
	if assert.NotNil(t, notes) {
		if assert.NotNil(t, notes.Fields) {
			assert.Equal(t, 3, len(notes.Fields))
			class, ok := notes.Fields["class"]
			assert.True(t, ok, "class field not present")
			assert.Equal(t, "class", class.Name)
			assert.Equal(t, 1, class.Ranges[0].Start)
			assert.Equal(t, 3, class.Ranges[0].End)
			assert.Equal(t, 5, class.Ranges[1].Start)
			assert.Equal(t, 7, class.Ranges[1].End)

			seat, ok := notes.Fields["seat"]
			assert.True(t, ok, "seat field not present")
			assert.Equal(t, "seat", seat.Name)
			assert.Equal(t, 13, seat.Ranges[0].Start)
			assert.Equal(t, 40, seat.Ranges[0].End)
			assert.Equal(t, 45, seat.Ranges[1].Start)
			assert.Equal(t, 50, seat.Ranges[1].End)
		}
		if assert.NotNil(t, notes.MyTicket) {
			assert.Equal(t, 3, len(notes.MyTicket))
			assert.EqualValues(t, Ticket{7, 1, 14}, notes.MyTicket)
		}
		if assert.NotNil(t, notes.NearbyTickets) {
			assert.Equal(t, 4, len(notes.NearbyTickets))
			assert.EqualValues(t, Ticket{7, 3, 47}, notes.NearbyTickets[0])
		}
	}
}

func TestSumInvalidValues(t *testing.T) {
	notes := ReadNotes(strings.NewReader(testNotes))
	sum, _ := SumInvalidValues(notes)
	assert.Equal(t, 71, sum)
}

var inputPart2 = `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`

func TestCalculateFieldOrder(t *testing.T) {
	notes := ReadNotes(strings.NewReader(inputPart2))
	_, validTickets := SumInvalidValues(notes)
	fieldOrder := CalculateFieldOrder(notes, validTickets)
	if assert.NotNil(t, fieldOrder) {
		assert.Equal(t, []string{"row", "class", "seat"}, fieldOrder)
	}
}

func TestCalculateFieldOrder2(t *testing.T) {
	f, err := os.Open("./day16.in")
	if assert.NoError(t, err) {
		notes := ReadNotes(f)
		_, validTickets := SumInvalidValues(notes)
		fieldOrder := CalculateFieldOrder(notes, validTickets)
		if assert.NotNil(t, fieldOrder) {
			assert.Equal(t, []string{"row", "class", "seat"}, fieldOrder)
		}
	}
}