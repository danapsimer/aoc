package main

import (
	"aoc/2019/utils"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type Axis int

const (
	X Axis = iota
	Y
	Z
)

func (v *vector) get(axis Axis) int {
	if axis < X || axis > Z {
		panic(fmt.Errorf("unexpected axis value: %d", axis))
	}
	return v[axis]
}

func (v *vector) put(axis Axis, value int) *vector {
	if axis < X || axis > Z {
		panic(fmt.Errorf("unexpected axis value: %d", axis))
	}
	v[axis] = value
	return v
}

func iabs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func calcGravityComponent(v1, v2 int) int {
	if v1 > v2 {
		return -1
	} else if v1 < v2 {
		return 1
	} else {
		return 0
	}
}

type vector [3]int

func (v vector) Equals(v2 vector) bool {
	for idx, value := range v {
		if v2[idx] != value {
			return false
		}
	}
	return true
}

func (v vector) Less(v2 vector) bool {
	for idx, value := range v {
		if v2[idx] < value {
			return false
		}
	}
	return true
}

func (v vector) energy() int {
	sum := 0
	for _, value := range v {
		sum += iabs(value)
	}
	return sum
}

type moon struct {
	pos vector
	vel vector
}

func (m *moon) Equals(m2 *moon) bool {
	return m.pos.Equals(m2.pos) && m.vel.Equals(m2.vel)
}

func (m *moon) Less(m2 *moon) bool {
	if m.pos.Less(m2.pos) {
		return true
	} else if m.pos.Equals(m2.pos) && m.vel.Less(m2.vel) {
		return true
	}
	return false
}

func (m *moon) energy() int {
	return m.pos.energy() * m.vel.energy()
}

func (m *moon) AxisEqual(m2 *moon, axis Axis) bool {
	return m.pos.get(axis) == m2.pos.get(axis) && m.vel.get(axis) == m2.pos.get(axis)
}

type Moons []*moon

func (moons Moons) Less(mm2 Moons) bool {
	for i, m := range moons {
		if !m.Less(mm2[i]) {
			return false
		}
	}
	return true
}

func (moons Moons) Equals(mm2 Moons) bool {
	for i, m := range moons {
		if !m.Equals(mm2[i]) {
			return false
		}
	}
	return true
}

func (moons Moons) Copy() Moons {
	cpy := make([]*moon, len(moons))
	for m, mn := range moons {
		cpyMoon := *mn
		cpy[m] = &cpyMoon
	}
	return cpy
}

func (moons Moons) energy() int {
	total := 0
	for _, m := range moons {
		total += m.energy()
	}
	return total
}

func step(moons []*moon) {
	for i, m1 := range moons {
		for j, m2 := range moons {
			if i != j {
				for a := X; a <= Z; a++ {
					m1.vel.put(a, m1.vel.get(a)+calcGravityComponent(m1.pos.get(a), m2.pos.get(a)))
				}
			}
		}
	}
	for _, m := range moons {
		for a := X; a <= Z; a++ {
			m.pos.put(a, m.pos.get(a)+m.vel.get(a))
		}
	}
}

func stepForAxis(axis Axis, moons []*moon) {
	for i, m1 := range moons {
		for j, m2 := range moons {
			if i != j {
				m1.vel.put(axis, m1.vel.get(axis)+calcGravityComponent(m1.pos.get(axis), m2.pos.get(axis)))
			}
		}
	}
	for _, m := range moons {
		m.pos.put(axis, m.pos.get(axis)+m.vel.get(axis))
	}
}

func Day12(moons Moons) int {
	for i := 0; i < 1000; i++ {
		step(moons)
	}
	return moons.energy()
}

type axisVec []int

type axisStateSet struct {
	step     int
	depth    int
	children map[int]*axisStateSet
}

func NewAxisStateSet() *axisStateSet {
	return &axisStateSet{0, 0, make(map[int]*axisStateSet)}
}

func (set *axisStateSet) Seen(vec axisVec) (int, bool) {
	node := set
	for _, value := range vec {
		var ok bool
		node, ok = node.children[value]
		if !ok {
			return -1, false
		}
	}
	return node.step, true
}

func (set *axisStateSet) See(step int, vec axisVec) {
	node := set
	for depth, value := range vec {
		child, ok := node.children[value]
		if !ok {
			child = &axisStateSet{step, depth + 1, make(map[int]*axisStateSet)}
			node.children[value] = child
		}
		node = child
	}
}

func Day12Part2(moons Moons) int {
	curr := moons.Copy()
	counters := make([]int, 4)
	done := make([]bool, 4)
	seen := make([]*axisStateSet, 4)
	seen[X] = NewAxisStateSet()
	seen[Y] = NewAxisStateSet()
	seen[Z] = NewAxisStateSet()
	initial := []axisVec{
		buildAxisStateVec(curr, X),
		buildAxisStateVec(curr, Y),
		buildAxisStateVec(curr, Z),
	}
	seen[X].See(0, initial[X])
	seen[Y].See(0, initial[Y])
	seen[Z].See(0, initial[Z])
	compareToFirst := true
	wg := &sync.WaitGroup{}
	minFirstStep := -1
	for axis := X; axis <= Z; axis++ {
		wg.Add(1)
		go func(axis Axis) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("ERROR: %+v", r)
				}
				log.Printf("Done calculating repeat of %d axis: %d", axis, counters[axis])
				wg.Done()
			}()
			steps := 0
			for !done[axis] {
				stepForAxis(axis, curr)
				steps += 1
				axisState := buildAxisStateVec(curr, axis)
				if compareToFirst {
					equal := true
					for ias, as := range axisState {
						if initial[axis][ias] != as {
							equal = false
							break
						}
					}
					if equal {
						done[axis] = true
						counters[axis] = steps
					}
				} else {
					if lastSteps, ok := seen[axis].Seen(axisState); ok {
						done[axis] = true
						counters[axis] = steps - lastSteps
						if minFirstStep < 0 {
							minFirstStep = lastSteps
						}
					} else {
						seen[axis].See(steps, axisState)
					}
				}
			}
		}(axis)
	}
	wg.Wait()
	return utils.LCM(counters[X], counters[Y], counters[Z])
}

func buildAxisStateVec(curr Moons, axis Axis) axisVec {
	var axisState axisVec = make([]int, len(curr)*2)
	for i, m := range curr {
		axisState[i] = m.pos.get(axis)
		axisState[len(curr)+i] = m.vel.get(axis)
	}
	return axisState
}

func readInput(reader io.Reader) Moons {
	moons := make([]*moon, 0, 4)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		var x, y, z int
		n, err := fmt.Sscanf(scanner.Text(), "<x=%d, y=%d, z=%d>", &x, &y, &z)
		if err != nil {
			panic(err)
		}
		if n != 3 {
			panic(fmt.Errorf("incomplete line found: %d", n))
		}
		moons = append(moons, &moon{pos: vector{x, y, z}})
	}
	return moons
}

func main() {
	moons := readInput(os.Stdin)
	log.Printf("part1 = %d", Day12(moons.Copy()))
	log.Printf("part2 = %d", Day12Part2(moons))
}
