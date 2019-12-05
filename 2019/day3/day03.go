package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type direction int

const (
	U direction = iota
	R
	D
	L
)

func DirectionFromString(s string) direction {
	switch s {
	case "U":
		return U
	case "R":
		return R
	case "D":
		return D
	case "L":
		return L
	default:
		panic(fmt.Errorf("unknown distance %s", s))
	}
}

type wire struct {
	d direction
	n int
}

type matrixNode struct {
	name  string
	steps int
}

type matrix struct {
	rows   []int
	cols   []int
	values [][]*matrixNode
}

func NewMatrix() *matrix {
	return &matrix{
		make([]int, 0, 10000),
		make([]int, 0, 10000),
		make([][]*matrixNode, 0, 10000),
	}
}

func (m *matrix) MarkWire(name string, steps, x, y int) {
	r := m.get(x, y)
	if r < len(m.rows) && m.rows[r] == y && m.cols[r] == x {
		c := sort.Search(len(m.values[r]), func(i int) bool { return m.values[r][i].name == name })
		if c == len(m.values[r]) || m.values[r][c].name != name {
			m.values[r] = append(m.values[r], nil)
			copy(m.values[r][c+1:], m.values[r][c:])
			m.values[r][c] = &matrixNode{name, steps}
		}
	} else {
		m.rows = append(m.rows, 0)
		copy(m.rows[r+1:], m.rows[r:])
		m.rows[r] = y

		m.cols = append(m.cols, 0)
		copy(m.cols[r+1:], m.cols[r:])
		m.cols[r] = x

		m.values = append(m.values, nil)
		copy(m.values[r+1:], m.values[r:])
		m.values[r] = []*matrixNode{{name, steps}}
	}
}

func (m *matrix) Get(x, y int) []*matrixNode {
	r := m.get(x, y)
	if r < len(m.rows) && m.rows[r] == y && m.cols[r] == x {
		return m.values[r]
	}
	return nil
}

func (m *matrix) get(x, y int) int {
	r := sort.SearchInts(m.rows, y)
	if r < len(m.rows) && m.rows[r] == y {
		for ; r < len(m.rows) && m.rows[r] == y && m.cols[r] < x; r++ {
			// do nothing.
		}
		if r < len(m.cols) && m.cols[r] == x {
			return r
		}
	}
	return r
}

func (m *matrix) AddWireSegment(name string, stepsin, sx, sy int, w *wire) (steps, x, y int) {
	steps = stepsin
	switch w.d {
	case U:
		x = sx
		for y = sy; y < sy+w.n; y += 1 {
			m.MarkWire(name, steps, x, y)
			steps += 1
		}
	case D:
		x = sx
		for y = sy; y > sy-w.n; y -= 1 {
			m.MarkWire(name, steps, x, y)
			steps += 1
		}
	case R:
		y = sy
		for x = sx; x < sx+w.n; x += 1 {
			m.MarkWire(name, steps, x, y)
			steps += 1
		}
	case L:
		y = sy
		for x = sx; x > sx-w.n; x -= 1 {
			m.MarkWire(name, steps, x, y)
			steps += 1
		}
	default:
		panic(fmt.Sprintf("unknown direction: %d", w.d))
	}
	return
}

func (m *matrix) AddWire(name string, ww []*wire) {
	steps, x, y := 0, 0, 0
	for _, w := range ww {
		steps, x, y = m.AddWireSegment(name, steps, x, y, w)
	}
}

func iabs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (m *matrix) FindClosestDistance() int {
	minD := math.MaxInt64
	for r := 0; r < len(m.values); r++ {
		if m.rows[r] != 0 && m.cols[r] != 0 && len(m.values[r]) > 1 {
			d := iabs(m.rows[r]) + iabs(m.cols[r])
			if d < minD {
				minD = d
			}
		}
	}
	return minD
}

func (m *matrix) FindShortestSteps() int {
	minSteps := math.MaxInt64
	for r := 0; r < len(m.values); r++ {
		if m.rows[r] != 0 && m.cols[r] != 0 && len(m.values[r]) > 1 {
			sum := 0
			for _, n := range m.values[r] {
				sum += n.steps
			}
			if sum < minSteps {
				minSteps = sum
			}
		}
	}
	return minSteps
}

func parseWire(line string) []*wire {
	ww := make([]*wire, 0, 1000)
	segments := strings.Split(line, ",")
	for _, segment := range segments {
		n, err := strconv.Atoi(segment[1:])
		if err != nil {
			panic(err)
		}
		ww = append(ww, &wire{
			DirectionFromString(segment[:1]),
			n,
		})
	}
	return ww
}

func Day03(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	matrix := NewMatrix()
	w := 0
	for scanner.Scan() {
		matrix.AddWire(fmt.Sprintf("wire%d", w), parseWire(scanner.Text()))
		w += 1
	}
	return fmt.Sprintf("%d", matrix.FindClosestDistance())
}

func Day03Part2(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	matrix := NewMatrix()
	w := 0
	for scanner.Scan() {
		start := time.Now()
		wire := parseWire(scanner.Text())
		parseWireTime := time.Now().Sub(start)
		wireName := fmt.Sprintf("wire%d", w)
		log.Printf("parsting %s took %s", wireName, parseWireTime.String())
		start = time.Now()
		matrix.AddWire(wireName, wire)
		addWireTime := time.Now().Sub(start)
		log.Printf("adding Wire %s took %s", wireName, addWireTime.String())
		w += 1
	}
	start := time.Now()
	shortestSteps := matrix.FindShortestSteps()
	findShortestStepsTime := time.Now().Sub(start)
	log.Printf("Finding the shortest steps took %s", findShortestStepsTime.String())
	return fmt.Sprintf("%d", shortestSteps)
}

func main() {
	fmt.Println(Day03Part2(os.Stdin))
}
