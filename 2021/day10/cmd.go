package main

import (
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/day10/chunks"
	"github.com/danapsimer/aoc/2021/utils"
)

func main() {

	flag.Parse()
	args := flag.Args()
	for _, arg := range args {
		part1(arg)
	}
}

func part1(filename string) {
	lines, err := utils.ReadLinesFromFile(filename)
	if err != nil {
		panic(err)
	}

	score, completionScore := chunks.CheckLines(lines)
	fmt.Printf("Part1: score = %d\n", score)
	fmt.Printf("Part2: completionScore = %d\n", completionScore)
}
