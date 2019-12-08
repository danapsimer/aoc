package main

import (
	"flag"
	"fmt"
)

var (
	serialNumber int
	minSize      = 3
	maxSize      = 3
)

func init() {
	flag.IntVar(&serialNumber, "s", -1, "the serial number to use")
	flag.IntVar(&minSize, "mins", 3, "the minimum size to use")
	flag.IntVar(&maxSize, "maxs", 3, "the maximum size to use")
}

func main() {
	flag.Parse()
	var grid [300][300]int

	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			rackId := x + 10
			power := rackId * y
			power += serialNumber
			power *= rackId
			power /= 100
			power %= 10
			power -= 5
			grid[x-1][y-1] = power
		}
	}
	maxPower := 0
	maxPowerX := 0
	maxPowerY := 0
	maxPowerSize := 0
	for size := minSize; size <= maxSize; size++ {
		for x := 1; x <= 300-size+1; x++ {
			for y := 1; y <= 300-size+1; y++ {
				totalPower := 0
				for dx1 := 0; dx1 < size; dx1++ {
					for dy1 := 0; dy1 < size; dy1++ {
						totalPower += grid[x+dx1-1][y+dy1-1]
					}
				}
				if totalPower > maxPower {
					maxPower = totalPower
					maxPowerX = x
					maxPowerY = y
					maxPowerSize = size
				}
			}
		}
	}
	fmt.Printf("%d,%d,%d\n", maxPowerX, maxPowerY, maxPowerSize)
}
