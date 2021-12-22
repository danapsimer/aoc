package main

import (
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/day09/heightmap"
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

func part1(filename string) {
	lines, err := utils.ReadLinesFromFile(filename)
	if err != nil {
		panic(err)
	}
	hm := heightmap.LoadHeightMap(lines)
	sum := hm.SumRiskScores()
	fmt.Printf("Part1: sum of risk scores = %d\n", sum)
}
func part2(filename string) {
	lines, err := utils.ReadLinesFromFile(filename)
	if err != nil {
		panic(err)
	}
	hm := heightmap.LoadHeightMap(lines)
	sum := hm.FindBasins()
	fmt.Printf("Part2: multiple of basin sizes = %d\n", sum)
}
