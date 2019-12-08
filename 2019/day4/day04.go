package main

import (
	"fmt"
	"strconv"
)

func checkPassword(password int) bool {
	passwordString := strconv.Itoa(password)
	if len(passwordString) == 6 {
		var last rune
		repeatCounts := make([]int, 10)
		for i, c := range passwordString {
			if i > 0 {
				if last > c {
					return false
				}
				if last == c {
					repeatCounts[c-'0'] += 1
				}
			}
			last = c
		}
		for _, c := range repeatCounts {
			if c == 1 {
				return true
			}
		}
	}
	return false
}

func findPasswordCount(start, end int) int {
	count := 0
	for password := start; password <= end; password++ {
		if checkPassword(password) {
			count += 1
		}
	}
	return count
}

func main() {
	fmt.Printf("%d\n", findPasswordCount(356261, 846303))
}
