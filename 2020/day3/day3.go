package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type grid []string

func (g grid) get(x, y int) string {
	xidx := x % len(g[0])
	return g[y][xidx : xidx+1]
}

func readGrid(reader io.Reader) grid {
	var g grid = make([]string, 0, 1000)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		g = append(g, scanner.Text())
	}
	return g
}

func main() {
	g := readGrid(os.Stdin)
	treeCount11 := countTrees(g, 1, 1)
	log.Printf("treeCount11 = %d", treeCount11)
	treeCount31 := countTrees(g, 3, 1)
	log.Printf("treeCount31 = %d", treeCount31)
	treeCount51 := countTrees(g, 5, 1)
	log.Printf("treeCount51 = %d", treeCount51)
	treeCount71 := countTrees(g, 7, 1)
	log.Printf("treeCount71 = %d", treeCount71)
	treeCount12 := countTrees(g, 1, 2)
	log.Printf("treeCount12 = %d", treeCount12)
	log.Printf("result = %d", treeCount11*treeCount12*treeCount31*treeCount51*treeCount71)
}

func countTrees(g grid, dx int, dy int) int {
	x, y := 0, 0
	treeCount := 0
	for y < len(g) {
		if g.get(x, y) == "#" {
			treeCount += 1
		}
		x += dx
		y += dy
	}
	return treeCount
}
