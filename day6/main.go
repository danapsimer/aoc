package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type point struct {
	x, y int
}

func NewPoint(line string) (*point, error) {
	var p point
	_, err := fmt.Sscanf("%d, %d", line, &p.x, &p.y)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func distance(p1, p2 *point) int {
	return abs(p1.x - p2.x) + abs(p1.y - p2.y)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	points := make([]*point,0,50)
	maxX, maxY := 0, 0
	for scanner.Scan() {
		point, err := NewPoint(scanner.Text())
		if err != nil {
			fmt.Printf("ERROR: reading points: %s", err.Error())
			os.Exit(-1)
		}
		points = append(points,point)
		if maxX < point.x {
			maxX = point.x
		}
		if maxY < point.y {
			maxY = point.y
		}
	}
	var grid [maxX+1][maxY+1]int
	for x := 0; x <= maxX; x += 1 {
		for y := 0; y <= maxY; y += 1 {
			point := &point{x,y}
			grid[x][y] = -1
			closestDistance := math.MaxInt32
			for idx, p := range points {
				d := distance(point,p)
				if closestDistance > d {
					grid[x][y] = idx
					closestDistance = d
				} else if closestDistance == d {
					grid[x][y] = -1
				}
			}
		}
	}
	maxDistanceOnGrid := distance(&point{0,0}, &point{maxX,maxY})
	maxArea := 0
	for idx, p := range points {
		area := 1
		for d := 0; d < maxDistanceOnGrid; d += 1 {

		}
	}
}
