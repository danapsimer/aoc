package chiton

import (
	"bytes"
	"container/list"
	"fmt"
)

type Grid [][]int
type point struct{ x, y int }
type path []point

func (pth path) Append(p point) (bool, path) {
	for _, pp := range pth {
		if pp == p {
			return false, pth
		}
	}
	return true, append(pth, p)
}

func (g Grid) Width() int {
	return len(g[0])
}

func (g Grid) Height() int {
	return len(g)
}

func (g Grid) CalculatePathRisk(pth path) int {
	risk := 0
	for i, p := range pth {
		if i > 0 {
			risk += g[p.y][p.x]
		}
	}
	return risk
}

func (g Grid) DijkstraShortestPath() path {
	startPoint := point{0, 0}
	endPoint := point{len(g[0]) - 1, len(g) - 1}
	distance := make(map[point]int)
	prev := make(map[point]point)
	visited := make(map[point]bool)
	queue := list.New()
	pushIf := func(u, v point) {
		if v.x < 0 || v.y < 0 || v.y >= len(g) || v.x >= len(g[v.y]) {
			return
		}
		if _, visited := visited[v]; !visited {
			alt := distance[u] + g[v.y][v.x]
			if distV, ok := distance[v]; !ok || alt < distV {
				distance[v] = alt
				prev[v] = u
				queue.PushBack(v)
			}
		}
	}
	visited[startPoint] = true
	distance[startPoint] = 0
	queue.PushBack(startPoint)
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		current := front.Value.(point)
		if current == endPoint {
			shortestPath := path{current}
			for {
				var ok bool
				if current, ok = prev[current]; ok {
					shortestPath = append(path{current}, shortestPath...)
				} else {
					return shortestPath
				}
			}
			return shortestPath
		}
		pushIf(current, point{current.x, current.y + 1})
		pushIf(current, point{current.x, current.y - 1})
		pushIf(current, point{current.x + 1, current.y})
		pushIf(current, point{current.x - 1, current.y})
	}
	return nil
}

func (g Grid) Expand() Grid {
	newGrid := make(Grid, len(g)*5)
	height := len(g)
	width := len(g[0])
	for y, row := range g {
		newGrid[y] = make([]int, width*5)
		copy(newGrid[y], g[y])
		for x, value := range row {
			for d := 1; d < 5; d++ {
				newGrid[y][x+d*width] = 1 + ((value-1)+d)%9
			}
		}
	}
	for dy := 1; dy < 5; dy++ {
		for y := dy * height; y < (dy+1)*height; y++ {
			newGrid[y] = make([]int, width*5)
			for dx := 0; dx < 5; dx++ {
				for x := 0; x < width; x++ {
					newGrid[y][x+dx*width] = 1 + ((newGrid[y-height][x+dx*width]-1)+1)%9
				}
			}
		}
	}
	return newGrid
}

func (g Grid) StringWithDividers(n int) string {
	buf := &bytes.Buffer{}
	width := len(g[0])
	for y, row := range g {
		if y > 0 && y%n == 0 {
			for x := 0; x < (width + width/n); x++ {
				fmt.Fprint(buf, "-")
			}
			fmt.Fprintln(buf)
		}
		for x, c := range row {
			if x > 0 && x%n == 0 {
				fmt.Fprint(buf, "|")
			}
			fmt.Fprint(buf, string([]rune{rune('0' + c)}))
		}
		fmt.Fprintln(buf)
	}
	return buf.String()
}

func (g Grid) String() string {
	return g.StringWithDividers(len(g))
}

func LoadGrid(lines <-chan string) Grid {
	g := make(Grid, 0, 1000)
	for line := range lines {
		row := make([]int, 0, 1000)
		for _, c := range line {
			row = append(row, int(c-'0'))
		}
		g = append(g, row)
	}
	return g
}
