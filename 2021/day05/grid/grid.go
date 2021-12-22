package grid

import (
	"fmt"
	"github.com/danapsimer/aoc/2021/utils"
	"log"
)

type Grid [][]int
type Point struct {
	x, y int
}
type Line struct {
	p1, p2 Point
}

func NewGrid(h, w int) Grid {
	g := make([][]int, h)
	for y, _ := range g {
		g[y] = make([]int, w)
	}
	return g
}

func (g Grid) drawLine(l Line) {
	if l.p1.x != l.p2.x && l.p1.y != l.p2.y {
		dx := 1
		if l.p1.x > l.p2.x {
			dx = -1
		}
		dy := 1
		if l.p1.y > l.p2.y {
			dy = -1
		}
		p := l.p1
		for {
			g[p.y][p.x] += 1
			if p == l.p2 {
				break
			}
			p.y += dy
			p.x += dx
		}
		return
	}
	if l.p1.x == l.p2.x {
		for y := utils.IMin(l.p1.y, l.p2.y); y <= utils.IMax(l.p1.y, l.p2.y); y++ {
			g[y][l.p1.x] += 1
		}
	} else if l.p1.y == l.p2.y {
		for x := utils.IMin(l.p1.x, l.p2.x); x <= utils.IMax(l.p1.x, l.p2.x); x++ {
			g[l.p1.y][x] += 1
		}
	}
}

func (g Grid) DrawLines(lines <-chan Line) {
	for l := range lines {
		g.drawLine(l)
	}
}

func (g Grid) CountOverlaps() int {
	count := 0
	for y, _ := range g {
		for x, _ := range g[y] {
			if g[y][x] >= 2 {
				count += 1
			}
		}
	}
	return count
}

func ParseLine(lstr string) (Line, error) {
	var l Line
	n, err := fmt.Sscanf(lstr, "%d,%d -> %d,%d", &l.p1.x, &l.p1.y, &l.p2.x, &l.p2.y)
	if err != nil {
		return l, err
	}
	if n != 4 {
		return l, fmt.Errorf("did not parse correct number of elements: %s", lstr)
	}
	return l, nil
}

func ParseLines(lstrs <-chan string) <-chan Line {
	lineCh := make(chan Line)
	go func() {
		defer close(lineCh)
		for lstr := range lstrs {
			l, err := ParseLine(lstr)
			if err != nil {
				log.Default().Printf("unable to parse Line string: %s: %s", lstr, err.Error())
			}
			lineCh <- l
		}
	}()
	return lineCh
}
