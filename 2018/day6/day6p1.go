package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	points := make([]*point, 0, 50)
	line, maxX, maxY := 0, 0, 0
	for scanner.Scan() {
		point, err := NewPoint(scanner.Text())
		if err != nil {
			fmt.Printf("ERROR: reading points: (%d) '%s': %s\n", line, scanner.Text(), err.Error())
			os.Exit(-1)
		}
		points = append(points, point)
		if maxX < point.x {
			maxX = point.x
		}
		if maxY < point.y {
			maxY = point.y
		}
		line += 1
	}
	counts := make([]int, len(points))
	for y := 0; y <= maxY; y += 1 {
		for x := 0; x <= maxX; x += 1 {
			point := &point{x, y}
			xyIdx := -1
			closestDistance := math.MaxInt32
			for idx, p := range points {
				d := distance(point, p)
				if closestDistance > d {
					xyIdx = idx
					closestDistance = d
				} else if closestDistance == d {
					xyIdx = -1
				}
			}
			if xyIdx >= 0 {
				if x == 0 || y == 0 || x == maxX || y == maxY {
					counts[xyIdx] = -1
				} else if counts[xyIdx] != -1 {
					counts[xyIdx] += 1
				}
			}
			if xyIdx == -1 {
				fmt.Print("**** ")
			} else if closestDistance == 0 {
				fmt.Printf("[%-2d] ", xyIdx)
			} else {
				fmt.Printf("(%-2d) ", xyIdx)
			}
		}
		fmt.Print("\n")
	}
	maxSize := 0
	for _, count := range counts {
		if count != -1 && maxSize < count {
			maxSize = count
		}
	}
	fmt.Printf("%d\n", maxSize)
}
