package main

import (
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/day06/lf"
	"github.com/danapsimer/aoc/2021/utils"
)

func main() {
	flag.Parse()

	args := flag.Args()
	for _, arg := range args {
		part1(80, arg)
		part1(256, arg)
	}
}

func part1(n int, arg string) {
	fish, err := utils.ReadIntegersFromFile(arg)
	if err != nil {
		panic(err)
	}
	count := lf.NDaysStatic(n, fish)
	fmt.Printf("count after %d days = %d\n", n, count)
}
