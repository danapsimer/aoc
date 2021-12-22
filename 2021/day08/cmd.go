package main

import (
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/day08/decoder"
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

func part1(arg string) {
	lines, err := utils.ReadLinesFromFile(arg)
	if err != nil {
		panic(err)
	}
	count := decoder.CountUniqueOutputs(decoder.LoadDecoderLines(lines))
	fmt.Printf("Part1: unique output count = %d\n", count)
}

func part2(arg string) {
	lines, err := utils.ReadLinesFromFile(arg)
	if err != nil {
		panic(err)
	}
	sum := decoder.SumOutputValues(decoder.LoadDecoderLines(lines))
	fmt.Printf("Part2: sum of outputs = %d\n", sum)
}
