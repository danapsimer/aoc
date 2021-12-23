package main

import (
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/day12/caves"
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
	cs, err := caves.LoadCaves(lines)
	if err != nil {
		panic(err)
	}
	count := len(cs.FindPaths(""))
	fmt.Printf("Part1: count of paths = %d\n", count)
}

func part2(filename string) {
	lines, err := utils.ReadLinesFromFile(filename)
	if err != nil {
		panic(err)
	}
	cs, err := caves.LoadCaves(lines)
	if err != nil {
		panic(err)
	}
	count := len(cs.FindPathWith1DuplicateSmallCave())
	fmt.Printf("Part2: count of paths with upto 1 duplicate = %d\n", count)
}
