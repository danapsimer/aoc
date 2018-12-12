package main

import (
	"bufio"
	"fmt"
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
	count := 0
	for y := 0; y <= maxY; y += 1 {
		for x := 0; x <= maxX; x += 1 {
			point := &point{x, y}
			totalOfDistances := 0
			for _, p := range points {
				totalOfDistances += distance(point, p)
				if totalOfDistances > 10000 {
					break
				}
			}
			if totalOfDistances < 10000 {
				count += 1
			}
		}
	}
	fmt.Printf("%d\n", count)
}
