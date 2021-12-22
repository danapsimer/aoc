package main

import (
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/utils"
	"math"
)

func main() {

	flag.Parse()
	args := flag.Args()
	for _, arg := range args {
		positions, err := utils.ReadIntegersFromFile(arg)
		if err != nil {
			panic(err)
		}
		Part1(positions)
		positions, err = utils.ReadIntegersFromFile(arg)
		if err != nil {
			panic(err)
		}
		Part2(positions)
	}
}

// x = level to align to
// c1 ... cN = current positions
// minimize | c1 - x | + | c2 - x | + ... + | cn - x |

func Part1(in <-chan int) {
	currentPositions := make([]int, 0, 1000)
	min, max := math.MaxInt, 0
	for cp := range in {
		min = utils.IMin(min, cp)
		max = utils.IMax(max, cp)
		currentPositions = append(currentPositions, cp)
	}
	minX := math.MaxInt
	minFuel := math.MaxInt
	for x := min; x <= max; x++ {
		fuel := 0
		for _, cp := range currentPositions {
			fuel += utils.IAbs(cp - x)
		}
		if minFuel > fuel {
			minX = x
			minFuel = fuel
		}
	}
	fmt.Printf("Part1: minX = %d, minFuel = %d\n", minX, minFuel)
}

func Part2(in <-chan int) {
	currentPositions := make([]int, 0, 1000)
	min, max := math.MaxInt, 0
	for cp := range in {
		min = utils.IMin(min, cp)
		max = utils.IMax(max, cp)
		currentPositions = append(currentPositions, cp)
	}
	minX := math.MaxInt
	minFuel := math.MaxInt
	for x := min; x <= max; x++ {
		fuel := 0
		for _, cp := range currentPositions {
			fuel += part2FuelCost(utils.IAbs(cp - x))
		}
		if minFuel > fuel {
			minX = x
			minFuel = fuel
		}
	}
	fmt.Printf("Part2: minX = %d, minFuel = %d\n", minX, minFuel)
}

func part2FuelCost(steps int) int {
	cost := 0
	for s := 1; s <= steps; s++ {
		cost += s
	}
	return cost
}
