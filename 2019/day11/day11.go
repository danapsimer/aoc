package main

import (
	"aoc/2019/intCode"
	"bytes"
	"log"
	"math"
	"os"
	"sort"
)

type point struct {
	x, y int
}

func (pt *point) Equal(other *point) bool {
	return pt.x == other.x && pt.y == other.y
}

type points []*point

func (pts points) Find(f *point) int {
	idx := sort.Search(len(pts), func(idx int) bool {
		if pts[idx].y >= f.y && pts[idx].x >= f.x {
			return true
		}
		return false
	})
	if pts[idx].Equal(f) {
		return idx
	}
	return -idx
}

func PaintRobot(start bool, prg *intCode.IntCodeProgram) map[point]bool {
	go prg.RunProgram()
	panels := make(map[point]bool)
	pos := point{0, 0}
	panels[pos] = start
	dir := 0
	for !prg.IsFinished() {
		c, ok := panels[pos]
		if ok && c {
			prg.GetInput() <- 1
		} else {
			prg.GetInput() <- 0
		}
		log.Printf("sent status of (%d,%d) = %v",pos.x,pos.y,c)
		paint := <-prg.GetOutput()
		turn := <-prg.GetOutput()
		panels[pos] = paint == 1
		switch turn {
		case 0:
			dir -= 1
			if dir < 0 {
				dir = 3
			}
		case 1:
			dir += 1
			if dir > 3 {
				dir = 0
			}
		}
		switch dir {
		case 0:
			pos.y += 1
		case 1:
			pos.x += 1
		case 2:
			pos.y -= 1
		case 3:
			pos.x -= 1
		}
	}
	return panels
}

func Render(painted map[point]bool) string {
	maxX, maxY := 0, 0
	minX, minY := math.MaxInt64, math.MaxInt64
	for k, _ := range painted {
		if k.x > maxX {
			maxX = k.x
		}
		if k.x < minX {
			minX = k.x
		}
		if k.y > maxY {
			maxY = k.y
		}
		if k.y < minY {
			minY = k.y
		}
	}
	w := bytes.NewBuffer(make([]byte, 0, (maxX-minX)*(maxY-minY)))
	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			pt, ok := painted[point{x, y}]
			if ok && pt {
				w.WriteRune('\u2588')
			} else {
				w.WriteRune(' ')
			}
		}
		w.WriteString("\n")
	}
	return w.String()
}

func main() {
	prg := intCode.ReadIntCodeProgram(os.Stdin)
	painted := PaintRobot(false, prg.Copy())
	log.Printf("part01 = %d", len(painted))
	painted = PaintRobot(true, prg.Copy())
	log.Printf("part02 = \n%s", Render(painted))

}
