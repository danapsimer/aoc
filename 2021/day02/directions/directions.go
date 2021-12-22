package directions

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type direction int

const (
	forward direction = iota
	up
	down
)

func directionFromString(dirStr string) direction {
	switch dirStr {
	case "forward":
		return forward
	case "up":
		return up
	case "down":
		return down
	default:
		panic(fmt.Sprintf("unknown direction %s", dirStr))
	}
}

type step struct {
	dir   direction
	delta int
}

func ParseStep(stepStr string) step {
	fields := strings.Split(stepStr, " ")
	if len(fields) != 2 {
		panic(fmt.Sprintf("expected 2 fields: %s", stepStr))
	}
	delta, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(fmt.Sprintf("cannot parse integer: %s: %s", fields[1], err.Error()))
	}
	return step{
		directionFromString(fields[0]),
		delta,
	}
}

func ReadDirectionsFromFile(ctx context.Context, filename string) (<-chan step, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	steps := make(chan step)
	go func() {
		defer f.Close()
		defer close(steps)
		s := bufio.NewScanner(f)
		for s.Scan() {
			steps <- ParseStep(s.Text())
		}
	}()
	return steps, nil
}

func CalculatePosition(x, y int, steps <-chan step) (int, int) {
	for step := range steps {
		switch step.dir {
		case forward:
			x += step.delta
		case up:
			y -= step.delta
		case down:
			y += step.delta
		}
	}
	return x, y
}

func CalculatePositionWithAim(x, y int, steps <-chan step) (int, int) {
	a := 0
	for step := range steps {
		switch step.dir {
		case forward:
			x += step.delta
			y += a * step.delta
		case up:
			a -= step.delta
		case down:
			a += step.delta
		}
	}
	return x, y
}
