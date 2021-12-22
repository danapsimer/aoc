package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/day02/directions"
)

var (
	Debug = flag.Bool("d", false, "turn on debug output")
)

func main() {
	flag.Parse()
	args := flag.Args()
	for _, fn := range args {
		stepCh, err := directions.ReadDirectionsFromFile(context.TODO(), fn)
		if err != nil {
			panic(err.Error())
		}
		x, y := directions.CalculatePosition(0, 0, stepCh)
		fmt.Printf("Part 1 %s: x = %d, y = %d, x*y = %d\n", fn, x, y, x*y)

		stepCh, err = directions.ReadDirectionsFromFile(context.TODO(), fn)
		if err != nil {
			panic(err.Error())
		}
		x, y = directions.CalculatePositionWithAim(0, 0, stepCh)
		fmt.Printf("Part 2 %s: x = %d, y = %d, x*y = %d\n", fn, x, y, x*y)
	}
}
