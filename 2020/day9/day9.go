package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func isSumOf(n int, a []int) bool {
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if a[i] + a[j] == n  {
				return true
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := make([]int,0,1001)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		input = append(input, n)
	}
	l := 0
	last25 := make([]int,25)
	var vulnerability int
	for _, n := range input {
		if l < 25 {
			last25[l] = n
		} else {
			if !isSumOf(n, last25) {
				vulnerability = n
				break
			}
			copy(last25, last25[1:25])
			last25[24] = n
		}
		l += 1
	}
	log.Printf("first invalid value %d\n", vulnerability)

outer:
	for i, _ := range input {
		min := math.MaxInt32
		max := 0
		sum := 0
		for j := i; j < len(input); j++ {
			sum += input[j];
			if min > input[j] {
				min = input[j]
			}
			if max < input[j] {
				max = input[j]
			}
			if sum == vulnerability {
				log.Printf("smallest = %d + largest = %d = %d", min, max, min + max)
				break outer
			}
		}
	}
}
