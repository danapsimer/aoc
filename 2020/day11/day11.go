package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type grid []string

func readGrid(reader io.Reader) grid {
	scanner := bufio.NewScanner(reader)
	g := make(grid, 0, 1000)
	for scanner.Scan() {
		g = append(g, scanner.Text())
	}
	return g
}

func (g grid) get(x, y int) string {
	if y < 0 || y >= len(g) || x < 0 || x >= len(g[y]) {
		return "."
	}
	return g[y][x : x+1]
}

func (g grid) height() int {
	return len(g)
}

func (g grid) width() int {
	return len(g[0])
}

func (g grid) equals(o grid) bool {
	if g.height() == o.height() && g.width() == o.width() {
		for y := range g {
			if g[y] != o[y] {
				return false
			}
		}
		return true
	}
	return false
}

func (g grid) applyRules() grid {
	cp := make(grid, g.height())
	copy(cp, g)
	for y := 0; y < g.height(); y++ {
		for x := 0; x < g.width(); x++ {
			c := g.get(x, y)
			if c == "." {
				continue
			}
			adjOccupiedCount := 0
			for yd := -1; yd <= 1; yd++ {
				for xd := -1; xd <= 1; xd++ {
					if (xd != 0 || yd != 0) && g.get(x+xd, y+yd) == "#" {
						adjOccupiedCount += 1
					}
				}
			}
			if c == "L" && adjOccupiedCount == 0 {
				cp[y] = cp[y][:x] + "#" + cp[y][x+1:]
			} else if c == "#" && adjOccupiedCount >= 4 {
				cp[y] = cp[y][:x] + "L" + cp[y][x+1:]
			}
		}
	}
	return cp
}

func (g grid) applyRulesPart2() grid {
	cp := make(grid, g.height())
	copy(cp, g)
	for y := 0; y < g.height(); y++ {
		for x := 0; x < g.width(); x++ {
			c := g.get(x, y)
			if c == "." {
				continue
			}
			visibleOccupiedCount := 0
			for yd := -1; yd <= 1; yd++ {
				for xd := -1; xd <= 1; xd++ {
					if xd == 0 && yd == 0 {
						continue
					}
					for nx, ny := x+xd, y+yd; 0 <= nx && nx < g.width() && 0 <= ny && ny < g.height(); nx, ny = nx+xd, ny+yd {
						nc := g.get(nx, ny)
						if nc == "#" {
							visibleOccupiedCount += 1
							break
						} else if nc == "L" {
							break
						}
					}
				}
			}
			if c == "L" && visibleOccupiedCount == 0 {
				cp[y] = cp[y][:x] + "#" + cp[y][x+1:]
			} else if c == "#" && visibleOccupiedCount >= 5 {
				cp[y] = cp[y][:x] + "L" + cp[y][x+1:]
			}
		}
	}
	return cp
}

func (g grid) String() string {
	return strings.Join(g, "\n")
}

func (original grid) applyRulesUntilStable(rulesFn func(g grid) grid) grid {
	g := original
	cp := g.applyRules()
	for !cp.equals(g) {
		g = cp
		cp = rulesFn(g)
	}
	return cp
}

func (g grid) countOccupiedSeats() int {
	occupied := 0
	for y := range g {
		for x := 0; x < g.width(); x++ {
			if g.get(x, y) == "#" {
				occupied += 1
			}
		}
	}
	return occupied
}

func main() {
	original := readGrid(os.Stdin)
	g := original.applyRulesUntilStable(func(g grid) grid {
		return g.applyRules()
	})
	log.Printf("part 1 occupied = %d", g.countOccupiedSeats())
	g = original.applyRulesUntilStable(func(g grid) grid {
		return g.applyRulesPart2()
	})
	log.Printf("part 2 occupied = %d", g.countOccupiedSeats())
}
