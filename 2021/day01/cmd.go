package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/day01/location"
)

var (
	Debug = flag.Bool("d", false, "turn on debug output")
)

func init() {
}

func main() {
	flag.Parse()
	args := flag.Args()
	for _, arg := range args {
		count, err := location.CountIncreasesFromPreviousInFile(context.TODO(), arg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Part 1 %s: %d\n", arg, count)
		count, err = location.CountIncreasesFromPreviousInFilePart2(context.TODO(), arg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Part 2 %s: %d\n", arg, count)
	}
}
