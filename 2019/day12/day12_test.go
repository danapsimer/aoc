package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	moons = []*moon{
		{pos: vector{-1, 0, 2}},
		{pos: vector{2, -10, -7}},
		{pos: vector{4, -8, 8}},
		{pos: vector{3, 5, -1}},
	}
	steps = [][]*moon{
		{
			{pos: vector{x: 2, y: -1, z: 1}, vel: vector{x: 3, y: -1, z: -1}},
			{pos: vector{x: 3, y: -7, z: -4}, vel: vector{x: 1, y: 3, z: 3}},
			{pos: vector{x: 1, y: -7, z: 5}, vel: vector{x: -3, y: 1, z: -3}},
			{pos: vector{x: 2, y: 2, z: 0}, vel: vector{x: -1, y: -3, z: 1}},
		},
		{
			{pos: vector{x: 5, y: -3, z: -1}, vel: vector{x: 3, y: -2, z: -2}},
			{pos: vector{x: 1, y: -2, z: 2}, vel: vector{x: -2, y: 5, z: 6}},
			{pos: vector{x: 1, y: -4, z: -1}, vel: vector{x: 0, y: 3, z: -6}},
			{pos: vector{x: 1, y: -4, z: 2}, vel: vector{x: -1, y: -6, z: 2}},
		},
		{
			{pos: vector{x: 5, y: -6, z: -1}, vel: vector{x: 0, y: -3, z: 0}},
			{pos: vector{x: 0, y: 0, z: 6}, vel: vector{x: -1, y: 2, z: 4}},
			{pos: vector{x: 2, y: 1, z: -5}, vel: vector{x: 1, y: 5, z: -4}},
			{pos: vector{x: 1, y: -8, z: 2}, vel: vector{x: 0, y: -4, z: 0}},
		},
		{
			{pos: vector{x: 2, y: -8, z: 0}, vel: vector{x: -3, y: -2, z: 1}},
			{pos: vector{x: 2, y: 1, z: 7}, vel: vector{x: 2, y: 1, z: 1}},
			{pos: vector{x: 2, y: 3, z: -6}, vel: vector{x: 0, y: 2, z: -1}},
			{pos: vector{x: 2, y: -9, z: 1}, vel: vector{x: 1, y: -1, z: -1}},
		},
		{
			{pos: vector{x: -1, y: -9, z: 2}, vel: vector{x: -3, y: -1, z: 2}},
			{pos: vector{x: 4, y: 1, z: 5}, vel: vector{x: 2, y: 0, z: -2}},
			{pos: vector{x: 2, y: 2, z: -4}, vel: vector{x: 0, y: -1, z: 2}},
			{pos: vector{x: 3, y: -7, z: -1}, vel: vector{x: 1, y: 2, z: -2}},
		},
		{
			{pos: vector{x: -1, y: -7, z: 3}, vel: vector{x: 0, y: 2, z: 1}},
			{pos: vector{x: 3, y: 0, z: 0}, vel: vector{x: -1, y: -1, z: -5}},
			{pos: vector{x: 3, y: -2, z: 1}, vel: vector{x: 1, y: -4, z: 5}},
			{pos: vector{x: 3, y: -4, z: -2}, vel: vector{x: 0, y: 3, z: -1}},
		},
		{
			{pos: vector{x: 2, y: -2, z: 1}, vel: vector{x: 3, y: 5, z: -2}},
			{pos: vector{x: 1, y: -4, z: -4}, vel: vector{x: -2, y: -4, z: -4}},
			{pos: vector{x: 3, y: -7, z: 5}, vel: vector{x: 0, y: -5, z: 4}},
			{pos: vector{x: 2, y: 0, z: 0}, vel: vector{x: -1, y: 4, z: 2}},
		},
		{
			{pos: vector{x: 5, y: 2, z: -2}, vel: vector{x: 3, y: 4, z: -3}},
			{pos: vector{x: 2, y: -7, z: -5}, vel: vector{x: 1, y: -3, z: -1}},
			{pos: vector{x: 0, y: -9, z: 6}, vel: vector{x: -3, y: -2, z: 1}},
			{pos: vector{x: 1, y: 1, z: 3}, vel: vector{x: -1, y: 1, z: 3}},
		},
		{
			{pos: vector{x: 5, y: 3, z: -4}, vel: vector{x: 0, y: 1, z: -2}},
			{pos: vector{x: 2, y: -9, z: -3}, vel: vector{x: 0, y: -2, z: 2}},
			{pos: vector{x: 0, y: -8, z: 4}, vel: vector{x: 0, y: 1, z: -2}},
			{pos: vector{x: 1, y: 1, z: 5}, vel: vector{x: 0, y: 0, z: 2}},
		},
		{
			{pos: vector{x: 2, y: 1, z: -3}, vel: vector{x: -3, y: -2, z: 1}},
			{pos: vector{x: 1, y: -8, z: 0}, vel: vector{x: -1, y: 1, z: 3}},
			{pos: vector{x: 3, y: -6, z: 1}, vel: vector{x: 3, y: 2, z: -3}},
			{pos: vector{x: 2, y: 0, z: 4}, vel: vector{x: 1, y: -1, z: -1}},
		},
	}
	energy10 = []int{36, 45, 80, 18, 179}
)

func TestStep(t *testing.T) {
	for i := 0; i < 10; i++ {
		step(moons)
		assert.EqualValues(t, steps[i], moons, "step %d doesn't match", i+1)
	}
	energySum := 0
	for m, moon := range moons {
		energy := moon.energy()
		assert.EqualValues(t, energy10[m], energy)
		energySum += energy
	}
	assert.EqualValues(t, energy10[4], energySum)
}

func TestAxisStateSet(t *testing.T) {
	var s1 axisVec = []int{1, 2, 3}
	var s2 axisVec = []int{1, 2, 4}
	var s3 axisVec = []int{1, 3, 3}
	var s4 axisVec = []int{7, 2, 3}
	set := NewAxisStateSet()
	set.See(1, s1)
	set.See(2, s2)
	step, seen := set.Seen(s1)
	assert.True(t, seen)
	assert.EqualValues(t, step, 1)
	step, seen = set.Seen(s2)
	assert.True(t, seen)
	assert.EqualValues(t, step, 2)
	step, seen = set.Seen(s3)
	assert.False(t, seen)
	step, seen = set.Seen(s4)
	assert.False(t, seen)

	set.See(3, s3)
	set.See(4, s4)
	step, seen = set.Seen(s1)
	assert.True(t, seen)
	assert.EqualValues(t, step, 1)
	step, seen = set.Seen(s2)
	assert.True(t, seen)
	assert.EqualValues(t, step, 2)
	step, seen = set.Seen(s3)
	assert.True(t, seen)
	assert.EqualValues(t, step, 3)
	step, seen = set.Seen(s4)
	assert.True(t, seen)
	assert.EqualValues(t, step, 4)
}

func TestDay12Part2(t *testing.T) {
	file, err := os.Open("./day12.in")
	if assert.NoError(t, err) {
		moons := readInput(file)
		assert.EqualValues(t, 0, Day12Part2(moons))
	}
}
