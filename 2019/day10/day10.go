package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"math/cmplx"
	"os"
	"sort"
)

type node struct {
	asteroid bool
	x, y     int
	visible  []*node
}

type grid struct {
	nodes     [][]*node
	asteroids []*node
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func ReduceSlope(dx, dy int) (int, int) {
	adx := dx
	if adx < 0 {
		adx = -adx
	}
	ady := dy
	if ady < 0 {
		ady = - ady
	}
	for {
		gcd := GCD(adx, ady)
		if gcd == 1 {
			break
		}
		ady = ady / gcd
		adx = adx / gcd
	}
	if dx < 0 {
		adx = -adx
	}
	if dy < 0 {
		ady = -ady
	}
	return adx, ady
}

func (g *grid) FindBestLocation() *node {

	for _, asteroid1 := range g.asteroids {
		g.populateVisible(asteroid1)
	}
	maxCount := 0
	var maxAsteroid *node
	for _, asteroid := range g.asteroids {
		if maxCount < len(asteroid.visible) {
			maxCount = len(asteroid.visible)
			maxAsteroid = asteroid
		}
	}
	return maxAsteroid
}

func (g *grid) populateVisible(asteroid1 *node) {
	maxX := len(g.nodes[0])
	maxY := len(g.nodes)
	for _, asteroid2 := range g.asteroids {
		if asteroid2 != asteroid1 {
			dx, dy := ReduceSlope(asteroid2.x-asteroid1.x, asteroid2.y-asteroid1.y)
			x, y := asteroid1.x, asteroid1.y
			blocked := false
			for !blocked {
				x, y = x+dx, y+dy
				if x < 0 || x >= maxX || y < 0 || y >= maxY {
					panic("exited the field before finding target asteroid")
				}
				if x == asteroid2.x && y == asteroid2.y {
					break
				}
				if g.nodes[y][x].asteroid {
					blocked = true
				}
			}
			if !blocked {
				asteroid1.visible = append(asteroid1.visible, asteroid2)
			}
		}
	}
}

func theta(x1, y1, x2, y2 int) float64 {
	dx := x1 - x2
	dy := y2 - y1
	r := -cmplx.Phase(complex(-float64(dy), float64(dx)))
	if r < 0 {
		r = 2*math.Pi + r
	}
	//log.Printf("theta(%d,%d,%d,%d) = %f", x1, y1, x2, y2, r)
	return r
}

func (n *node) Len() int {
	return len(n.visible)
}

func (n *node) Less(i, j int) bool {
	iTheta := theta(n.x, n.y, n.visible[i].x, n.visible[i].y)
	jTheta := theta(n.x, n.y, n.visible[j].x, n.visible[j].y)
	return iTheta < jTheta
}

func (n *node) Swap(i, j int) {
	tmp := n.visible[i]
	n.visible[i] = n.visible[j]
	n.visible[j] = tmp
}

func (g *grid) EliminateVisibleUntil200th(n *node, c int) (*node, int) {
	sort.Sort(n)
	nodes := n.visible
	n.visible = make([]*node, 0, 100)
	for _, v := range nodes {
		v.asteroid = false
		for i := range g.asteroids {
			if g.asteroids[i] == v {
				if i+1 < len(g.asteroids) {
					g.asteroids = append(g.asteroids[:i], g.asteroids[i+1:]...)
				} else {
					g.asteroids = g.asteroids[:i]
				}
				break
			}
		}
		c += 1
		log.Printf("Eliminated #%d asteroid at (%d, %d)", c, v.x, v.y)
		if c == 200 {
			return v, c
		}
	}
	return nil, c
}

func ReadMap(in io.Reader) *grid {
	mm := make([][]*node, 0, 1000)
	wa := make([]*node, 0, 1000)
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]*node, len(line))
		for i, c := range line {
			row[i] = &node{c == '#', i, len(mm), make([]*node, 0, 100)}
			if row[i].asteroid {
				wa = append(wa, row[i])
			}
		}
		mm = append(mm, row)
	}
	return &grid{mm, wa}
}

func main() {
	g := ReadMap(os.Stdin)
	best := g.FindBestLocation()
	log.Printf("part1 = %d @ (%d, %d)", len(best.visible), best.x, best.y)
	c := 0
	sweep := 0
	for {
		last, c := g.EliminateVisibleUntil200th(best, c)
		if c == 200 {
			log.Printf("part2 = %d", last.x*100+last.y)
			os.Exit(0)
		}
		sweep += 1
		log.Printf("Completed sweep #%d", sweep)
		g.populateVisible(best)
	}
}
