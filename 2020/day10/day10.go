package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

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
	// For this graph, an ascending sort is a valid topological sort
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	countOf1JoltDiffs := 0
	countOf3JoltDiffs := 0
	n := len(adapters)
	depth := make([]int, n)
	depth[n-1] = 1
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
		}
		i := n - idx - 1
		for j := i - 1; j >= 0 && j >= i-3; j -= 1 {
			if adapters[i]-adapters[j] <= 3 {
				depth[j] += depth[i]
			}
		}
	}
	log.Printf("countOf1JoltDiffs(%d) * countOf3JoltDiffs(%d) = %d", countOf1JoltDiffs, countOf3JoltDiffs, countOf1JoltDiffs*countOf3JoltDiffs)
	log.Printf("maximum number of paths from 0 to %d is %d", adapters[n-1], depth[0])
}
