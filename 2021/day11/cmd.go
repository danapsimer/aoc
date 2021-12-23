package main

import (
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/day11/octo"
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
	octopi := octo.LoadOctopi(lines)
	count := octo.NRounds(100, octopi)
	fmt.Printf("Part1: flash count = %d\n", count)
}

func part2(filename string) {
	lines, err := utils.ReadLinesFromFile(filename)
	if err != nil {
		panic(err)
	}
	octopi := octo.LoadOctopi(lines)
	round := 1
	for {
		count := octo.Round(octopi)
		if count == 100 {
			break
		}
		round += 1
	}
	fmt.Printf("Part2: first round when all octopuses flash: %d\n", round)
}
