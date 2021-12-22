package main

import (
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/day05/grid"
	"github.com/danapsimer/aoc/2021/utils"
)

func main() {

	flag.Parse()

	args := flag.Args()
	for _, arg := range args {
		lineStrs, err := utils.ReadLinesFromFile(arg)
		if err != nil {
			panic(err)
		}
		g := grid.NewGrid(1000, 1000)
		g.DrawLines(grid.ParseLines(lineStrs))
		count := g.CountOverlaps()
		fmt.Printf("Part 1: overlap count = %d\n", count)
	}
}
