package main

import (
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/day13/paper"
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

	paper, instructions, err := paper.LoadPaperAndInstructions(lines)
	if err != nil {
		panic(err)
	}
	paper.Fold(instructions[0])
	count := paper.CountDots()
	fmt.Printf("Part1: dots after first fold: %d\n", count)
}

func part2(fileName string) {
	lines, err := utils.ReadLinesFromFile(fileName)
	if err != nil {
		panic(err)
	}

	paper, instructions, err := paper.LoadPaperAndInstructions(lines)
	if err != nil {
		panic(err)
	}
	for _, instruction := range instructions {
		paper.Fold(instruction)
	}
	str := paper.String()
	fmt.Printf("Part2:\n%s\n", str)
}
