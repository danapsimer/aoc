package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type dag struct {
	N           int
	V           [][]int
	ParentCount []int
}

func NewDAG(n int) *dag {
	return &dag{n, make([][]int, n), make([]int, n)}
}

func (d *dag) addEdge(source, destination int) {
	d.V[source] = append(d.V[source], destination)
	d.ParentCount[destination] += 1
}

func (d *dag) sort() []int {
	q := list.New()
	pCounts := make([]int, d.N)
	copy(pCounts, d.ParentCount)
	for idx, f := range pCounts {
		if f == 0 {
			q.PushBack(idx)
		}
	}
	l := make([]int, 0, d.N)
	for q.Len() != 0 {
		ue := q.Front()
		if ue == nil {
			panic("Front returned nil")
		}
		q.Remove(ue)
		u := ue.Value.(int)
		log.Printf("u = %d", u)
		l = append(l, u)
		for _, child := range d.V[u] {
			pCounts[child] -= 1
			if pCounts[child] == 0 {
				q.PushBack(child)
			}
		}
	}
	return l
}

func (d *dag) numberOfPaths(source, destination int) int {
	//s := d.sort()
	//log.Printf("topologicalSort = %v", s)
	dp := make([]int, d.N)
	dp[destination] = 1
	for i := d.N - 1; i >= 0; i -= 1 {
		node := i //s[i]
		for _, child := range d.V[node] {
			dp[node] += dp[child]
		}
	}
	return dp[source]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	adapters := make([]int, 0, 1000)
	adapters = append(adapters, 0)
	for scanner.Scan() {
		adapter, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		adapters = append(adapters, adapter)
	}
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	countOf1JoltDiffs := 0
	countOf3JoltDiffs := 0
	n := len(adapters)
	d := NewDAG(n)
	for idx, adapter := range adapters {
		if idx < n-1 {
			diff := adapters[idx+1] - adapter
			if diff == 1 {
				countOf1JoltDiffs += 1
			} else if diff == 3 {
				countOf3JoltDiffs += 1
			} else if diff > 3 {
				panic(fmt.Errorf("found adapter difference > 3 at %d(%d,%d)", idx, adapter, adapters[idx+1]))
			}
			d.addEdge(idx, idx+1)
		}
		if idx < n-2 && adapters[idx+2]-adapter <= 3 {
			d.addEdge(idx, idx+2)
		}
		if idx < n-3 && adapters[idx+3]-adapter <= 3 {
			d.addEdge(idx, idx+3)
		}
	}
	log.Printf("countOf1JoltDiffs(%d) * countOf3JoltDiffs(%d) = %d", countOf1JoltDiffs, countOf3JoltDiffs, countOf1JoltDiffs*countOf3JoltDiffs)
	log.Printf("maximum number of paths from 0 to %d is %d", adapters[n-1], d.numberOfPaths(0, n-1))
}
