package main

import (
	"strings"
	"testing"
)

func TestMatrix_MarkWire(t *testing.T) {
	matrix := NewMatrix()

	matrix.MarkWire("w1", 1, 0, 1)
	matrix.MarkWire("w1", 2, -5, 5)
	matrix.MarkWire("w2", 1, -5, 5)
	matrix.MarkWire("w1", 3, 1, 3)
	matrix.MarkWire("w2", 2, 1, 3)
	matrix.MarkWire("w3", 1, 1, 3)
	matrix.MarkWire("w1", 4, 7, 10)
	matrix.MarkWire("w2", 3, 7, 10)
	matrix.MarkWire("w3", 2, 7, 10)
	matrix.MarkWire("w4", 1, 7, 10)
	matrix.MarkWire("w5", 1, 7, 10)
	matrix.MarkWire("w6", 1, 7, 10)

	v := matrix.Get(0, 0)
	if len(v) != 0 {
		t.Errorf("expected 0 but got %v", v)
	}
	v = matrix.Get(0, 1)
	if len(v) != 1 {
		t.Errorf("expected 1 but got %v", v)
	}
	v = matrix.Get(-5, 5)
	if len(v) != 2 {
		t.Errorf("expected 2 but got %v", v)
	}
	v = matrix.Get(1, 3)
	if len(v) != 3 {
		t.Errorf("expected 3 but got %v", v)
	}
	v = matrix.Get(7, 10)
	if len(v) != 6 {
		t.Errorf("expected 6 but got %v", v)
	}
}

var (
	testIn = []string{
		`R8,U5,L5,D3
U7,R6,D4,L4`,
		`R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`,
		`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
	}
	testExpected      = []string{"6", "159", "135"}
	testExpectedPart2 = []string{"30", "610", "410"}
)

func TestDay03(t *testing.T) {
	for i, in := range testIn {
		result := Day03(strings.NewReader(in))
		if result != testExpected[i] {
			t.Errorf("Expected %s but got %s for input\n%s", testExpected[i], result, in)
		}
	}
}
func TestDay03Part2(t *testing.T) {
	for i, in := range testIn {
		result := Day03Part2(strings.NewReader(in))
		if result != testExpectedPart2[i] {
			t.Errorf("Expected %s but got %s for input\n%s", testExpectedPart2[i], result, in)
		}
	}
}
