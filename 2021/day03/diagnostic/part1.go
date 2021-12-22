package diagnostic

import (
	"fmt"
)

func ReadInputFromChannel(lines <-chan string) (<-chan uint, int, error) {
	intCh := make(chan uint, 1)
	firstLine := <-lines
	value := parseLine(firstLine)
	intCh <- value

	go func() {
		defer close(intCh)
		for line := range lines {
			intCh <- parseLine(line)
		}
	}()
	return intCh, len(firstLine), nil
}

func parseLine(str string) uint {
	var value uint
	n, err := fmt.Sscanf(str, "%b", &value)
	if err != nil || n != 1 {
		panic(fmt.Sprintf("cannot parse %s as a binary integer", str))
	}
	return value
}

func CalculateGammaAndEpsilon(values <-chan uint, width int) (uint, uint) {
	counts := make([]int, width, width)
	n := 0
	for value := range values {
		for b, _ := range counts {
			if value&(1<<b) != 0 {
				counts[b] += 1
			}
		}
		n += 1
	}
	gamma := uint(0)
	epsilon := uint(0)
	for b, c := range counts {
		if n-c < n/2 {
			gamma = gamma | (1 << b)
		} else {
			epsilon = epsilon | (1 << b)
		}
	}
	return gamma, epsilon
}
