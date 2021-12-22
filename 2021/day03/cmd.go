package main

import (
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/day03/diagnostic"
	"github.com/danapsimer/aoc/2021/utils"
)

var (
	Debug = flag.Bool("d", false, "turn on debug output")
)

func main() {
	flag.Parse()
	args := flag.Args()
	for _, fn := range args {
		part1(fn)
		part2(fn)
	}
}

func part1(fn string) {
	lines, err := utils.ReadLinesFromFile(fn)
	if err != nil {
		fmt.Printf("error reading from file: %s: %s", fn, err.Error())
	} else {
		input, width, err := diagnostic.ReadInputFromChannel(lines)
		if err != nil {
			fmt.Printf("error parsing lines from file: %s: %s", fn, err.Error())
		} else {
			gamma, epsilon := diagnostic.CalculateGammaAndEpsilon(input, width)
			fmt.Printf("Part 1: Gamma = %b, Epsilon = %b, Gamma * Epsilon = %d\n", gamma, epsilon, gamma*epsilon)
		}
	}
}

func part2(fn string) {
	lines, err := utils.ReadLinesFromFile(fn)
	if err != nil {
		fmt.Printf("error reading from file: %s: %s", fn, err.Error())
	} else {
		input, width, err := diagnostic.ReadInputFromChannel(lines)
		if err != nil {
			fmt.Printf("error parsing lines from file: %s: %s", fn, err.Error())
		} else {
			oxyRating, co2Rating := diagnostic.CalculateRatings(input, width)
			fmt.Printf("Part 2: oxyRating = %b, co2Rating = %b, oxyRating * co2Rating = %d\n", oxyRating, co2Rating, oxyRating*co2Rating)
		}
	}
}
