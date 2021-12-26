package chiton

import (
	"container/list"
	"math"
)

type grid [][]int
type point struct{ x, y int }

func (g grid) CalculatePathRisk(path []point) int {
	risk := 0
	for i, p := range path {
		if i > 0 {
			risk += g[p.y][p.x]
		}
	}
	return risk
}

func (g grid) FindShortestPath() ([]point, int) {
	queue := list.New()
	startPoint := point{0, 0}
	endPoint := point{len(g[0]) - 1, len(g) - 1}
	queue.PushBack([]point{startPoint})
	visited := make(map[point]bool)
	pushIf := func(path []point) {
		p := path[len(path)-1]
		if p.x > 0 && p.y > 0 && p.x < len(g[0]) && p.y < len(g) {
			for _, pp := range path[:len(path)-1] {
				if pp == p {
					return
				}
			}
			queue.PushBack(path)
		}
	}

	completedPaths := make([][]point, 0, 1000)
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		path := front.Value.([]point)
		pathEnd := path[len(path)-1]
		if pathEnd == endPoint {
			completedPaths = append(completedPaths, path)
		}
		pushIf(append(path, point{pathEnd.x + 1, pathEnd.y}))
		pushIf(append(path, point{pathEnd.x - 1, pathEnd.y}))
		pushIf(append(path, point{pathEnd.x, pathEnd.y + 1}))
		pushIf(append(path, point{pathEnd.x, pathEnd.y - 1}))
	}

	var leastRiskyPath []point
	leastRiskyPathValue := math.MaxInt
	for _, path := range completedPaths {
		risk := g.CalculatePathRisk(path)
		if risk < leastRiskyPathValue {
			leastRiskyPath = path
			leastRiskyPathValue = risk
		}
	}
	return leastRiskyPath, leastRiskyPathValue
}

func LoadGrid(lines <-chan string) grid {
	g := make(grid, 0, 1000)
	for line := range lines {
		row := make([]int, 0, 1000)
		g = append(g, row)
		for _, c := range line {
			row = append(row, int(c-'0'))
		}
	}
	return g
}
