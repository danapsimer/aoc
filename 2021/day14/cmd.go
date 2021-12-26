package main

import (
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/day14/polymers"
	"github.com/danapsimer/aoc/2021/utils"
	"math"
)

func main() {

	flag.Parse()
	args := flag.Args()
	for _, arg := range args {
		part1(arg)
	}
}

func part1(fileName string) {
	lines, err := utils.ReadLinesFromFile(fileName)
	if err != nil {
		panic(err)
	}
	polymer, rules, err := polymers.LoadPolymerAndRules(lines)
	if err != nil {
		panic(err)
	}
	counts := polymer.ApplyRules(rules, 10)
	mostCommon, mostCommonCount, leastCommon, leastCommonCount := mostAndLeastCommon(counts)
	fmt.Printf("Part1: Most Common = (%s, %d), LeastCommon = (%s, %d), Most Common Count - Least Common Count = %d\n",
		mostCommon, mostCommonCount, leastCommon, leastCommonCount, mostCommonCount-leastCommonCount)
	counts = polymer.ApplyRules(rules, 40)
	mostCommon, mostCommonCount, leastCommon, leastCommonCount = mostAndLeastCommon(counts)
	fmt.Printf("Part2: Most Common = (%s, %d), LeastCommon = (%s, %d), Most Common Count - Least Common Count = %d\n",
		mostCommon, mostCommonCount, leastCommon, leastCommonCount, mostCommonCount-leastCommonCount)
}

func mostAndLeastCommon(counts map[rune]int) (string, int, string, int) {
	var mostCommon string
	var mostCommonCount int
	var leastCommon string
	var leastCommonCount int = math.MaxInt
	for e, c := range counts {
		if mostCommonCount < c {
			mostCommon = string([]rune{e})
			mostCommonCount = c
		}
		if leastCommonCount > c {
			leastCommon = string([]rune{e})
			leastCommonCount = c
		}
	}
	return mostCommon, mostCommonCount, leastCommon, leastCommonCount
}
