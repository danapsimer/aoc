package main

import (
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/day15/chiton"
	"github.com/danapsimer/aoc/2021/utils"
)

func main() {

	flag.Parse()

	args := flag.Args()
	for _, arg := range args {
		part1(arg)
		part2(arg)
	}

}

func part1(fileName string) {
	lines, err := utils.ReadLinesFromFile(fileName)
	if err != nil {
		panic(err)
	}
	g := chiton.LoadGrid(lines)
	shortestPath := g.DijkstraShortestPath()
	risk := g.CalculatePathRisk(shortestPath)
	fmt.Printf("Part1: risk of shortest path = %d\n", risk)
}

func part2(fileName string) {
	lines, err := utils.ReadLinesFromFile(fileName)
	if err != nil {
		panic(err)
	}
	g := chiton.LoadGrid(lines).Expand()
	shortestPath := g.DijkstraShortestPath()
	risk := g.CalculatePathRisk(shortestPath)
	fmt.Printf("Part2: risk of shortest path = %d\n", risk)
}
