package heightmap

import (
	"container/list"
	"sort"
)

type HeightMap [][]int

func LoadHeightMap(lines <-chan string) HeightMap {
	heightMap := make([][]int, 0, 100)
	for line := range lines {
		p := len(heightMap)
		heightMap = append(heightMap, make([]int, 0, 100))
		for _, c := range line {
			heightMap[p] = append(heightMap[p], int(c-'0'))
		}
	}
	return heightMap
}

func (hm HeightMap) IsLowPoint(x, y int) bool {
	// Up
	if y > 0 && hm[y-1][x] <= hm[y][x] {
		return false
	}
	// Down
	if y < len(hm)-1 && hm[y+1][x] <= hm[y][x] {
		return false
	}
	// Left
	if x > 0 && hm[y][x-1] <= hm[y][x] {
		return false
	}
	// Right
	if x < len(hm[y])-1 && hm[y][x+1] <= hm[y][x] {
		return false
	}
	return true
}

func (hm HeightMap) FindLowPointRiskScores() <-chan int {
	scoresCh := make(chan int)
	go func() {
		defer close(scoresCh)
		for y, row := range hm {
			for x, h := range row {
				if hm.IsLowPoint(x, y) {
					scoresCh <- (h + 1)
				}
			}
		}
	}()
	return scoresCh
}

func (hm HeightMap) SumRiskScores() int {
	scores := hm.FindLowPointRiskScores()
	sum := 0
	for score := range scores {
		sum += score
	}
	return sum
}

func (hm HeightMap) CalcBasinSize(x, y int) int {
	calcKey := func(xx, yy int) int {
		return yy*100 + xx
	}
	extractKey := func(key int) (int, int) {
		return key % 100, key / 100
	}
	seen := make(map[int]interface{})
	next := list.New()
	next.PushBack(calcKey(x, y))
	pushIf := func(xx, yy int) {
		key := calcKey(xx, yy)
		if _, ok := seen[key]; !ok && xx >= 0 && yy >= 0 && yy < len(hm) && xx < len(hm[yy]) {
			next.PushBack(key)
		}
	}

	size := 0
	for next.Len() > 0 {
		front := next.Front()
		next.Remove(front)
		if _, ok := seen[front.Value.(int)]; !ok {
			seen[front.Value.(int)] = -1
			x, y := extractKey(front.Value.(int))
			if hm[y][x] != 9 {
				size += 1
				pushIf(x, y-1)
				pushIf(x, y+1)
				pushIf(x-1, y)
				pushIf(x+1, y)
			}
		}
	}
	return size
}

func (hm HeightMap) FindBasins() int {
	basinSizes := make([]int, 0, 100)
	for y, row := range hm {
		for x, _ := range row {
			if hm.IsLowPoint(x, y) {
				basinSizes = append(basinSizes, hm.CalcBasinSize(x, y))
			}
		}
	}
	sort.Ints(basinSizes)
	accumulator := 1
	for i := len(basinSizes) - 1; i >= 0 && i >= len(basinSizes)-3; i-- {
		accumulator *= basinSizes[i]
	}
	return accumulator
}
